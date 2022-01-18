package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/ishitb/mossgo/moss"
	"github.com/ishitb/mossgo/utils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// reviewCmd represents the review command
var reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Send files for review",
	Long: `
	Usage Instructions:
		The -l option specifies the source language of the tested programs.
		Moss supports many different languages; see below for the
		full list.
		["c", "cc", "java", "ml", "pascal", "ada", "lisp", "scheme", "haskell", "fortran", "ascii", "vhdl", "perl", "matlab", "python", "mips", "prolog", "spice", "vb", "csharp", "modula2", "a8086", "javascript", "plsql"]
		Example: Compare the lisp programs foo.lisp and bar.lisp:
			moss -l lisp foo.lisp bar.lisp
	
		The -d option specifies that submissions are by directory, not by file.
		That is, files in a directory are taken to be part of the same program,
		and reported matches are organized accordingly by directory.
		Example: Compare the programs foo and bar, which consist of .c and .h
		files in the directories foo and bar respectively.
			moss -d foo/*.c foo/*.h bar/*.c bar/*.h
			
		Example: Each program consists of the *.c and *.h files in a directory under
		the directory "assignment1."
			moss -d assignment1/*/*.h assignment1/*/*.c

		The -b option names a "base file". When a base file is supplied,
		program code that also appears in the base file is not counted in matches.
		A typical base file will include, for example, the instructor-supplied 
		IMPORTANT: Unlike previous versions of moss, the -b option *always*
		takes a single filename, even if the -d option is also used.
		Examples: Submit all of the C++ files in the current directory, using skeleton.cc
		as the base file:
			moss -l cc -b skeleton.cc *.cc
		Submit all of the ML programs in directories asn1.96/* and asn1.97/*, where
		asn1.97/instructor/example.ml and asn1.96/instructor/example.ml contain the base files.
			moss -l ml -b asn1.97/instructor/example.ml -b asn1.96/instructor/example.ml -d asn1.97/*/*.ml asn1.96/*/*.ml 
		
		
		The -m option sets the maximum number of times a given passage may appear
		before it is ignored.  A passage of code that appears in many programs
		is probably legitimate sharing and not the result of plagiarism.  With -m N,
		any passage appearing in more than N programs is treated as if it appeared in 
		a base file (i.e., it is never reported). The default for -m is 10.
		Examples:
			moss -l pascal -m 2 *.pascal 
			moss -l cc -m 1000000 -b mycode.cc asn1/*.cc
		
		The -c option supplies a comment string that is attached to the generated
		report.
		Example:
			moss -l scheme -c "Scheme programs" *.sch
		
		The -n option determines the number of matching files to show in the results.
		The default is 250.
		Example:
			moss -c java -n 200 *.java
		
		The -x option sends queries to the current experimental version of the server.
		Example:
			moss -x -l ml *.ml
	`,
	Run: func(cmd *cobra.Command, args []string) {
		review(cmd, args)
	},
	Args: cobra.MinimumNArgs(1),
}

func review(cmd *cobra.Command, args []string) {
	files := args

	basefiles, _ := cmd.Flags().GetStringArray("basefile")
	comment, _ := cmd.Flags().GetString("comment")
	directory, _ := cmd.Flags().GetBool("directory")
	experimental, _ := cmd.Flags().GetBool("experimental")
	language, _ := cmd.Flags().GetString("language")
	maxSimilarities, _ := cmd.Flags().GetInt64("maxSimilarities")
	noOfMatchingFiles, _ := cmd.Flags().GetInt64("show")

	// Choosing language if flag not provided
	if language == "" {
		language = chooseLanguage()
		fmt.Println("Language:", language)
	}

	// Reading creds.json for unique ID
	credsJson := utils.ReadFile("creds.json")
	var creds map[string]interface{}
	err := json.Unmarshal([]byte(credsJson), &creds)

	if err != nil {
		utils.ErrorP(err.Error())
	}

	// Getting Unique ID
	uniqueIdRaw := creds["uniqueId"]
	uniqueId := ""
	if uniqueIdRaw == nil {
		reader := bufio.NewReader(os.Stdin)
		uniqueId = utils.GetInput("Enter your one time unique MOSS ID", reader)
		utils.Warn("Please use the login command to save your unique ID\n")
	} else {
		uniqueId = fmt.Sprintf("%0.0f", uniqueIdRaw.(float64))
	}

	mossObj := moss.NewMoss(uniqueId)

	// Adding Base Files
	for _, basefile := range basefiles {
		if strings.Contains(basefile, "*") {
			mossObj.AddBaseFilesByWildcard(basefile)
		} else {
			mossObj.AddBaseFile(basefile)
		}
	}

	// Adding Files
	for _, file := range files {
		if strings.Contains(file, "*") {
			mossObj.AddFilesByWildcard(file)
		} else {
			mossObj.AddFile(file)
		}
	}

	// Setting flags on moss object
	mossObj.SetComment(comment)
	if directory {
		mossObj.SetDirectory()
	}
	if experimental {
		mossObj.SetExperimental()
	}
	mossObj.SetLanguage(language)
	mossObj.SetMaxLimit(int(maxSimilarities))
	mossObj.SetNoOfMatchingFiles(int(noOfMatchingFiles))

	result := mossObj.SendForReview()
	utils.Warn(result + "\n")
}

func chooseLanguage() string {
	languages := moss.Languages

	utils.Info("Choose a language:")
	prompt := promptui.Select{
		Label: "Select Source Language",
		Items: languages,
	}
	result, _, err := prompt.Run()

	if err != nil {
		utils.ErrorP("Prompt failed %v\n", err)
	}

	return languages[result]
}

func init() {
	rootCmd.AddCommand(reviewCmd)

	reviewCmd.Flags().StringP(
		"language",
		"l",
		"",
		"Specify the source language of the tested programs | If not specified then a menu will be displayed asking the same",
	)
	reviewCmd.Flags().BoolP(
		"directory",
		"d",
		false,
		"Specifies that submissions are by directory, not by file.",
	)
	reviewCmd.Flags().StringArrayP(
		"basefile",
		"b",
		[]string{""},
		"Specifies base file paths which are instructor-supplied code for an assignment.  Multiple -b options are allowed.",
	)
	reviewCmd.Flags().Int64P(
		"maxSimilarities",
		"m",
		10,
		"Specifies the maximum number of times a given passage may appear before it is ignored",
	)
	reviewCmd.Flags().Int64P(
		"show",
		"n",
		250,
		"Specifies the maximum number of file comaprisons to be made",
	)
	reviewCmd.Flags().StringP(
		"comment",
		"c",
		"",
		"Supplies a comment string that is attached to the generated report",
	)
	reviewCmd.Flags().BoolP(
		"experimental",
		"x",
		false,
		"Sends queries to the current experimental version of the server.",
	)
}
