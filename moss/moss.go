package moss

import "github.com/ishitb/moss-client-go/utils"

type Moss struct {
	Languages []string
	server    string
	port      int
	unique_id string

	// Flags
	language             string
	directory            string
	base_files           []string
	max_limit            int
	comment              string
	no_of_matching_files int
	experimental         bool

	// Args
	files []string
}

func NewMoss(unique_id string) Moss {
	moss := Moss{
		Languages: []string{
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
		},
		server:               "moss.stanford.edu",
		port:                 7690,
		unique_id:            unique_id,
		language:             "c",
		base_files:           []string{},
		max_limit:            10,
		files:                []string{},
		no_of_matching_files: 250,
		experimental:         false,
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

func (moss *Moss) SetExperimental(experimental bool) {
	(*moss).experimental = experimental
}

func (moss *Moss) SetComment(comment string) {
	(*moss).comment = comment
}

func (moss *Moss) SetDirectory(directory string) {
	(*moss).directory = directory
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
