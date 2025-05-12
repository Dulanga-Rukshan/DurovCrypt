package DurovCrypt

func ShowHelp() string {
	help := `
 DurovCrypt - A go file password base file encryptor

USAGE:
  DurovCrypt [command] [options]

COMMANDS:
  encrypt (e)-----> Encrypt a file
  decrypt (d)-----> Decrypt a file
  help (h)   -----> Show this help message

OPTIONS:
  -f, --file     ---> Specify file to process

EXAMPLES:
  1. Encrypt a file:
     DurovCrypt e -f secret.txt
     (You'll be prompted for password)

  2. Decrypt a file:
     DurovCrypt d -f secret.txt.enc
     (Password prompt will appear)

  3. Quick help:
     DurovCrypt help

SECURITY TIPS:
  • Use strong passwords (mix letters, numbers, symbols)
  • Never share your password
  • Keep backups of original files
  • Delete unencrypted files after encryption

VERSION: 1.0.0
`
	return help
}

func HelpMsg() string {
	help := `COMMANDS:
		encrypt (e)    - Encrypt a file
		decrypt (d)    - Decrypt a file
		help (h)       - Show this help message`

	return help
}
