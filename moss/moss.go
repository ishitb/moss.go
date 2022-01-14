package moss

type Moss struct {
	Languages []string
	server    string
	port      int
	uniqueId  string

	// Flags
	language   string
	directory  string
	base_files []string
	max_limit  int

	// Args
	files []string
}

func NewMoss(uniqueId string) Moss {
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
		server:     "moss.stanford.edu",
		port:       7690,
		uniqueId:   uniqueId,
		language:   "c",
		base_files: []string{},
		max_limit:  10,
		files:      []string{},
	}

	return moss
}

func (moss *Moss) ChangeLanguage(language string) {
	(*moss).language = language
}

func (moss *Moss) ChangeMaxLimit(max_limit int) {
	(*moss).max_limit = max_limit
}

func (moss *Moss) SetDirectory(directory string) {
	(*moss).directory = directory
}

func (moss *Moss) SetBaseFiles(base_files ...string) {
	(*moss).base_files = base_files
}

func (moss *Moss) AddBaseFile(base_file string) {
	(*moss).base_files = append((*moss).base_files, base_file)
}

func (moss *Moss) SetFiles(files ...string) {
	(*moss).files = files
}

func (moss *Moss) AddFile(file string) {
	(*moss).files = append((*moss).files, file)
}
