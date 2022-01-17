package moss

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/ishitb/moss.go/utils"
)

type Moss struct {
	Languages []string
	server    string
	port      int
	unique_id string

	// Flags
	language             string
	directory            int
	base_files           []string
	max_limit            int
	comment              string
	no_of_matching_files int
	experimental         int

	// Args
	files []string
}

var Languages = []string{
	"c",
	"cc",
	"java",
	"ml",
	"pascal",
	"ada",
	"lisp",
	"scheme",
	"haskell",
	"fortran",
	"ascii",
	"vhdl",
	"verilog",
	"perl",
	"matlab",
	"python",
	"mips",
	"prolog",
	"spice",
	"vb",
	"csharp",
	"modula2",
	"a8086",
	"javascript",
	"plsql",
}

func NewMoss(unique_id string) Moss {
	moss := Moss{
		Languages:            Languages,
		server:               "moss.stanford.edu",
		port:                 7690,
		unique_id:            unique_id,
		language:             "c",
		directory:            0,
		base_files:           []string{},
		max_limit:            10,
		files:                []string{},
		no_of_matching_files: 250,
		experimental:         0,
	}

	return moss
}

func (moss *Moss) SetLanguage(language string) {
	if language == "cpp" {
		language = "cc"
	}
	(*moss).language = language
}

func (moss *Moss) SetMaxLimit(max_limit int) {
	(*moss).max_limit = max_limit
}

func (moss *Moss) SetNoOfMatchingFiles(no_of_matching_files int) {
	(*moss).no_of_matching_files = no_of_matching_files
}

func (moss *Moss) SetExperimental() {
	(*moss).experimental = 1
}

func (moss *Moss) SetComment(comment string) {
	(*moss).comment = comment
}

func (moss *Moss) SetDirectory() {
	(*moss).directory = 1
}

func (moss *Moss) SetBaseFiles(base_files ...string) {
	for _, base_file := range base_files {
		size, err := utils.GetSize(base_file)

		if size > 0 {
			(*moss).base_files = append((*moss).base_files, base_file)

		} else {
			utils.ErrorP(err.Error())
		}
	}
}

func (moss *Moss) AddBaseFile(base_file string) {
	size, err := utils.GetSize(base_file)

	if size > 0 {
		(*moss).base_files = append((*moss).base_files, base_file)

	} else {
		utils.ErrorP(err.Error())
	}
}

func (moss *Moss) AddBaseFilesByWildcard(wildcard string) {
	(*moss).base_files = append((*moss).base_files, utils.GetFilesByWildcard(wildcard)...)
}

func (moss *Moss) SetFiles(files ...string) {
	for _, file := range files {
		size, err := utils.GetSize(file)

		if size > 0 {
			(*moss).files = append((*moss).files, file)

		} else {
			utils.ErrorP(err.Error())
		}
	}
}

func (moss *Moss) AddFile(file string) {
	size, err := utils.GetSize(file)

	if size > 0 {
		(*moss).files = append((*moss).files, file)

	} else {
		utils.ErrorP(err.Error())
	}
}

func (moss *Moss) AddFilesByWildcard(wildcard string) {
	(*moss).files = append((*moss).files, utils.GetFilesByWildcard(wildcard)...)
}

func (moss Moss) UploadFile(connection net.Conn, file string, file_id int) {
	display_name := strings.ReplaceAll(strings.ReplaceAll(file, " ", "_"), "\\", "/")

	file_size, _ := utils.GetSize(file)
	fmt.Fprintf(connection, "file %v %v %v %v\n", file_id, moss.language, file_size, display_name)

	file_data := utils.ReadFile(file)
	fmt.Fprint(connection, file_data)
}

func (moss Moss) SendForReview() string {
	connection, error := net.Dial("tcp", fmt.Sprintf("%v:%v", moss.server, moss.port))

	if error != nil {
		utils.ErrorP(error.Error())
		connection.Close()
	}

	fmt.Fprintf(connection, "moss %v\n", moss.unique_id)
	fmt.Fprintf(connection, "directory %v\n", moss.directory)
	fmt.Fprintf(connection, "X %v\n", moss.experimental)
	fmt.Fprintf(connection, "maxmatches %v\n", moss.max_limit)
	fmt.Fprintf(connection, "show %v\n", moss.no_of_matching_files)
	fmt.Fprintf(connection, "language %v\n", moss.language)

	confirmation, error := bufio.NewReader(connection).ReadString('\n')

	if error != nil || strings.TrimSpace(confirmation) != "yes" {
		connection.Close()
		utils.ErrorP(error.Error())
	} else if strings.HasPrefix(confirmation, "Error") {
		connection.Close()
		utils.ErrorP(confirmation)
	}

	for _, file := range moss.base_files {
		moss.UploadFile(connection, file, 0)
	}

	for index, file := range moss.files {
		moss.UploadFile(connection, file, index+1)
	}

	fmt.Fprintf(connection, "query 0 %v\n", moss.comment)

	response, error := bufio.NewReader(connection).ReadString('\n')

	connection.Close()

	if error != nil {
		utils.ErrorP(error.Error())
	} else if strings.HasPrefix(response, "Error") {
		utils.ErrorP(response)
	}

	return response
}
