package DurovCrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"

	"golang.org/x/crypto/argon2"
)

const (
	saltSize  = 32
	nonceSize = 12

	//argon2
	ArgonIterations = 3
	ArgonMemory     = 64 * 1024 //64MB
	ArgonThreads    = 4
	ArgonKeyLength  = 32 //256bit key
)

// aes256 salt generator
func DataKey() []byte {
	datakey := make([]byte, saltSize)
	if _, err := io.ReadFull(rand.Reader, datakey); err != nil {
		fmt.Println("Err: Aragon2 salt generate failed: %w", err)
	}
	return datakey
}

// KeyGen generates consistent keys for both encryption and decryption
func KeyGen(password string, salt []byte) ([]byte, []byte, error) {
	if len(password) == 0 {
		return nil, nil, errors.New("empty password")
	}
	if len(salt) < 8 {
		return nil, nil, errors.New("salt too short")
	}

	derivedKey := argon2.IDKey(
		[]byte(password),
		salt,
		ArgonIterations,
		ArgonMemory,
		ArgonThreads,
		ArgonKeyLength,
	)

	return derivedKey, salt, nil
}

// key generation
func KeyGenerator(password string, salt []byte) ([]byte, []byte, error) {

	derivedKey, salt, err := KeyGen(password, salt)
	if err != nil {
		return nil, nil, fmt.Errorf("KEYGENERR: %v", err)
	}

	return derivedKey, salt, nil
}

// encrypt data
func Encrypt(password string, plaintext []byte) ([]byte, []byte, []byte, error) {
	//ERRORTYPE
	errorType := "ENCRYPTERR: "

	//salt generation
	salt := DataKey()

	//key generation
	derivedKey, salt, err := KeyGenerator(password, salt)
	MainErr("", err)

	//aes
	AES, err := aes.NewCipher(derivedKey)
	MainErr(errorType, err)

	//gcm
	gcm, err := cipher.NewGCM(AES)
	MainErr(errorType, err)

	//gcm nonce or iv generate
	nonce := make([]byte, nonceSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println("%v GCM nonce for encrypt err: %v", errorType, err)

		return nil,
			nil,
			nil,
			err
	}

	//gcm encrypt the plaintext
	cipherText := gcm.Seal(nil, nonce, plaintext, nil)

	return cipherText,
		salt,
		nonce,
		nil
}

// decrypt data
func Decrypt(password string, saltFromFile []byte, nonce []byte, cipherText []byte) (string, error) {
	//ERRORTYPE
	errorType := "DECRYPTERR: "

	derivedKey, _, err := KeyGenerator(password, saltFromFile)

	//key validation
	if len(derivedKey) != 32 {
		return "", fmt.Errorf("%vInvalid key size", errorType)
	}

	//nonce validation
	if len(nonce) != nonceSize {
		return "", fmt.Errorf("%vInvalid nonce size", errorType)
	}

	//aes key generate
	AES, err := aes.NewCipher(derivedKey)
	MainErr(errorType, err)

	//gcm
	gcm, err := cipher.NewGCM(AES)
	MainErr(errorType, err)

	//checking ciphertext size
	if len(cipherText) < gcm.Overhead() {
		return "", fmt.Errorf("%vCiphertext too short", errorType)
	}
	//gcm encrypt the plaintext
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "",
			fmt.Errorf("%v decryption failed: %w", errorType, err)
	}

	return string(plainText),
		nil
}
