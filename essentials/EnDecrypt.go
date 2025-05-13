package DurovCrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"runtime"

	"golang.org/x/crypto/argon2"
)

// key gernerating for encrypt and decrypt
func (masterKey *Aragon2Key) Generate() ([]byte, error) {
	if masterKey == nil {
		return nil, errors.New("Err: Argon key not found!")
	}
	if len(masterKey.Password) == 0 {
		return nil, errors.New("Err: Password is empty!")
	}
	if len(masterKey.Salt) < 8 {
		return nil, errors.New("Err: Salt is weak!")
	}
	if masterKey.Iteration < 1 {
		return nil, errors.New("Err: Iteration count is low!")
	}

	derivedKey := argon2.IDKey(
		masterKey.Password,
		masterKey.Salt,
		masterKey.Iteration,
		masterKey.MemSize,
		masterKey.Threads,
		uint32(masterKey.KeyLength),
	)
	return derivedKey, nil
}

// key generation
func KeyGen(password string) ([]byte, error) {

	//key generation
	salt := DataKey()
	NewKey := Aragon2Key{
		Password:  []byte(password),
		Salt:      salt,
		Iteration: uint32(3),
		MemSize:   uint32(64 * 1024),
		Threads:   uint8(runtime.NumCPU()),
		KeyLength: 32,
	}

	derivedKey, err6 := NewKey.Generate()
	if err6 != nil {
		return nil, fmt.Errorf("\nERROR: %v\n", err6)
	}

	return derivedKey, nil
}

// aes256 Key
func DataKey() []byte {
	datakey := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, datakey); err != nil {
		fmt.Println("Err: Aragon salt generate failed: %w", err)
	}
	return datakey
}

// encrypt data
func Encrypt(Key []byte, plaintext string) (string, error) {
	//aes
	AES, err := aes.NewCipher(Key)
	MainErr(err)

	//gcm
	gcm, err := cipher.NewGCM(AES)
	MainErr(err)

	//gcm nonce or iv generate
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Printf("Err: GCM nonce for encrypt err: %v", err)
		return "", err
	}

	//gcm encrypt the plaintext
	cipherText := gcm.Seal(nil, nonce, []byte(plaintext), nil)

	return hex.EncodeToString(cipherText), nil
}

// decrypt data
func Decrypt(Key []byte, cipherText string) (string, error) {
	//
	ciphertext, _ := hex.DecodeString(cipherText)

	//aes
	AES, err := aes.NewCipher(Key)
	MainErr(err)

	//gcm
	gcm, err := cipher.NewGCM(AES)
	MainErr(err)

	//nonceSize
	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return "", fmt.Errorf("Err: Cipher text is shorter than nonce.")

	}

	//nonce and message spliting from cipherText
	nonce, message := ciphertext[:nonceSize], ciphertext[nonceSize:]

	//decrypting the ciphertext
	plainText, err := gcm.Open(nil, nonce, message, nil)
	if err != nil {
		return "", fmt.Errorf("Err: decryption failed: %w", err)
	}
	return string(plainText), nil
}
