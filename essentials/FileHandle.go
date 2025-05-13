package DurovCrypt

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
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
	return strings.TrimSpace(input), err
}

// operation asker function
func DefaultAskInput(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

// input file name validation
func IsValidFileName(filename string, checkFunction FileChecker, helpMsg string) error {
	fileName := strings.TrimSpace(filename)

	//if userinput is empty
	if fileName == "" {
		return NewFileError("Come on dude, enter a filename!!", helpMsg)
	}

	//if filename is more that 255 chracters
	if utf8.RuneCountInString(fileName) > 255 {
		return NewFileError("You kidding man, filename more than 255 chracters ??", helpMsg)
	}

	//clean the filepath
	cleaned := filepath.Clean(fileName)
	if cleaned != fileName {
		return NewFileError("file path is wrong!", helpMsg)
	}

	//file info
	fileInfo, err := os.Stat(cleaned)
	if err != nil {
		if os.IsNotExist(err) {
			return NewFileError("File not exist!", helpMsg)
		}
		return NewFileError("Can't read file info", helpMsg)
	}

	//base file name
	baseFileName := filepath.Base(cleaned)
	if baseFileName == "." || baseFileName == ".." {
		return NewFileError("Can't help you dude, file name has invalid chracters.", helpMsg)

	}

	//check for invalid characters in base file name
	for _, r := range baseFileName {
		if r == '/' || r == '\\' || r == ':' || r == '*' ||
			r == '?' || r == '"' || r == '<' || r == '>' || r == '|' {
			return NewFileError("Can't help you dude, file name has invalid chracters.", helpMsg)
		}
	}

	//checking whether user inputed file is valid for encryption or decryption
	if fileInfo.IsDir() {
		return NewFileError("You entered a Directory name.", helpMsg)
	}

	//check file size
	if checkFunction.MaxFileSize > 0 && fileInfo.Size() > checkFunction.MaxFileSize {
		return NewFileError("You can't enter file that higher than 10 MB.", helpMsg)
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
			return NewFileError("Invalid file extension.", helpMsg)
		}
	}
	return nil
}

// user input prompt for encrypt file path
func FilePathInput(operation string) (string, error) {

	//ask user for fileName for encrypt
	filePath, err := InputPrompt(operation)
	if err != nil {
		return "", err
	}

	//inputed file path checking
	fileCheckerOptions := FileChecker{
		MaxFileSize: 10 * 1024 * 1024, //10 mb file max filesize
		AllowdExt:   AllExtentions,
	}

	if err := IsValidFileName(filePath, fileCheckerOptions, HelpMsg()); err != nil {
		return "", fmt.Errorf("%v", err)
	}

	return filePath, nil
}

// file reading function
func FileRead(fileName string) ([]byte, error) {
	//opening a file
	filename, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Error opening the file. \n\n%w", err)
	}

	//closing file
	defer filename.Close()

	data, err := io.ReadAll(filename)
	if len(data) < 0 {
		return nil, fmt.Errorf("File is empty.")
	}
	if err != nil {
		return nil, fmt.Errorf("Error reading the file. \n\n%w", err)
	}
	return data, nil
}

func FileWrite(cipherText []byte, fileName string) (string, error) {
	if len(cipherText) < 0 {
		return "", fmt.Errorf("cipherText is empty.")
	}

	fileCheckerOptions := FileChecker{
		MaxFileSize: 10 * 1024 * 1024, //10 mb file max filesize
		AllowdExt:   AllExtentions,
	}
	if err := IsValidFileName(fileName, fileCheckerOptions, HelpMsg()); err != nil {
		return "", fmt.Errorf("Error: %w", err)
	}
	basepath := filepath.Base(fileName)
	rmextension := filepath.Ext(fileName)
	filename := basepath[:len(basepath)-len(rmextension)]

	extension := "drv"
	err := os.WriteFile(filename+"."+extension, cipherText, 0644)
	if err != nil {
		return "", fmt.Errorf("Can't write Data to file! %w\n", err)
	}
	return fmt.Sprintf(" file saved as %v.%v!", filename, extension), nil
}
package DurovCrypt

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
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
	return strings.TrimSpace(input), err
}

// operation asker function
func DefaultAskInput(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

// input file name validation
func IsValidFileName(filename string, checkFunction FileChecker, helpMsg string) error {
	fileName := strings.TrimSpace(filename)

	//if userinput is empty
	if fileName == "" {
		return NewFileError("Come on dude, enter a filename!!", helpMsg)
	}

	//if filename is more that 255 chracters
	if utf8.RuneCountInString(fileName) > 255 {
		return NewFileError("You kidding man, filename more than 255 chracters ??", helpMsg)
	}

	//clean the filepath
	cleaned := filepath.Clean(fileName)
	if cleaned != fileName {
		return NewFileError("file path is wrong!", helpMsg)
	}

	//file info
	fileInfo, err := os.Stat(cleaned)
	if err != nil {
		if os.IsNotExist(err) {
			return NewFileError("File not exist!", helpMsg)
		}
		return NewFileError("Can't read file info", helpMsg)
	}

	//base file name
	baseFileName := filepath.Base(cleaned)
	if baseFileName == "." || baseFileName == ".." {
		return NewFileError("Can't help you dude, file name has invalid chracters.", helpMsg)

	}

	//check for invalid characters in base file name
	for _, r := range baseFileName {
		if r == '/' || r == '\\' || r == ':' || r == '*' ||
			r == '?' || r == '"' || r == '<' || r == '>' || r == '|' {
			return NewFileError("Can't help you dude, file name has invalid chracters.", helpMsg)
		}
	}

	//checking whether user inputed file is valid for encryption or decryption
	if fileInfo.IsDir() {
		return NewFileError("You entered a Directory name.", helpMsg)
	}

	//check file size
	if checkFunction.MaxFileSize > 0 && fileInfo.Size() > checkFunction.MaxFileSize {
		return NewFileError("You can't enter file that higher than 10 MB.", helpMsg)
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
			return NewFileError("Invalid file extension.", helpMsg)
		}
	}
	return nil
}

// user input prompt for encrypt file path
func FilePathInput(operation string) (string, error) {

	//ask user for fileName for encrypt
	filePath, err := InputPrompt(operation)
	if err != nil {
		return "", err
	}

	//inputed file path checking
	fileCheckerOptions := FileChecker{
		MaxFileSize: 10 * 1024 * 1024, //10 mb file max filesize
		AllowdExt:   AllExtentions,
	}

	if err := IsValidFileName(filePath, fileCheckerOptions, HelpMsg()); err != nil {
		return "", fmt.Errorf("%v", err)
	}

	return filePath, nil
}

// file reading function
func FileRead(fileName string) ([]byte, error) {
	//opening a file
	filename, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Error opening the file. \n\n%w", err)
	}

	//closing file
	defer filename.Close()

	data, err := io.ReadAll(filename)
	if len(data) < 0 {
		return nil, fmt.Errorf("File is empty.")
	}
	if err != nil {
		return nil, fmt.Errorf("Error reading the file. \n\n%w", err)
	}
	return data, nil
}

func FileWrite(cipherText []byte, fileName string) (string, error) {
	if len(cipherText) < 0 {
		return "", fmt.Errorf("cipherText is empty.")
	}

	fileCheckerOptions := FileChecker{
		MaxFileSize: 10 * 1024 * 1024, //10 mb file max filesize
		AllowdExt:   AllExtentions,
	}
	if err := IsValidFileName(fileName, fileCheckerOptions, HelpMsg()); err != nil {
		return "", fmt.Errorf("Error: %w", err)
	}
	basepath := filepath.Base(fileName)
	rmextension := filepath.Ext(fileName)
	filename := basepath[:len(basepath)-len(rmextension)]

	extension := "drv"
	err := os.WriteFile(filename+"."+extension, cipherText, 0644)
	if err != nil {
		return "", fmt.Errorf("Can't write Data to file! %w\n", err)
	}
	return fmt.Sprintf(" file saved as %v.%v!", filename, extension), nil
}
