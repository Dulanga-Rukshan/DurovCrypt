package DurovCrypt

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/AlecAivazis/survey/v2"
)

// all allowed extensions
var (
	AllExtentions = []string{
		//images files
		".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp", ".tiff", ".svg", ".ico",
		//documents files
		".pdf", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".txt", ".rtf",
		".odt", ".ods", ".odp", ".pages", ".numbers", ".key",
		//archive files
		".zip", ".rar", ".7z", ".tar", ".gz", ".bz2", ".xz", ".z", ".lz", ".lzma",
		//audio files
		".mp3", ".wav", ".ogg", ".flac", ".aac", ".m4a", ".wma", ".opus", ".mid",
		//video files
		".mp4", ".avi", ".mov", ".wmv", ".flv", ".mkv", ".webm", ".mpeg", ".3gp",
		//code files
		".go", ".js", ".py", ".java", ".c", ".cpp", ".h", ".html", ".css", ".php",
		".rb", ".sh", ".pl", ".swift", ".kt", ".ts", ".rs", ".lua",
		//data files
		".json", ".xml", ".csv", ".yaml", ".yml", ".sql", ".db", ".mdb", ".accdb",
		//excutable files
		".exe", ".dll", ".so", ".dylib", ".bin", ".app", ".apk", ".deb", ".rpm",
		//for the decryption
		".drv"}
)

// input prompt for user
func InputPrompt(operation string) (string, error) {
	fmt.Printf("Enter a FilePath/FileName to %s: ", operation)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input),
		err
}

// operation asker function
func DefaultAskInput() (string, error) {
	option := ""
	prompt := &survey.Select{
		Message: "What do you wanna perform in DurovCrypt?",
		Options: []string{"Encrypt", "Decrypt", "Help"},
		Description: func(value string, index int) string {
			if value == "Encrypt" {
				return value

			}
			if value == "Decrypt" {
				return value
			}
			if value == "Help" {
				return value
			}
			return ""
		},
	}
	survey.AskOne(prompt, &option, survey.WithIcons(func(icons *survey.IconSet) {
		icons.Question.Text = ">>>"
		icons.Question.Format = "red"
	}))
	return option,
		nil
}

// get the abs filepath
func AbsPath(filename string) string {
	if filepath.IsAbs(filename) {
		return filename
	}

	absPath, err := filepath.Abs(filename)
	if err != nil {
		return filepath.Join(".", filename)
	}
	return absPath
}

// input file name validation
func IsValidFileName(filename string, checkFunction FileChecker, helpMsg string) error {
	fileName := strings.TrimSpace(filename)

	//if userinput is empty
	if fileName == "" {
		return NewFileError("INVALIDFNAMEERR: Come on dude, enter a filename!!", helpMsg)
	}

	//if filename is more that 255 chracters
	if utf8.RuneCountInString(fileName) > 255 {
		return NewFileError("INVALIDFNAMEERR: You kidding man, filename more than 255 chracters ??", helpMsg)
	}

	//clean the filepath
	cleaned := filepath.Clean(fileName)
	if cleaned != fileName {
		return NewFileError("INVALIDFNAMEERR: file path is wrong!", helpMsg)
	}

	//file info
	fileInfo, err := os.Stat(cleaned)
	if err != nil {
		if os.IsNotExist(err) {
			return NewFileError("INVALIDFNAMEERR: File not exist!", helpMsg)
		}
		return NewFileError("INVALIDFNAMEERR: Can't read file info", helpMsg)
	}

	//base file name
	baseFileName := filepath.Base(cleaned)
	if baseFileName == "." || baseFileName == ".." {
		return NewFileError("INVALIDFNAMEERR: Can't help you dude, file name has invalid chracters.", helpMsg)

	}

	//check for invalid characters in base file name
	for _, r := range baseFileName {
		if r == '/' || r == '\\' || r == ':' || r == '*' ||
			r == '?' || r == '"' || r == '<' || r == '>' || r == '|' {
			return NewFileError("INVALIDFNAMEERR: Can't help you dude, file name has invalid chracters.", helpMsg)
		}
	}

	//checking whether user inputed file is valid for encryption or decryption
	if fileInfo.IsDir() {
		return NewFileError("INVALIDFNAMEERR: You entered a Directory name.", helpMsg)
	}

	//check file size
	if checkFunction.MaxFileSize > 0 && fileInfo.Size() > checkFunction.MaxFileSize {
		return NewFileError("INVALIDFNAMEERR: You can't enter file that higher than 10 MB.", helpMsg)
	}

	//check for file extension
	if len(AllExtentions) > 0 {

		ext := strings.ToLower(filepath.Ext(cleaned))
		validExtension := false
		for _, extensionInSlice := range AllExtentions {
			if strings.ToLower(extensionInSlice) == ext {
				validExtension = true
				break
			}
		}
		if !validExtension {
			return NewFileError("INVALIDFNAMEERR: Invalid file extension.", helpMsg)
		}
	}
	return nil
}

// user input prompt for encrypt file path
func FilePathInput(operation string) (string, error) {
	errorType := "FILEPATHERR: "
	for {
		//ask user for fileName for encrypt
		filePath, err := InputPrompt(operation)
		if err != nil {
			fmt.Printf("%sError reading input: %v\n", errorType, err)
			continue
		}

		//chheck if the user wants to cancel
		if strings.ToUpper(filePath) == "CANCEL" || strings.ToUpper(filePath) == "EXIT" {
			return "", fmt.Errorf("operation cancelled by user")
		}

		//inputed file path checking
		fileCheckerOptions := FileChecker{
			MaxFileSize: 10 * 1024 * 1024, //10 mb file max filesize
			AllowdExt:   AllExtentions,
		}

		if err := IsValidFileName(filePath, fileCheckerOptions, HelpMsg()); err != nil {
			fmt.Printf("\n%v\n", err)
			fmt.Println("\nPlease try again or type 'cancel' to exit.\n")
			continue
		}

		return filePath,
			nil
	}

}

// file reading function
func FileRead(fileName string, operation string) ([]byte, []byte, []byte, error) {

	//absolute file path /home/user/Desktop/file.txt
	absFilePath := AbsPath(fileName)

	//opening a file
	filename, err := os.OpenFile(absFilePath, os.O_RDONLY, 0644)
	if err != nil {
		return nil,
			nil,
			nil,
			fmt.Errorf("FILEREADERR: Error opening the file. \n\n%w", err)
	}

	//closing file
	defer filename.Close()

	switch strings.ToUpper(operation) {
	case "ENCRYPT", "E":
		//fileExtension .txt
		fileExtension := filepath.Ext(absFilePath)

		if strings.ToUpper(fileExtension) == ".DRV" {
			return nil,
				nil,
				nil,
				fmt.Errorf("FILEWRITEERR: File type is not valid for encrypt.")
		}

		//read the file for encrypt
		data, err := io.ReadAll(filename)

		//if file has 0 data file is empty error
		if len(data) == 0 {
			return nil,
				nil,
				nil,
				fmt.Errorf("FILEREADERR: File is empty.")
		}

		if err != nil {
			return nil,
				nil,
				nil,
				fmt.Errorf("FILEREADERR: Error reading the file. \n\n%w", err)
		}

		return data,
			nil,
			nil,
			nil

	case "DECRYPT", "D":
		//fileExtension .txt
		fileExtension := filepath.Ext(absFilePath)

		if strings.ToUpper(fileExtension) != ".DRV" {
			return nil,
				nil,
				nil,
				fmt.Errorf("FILEWRITEERR: File type is not valid for Decrypt Enter only Durocrypt encrypted file.")

		}

		//read from the encrypted file
		data, err := io.ReadAll(filename)

		//if file len < 0
		if len(data) == 0 {
			return nil,
				nil,
				nil,
				fmt.Errorf("FILEREADERR: File is empty for decrypt.")
		}

		if err != nil {
			return nil,
				nil,
				nil,
				fmt.Errorf("READERR: Error reading the file for decrypt. \n\n%w", err)
		}

		// Calculate positions
		ciphertextLen := len(data) - (saltSize + nonceSize)

		if ciphertextLen <= 0 {
			return nil, nil, nil, fmt.Errorf("INVALIDFILEERR: File too small")
		}

		//reading cipher text
		ciphertext := data[:ciphertextLen]

		//reading the salt
		salt := data[ciphertextLen : ciphertextLen+saltSize]

		//reading the nonce
		nonce := data[ciphertextLen+saltSize:]

		return ciphertext,
			salt,
			nonce,
			nil

	}
	return nil,
		nil,
		nil,
		nil
}

func FileWrite(data []byte, fileName string, salt []byte, nonce []byte, operation string) (string, error) {
	if len(data) == 0 {
		return "",
			fmt.Errorf("FILEWRITEERR: cipherText is empty.")
	}

	fileCheckerOptions := FileChecker{
		MaxFileSize: 10 * 1024 * 1024, //10 mb file max filesize
		AllowdExt:   AllExtentions,
	}
	if err := IsValidFileName(fileName, fileCheckerOptions, HelpMsg()); err != nil {
		return "",
			fmt.Errorf("FILEWRITEERR: %w", err)
	}

	//absolute file path /home/user/Desktop/file.txt
	absFilePath := AbsPath(fileName)

	//file directory name  desktop
	fileDirectory := filepath.Dir(absFilePath)

	//filebasename file.txt
	fileBaseName := filepath.Base(absFilePath)

	//fileExtension .txt
	fileExtension := filepath.Ext(absFilePath)

	//get the filename only  without extension
	fileNameOnly := strings.TrimSuffix(fileBaseName, fileExtension)

	switch strings.ToUpper(operation) {
	case "ENCRYPT", "E":
		if strings.ToUpper(fileExtension) == "DRV" {
			return "",
				fmt.Errorf("FILEWRITEERR: File is already been encrypted by DurovCrypt.")
		}
		if len(salt) < 32 {
			return "",
				fmt.Errorf("FILEWRITEERR: Salt is less than 32 bytes.")
		}

		if len(nonce) == 0 {
			return "",
				fmt.Errorf("FILEWRITEERR: Nonce is empty.")
		}

		//preallocation the memory by specifying combinedata and add salt,nonce and data separatly
		combinedData := make([]byte, len(data)+len(salt)+len(nonce))
		copy(combinedData[:len(data)], data)                   //data first
		copy(combinedData[len(data):len(data)+saltSize], salt) //salt next
		copy(combinedData[len(data)+saltSize:], nonce)         //nonce last

		//specifying file extension for encrypted file
		drvExtension := ".drv"

		//go to the basefilepath and create the encrypted file there
		fileOutputPath := filepath.Join(fileDirectory, fileNameOnly+fileExtension+drvExtension)

		//writing to a file with fileOutputPath
		err := os.WriteFile(fileOutputPath, combinedData, 0600)
		if err != nil {
			return "",
				fmt.Errorf("FILEWRITEERR: Can't write Data to file! %w\n", err)
		}

		//returning the  success message.
		return fmt.Sprintf("File saved as %v!", fileOutputPath),
			nil

	case "DECRYPT", "D":
		if strings.ToUpper(fileExtension) != ".DRV" {
			return "",
				fmt.Errorf("FILEWRITEERR: File extension is not drv ")
		}

		//this get the original file extension file.txt.drv get the txt

		//go to the basefilepath and create the decrypted file there
		fileOutputPath := filepath.Join(fileDirectory, "Decrypt"+fileNameOnly)

		//output filename
		outputFilename := fmt.Sprintf("%v", fileOutputPath)

		//wrting to the file name
		err := os.WriteFile(outputFilename, data, 0644)
		if err != nil {
			return "",
				fmt.Errorf("FILEWRITEERR: Can't write Data to file! %w\n", err)
		}
		return fmt.Sprintf("File saved as %v!", outputFilename),
			nil
	}
	return "",
		nil
}
