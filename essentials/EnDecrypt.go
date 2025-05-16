package DurovCrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"runtime"
	"strings"

	"golang.org/x/crypto/argon2"
)

// aes256 salt generator
func DataKey() []byte {
	datakey := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, datakey); err != nil {
		fmt.Println("Err: Aragon2 salt generate failed: %w", err)
	}
	return datakey
}

// key gernerating for encrypt and decrypt
func (masterKey *Aragon2Key) Generate() ([]byte, []byte, error) {
	if masterKey == nil {
		return nil, nil, errors.New("Err: Argon key not found!")
	}
	if len(masterKey.Password) == 0 {
		return nil, nil, errors.New("Err: Password is empty!")
	}
	if len(masterKey.Salt) < 8 {
		return nil, nil, errors.New("Err: Salt is weak!")
	}
	if masterKey.Iteration < 1 {
		return nil, nil, errors.New("Err: Iteration count is low!")
	}

	derivedKey := argon2.IDKey(
		masterKey.Password,
		masterKey.Salt,
		masterKey.Iteration,
		masterKey.MemSize,
		masterKey.Threads,
		uint32(masterKey.KeyLength),
	)

	return derivedKey, masterKey.Salt, nil
}

// key generation
func KeyGen(password string, operation string, salt []byte) ([]byte, []byte, error) {

	switch strings.ToUpper(operation) {
	case "ENCRYPT", "E":

		NewKey := Aragon2Key{
			Password:  []byte(password),
			Salt:      salt,
			Iteration: uint32(3),
			MemSize:   uint32(64 * 1024),
			Threads:   uint8(runtime.NumCPU()),
			KeyLength: 32,
		}

		derivedKey, derivedSalt, err6 := NewKey.Generate()
		if err6 != nil {
			return nil, nil, fmt.Errorf("\nERROR: %v\n", err6)
		}

		return derivedKey, derivedSalt, nil

	case "DECRYPT", "D":
		//key generation
		NewKey := Aragon2Key{
			Password:  []byte(password),
			Salt:      salt,
			Iteration: uint32(3),
			MemSize:   uint32(64 * 1024),
			Threads:   uint8(runtime.NumCPU()),
			KeyLength: 32,
		}
		derivedKey, derivedSalt, err6 := NewKey.Generate()
		if err6 != nil {
			return nil, nil, fmt.Errorf("\nERROR: %v\n", err6)
		}

		return derivedKey, derivedSalt, nil
	}

	return nil, nil, nil
}

// encrypt data
func Encrypt(password string, operation string, plaintext string) (string, string, error) {
	//salt generation
	salt := DataKey()

	//key generation
	derivedKey, salt, err5 := KeyGen(password, operation, salt)
	MainErr(err5)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	fmt.Print(b64Salt)

	//aes
	AES, err := aes.NewCipher(derivedKey)
	MainErr(err)

	//gcm
	gcm, err := cipher.NewGCM(AES)
	MainErr(err)

	//gcm nonce or iv generate
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Printf("Err: GCM nonce for encrypt err: %v", err)
		return "", "", err
	}

	//gcm encrypt the plaintext
	cipherText := gcm.Seal(nil, nonce, []byte(plaintext), nil)

	return hex.EncodeToString(cipherText), b64Salt, nil
}

// decrypt data
func Decrypt(Key []byte, cipherText string) (string, error) {
	//
	ciphertext, err := hex.DecodeString(cipherText)
	MainErr(err)

	//aes
	AES, err := aes.NewCipher(Key)
	MainErr(err)

	if err != nil {
		return "", fmt.Errorf("Err: decryption failed: %w", err)
	}
	return string(plainText), nil
}
