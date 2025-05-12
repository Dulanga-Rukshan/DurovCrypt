package DurovCrypt

type PasswordPolicyCheck struct {
	MinLength           int
	MaxLength           int
	RequireUpper        bool
	RequireLower        bool
	RequireSymbol       bool
	RequireNumber       bool
	DenyWhiteSpace      bool
	AllowedSpecialChars string
}

type Aragon2Key struct {
	Password  []byte
	Salt      []byte
	Iteration uint32
	MemSize   uint32
	Threads   uint8
	KeyLength uint32
}

type FileChecker struct {
	MaxFileSize int64
	AllowdExt   []string
}

type ErrorHandeling struct {
	Message string
	HelpMsg string
}
