package hash

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"strconv"
)

// AdjustKey ensures the key length is suitable for AES (16, 24, or 32 bytes)
func AdjustKey(key string) []byte {
	k := []byte(key)
	switch len(k) {
	case 16, 24, 32:
		return k
	default:
		if len(k) > 32 {
			return k[:32]
		}
		return append(k, make([]byte, 32-len(k))...)
	}
}

// HashID generates a hash from an auto-incremented ID using a secret key
func HashID(id int, secret string) (string, error) {
	key := AdjustKey(secret)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(strconv.Itoa(id)), nil)
	return hex.EncodeToString(ciphertext), nil
}

// DecryptID decrypts a hash back into an ID using the same secret key
func DecryptID(hash string, secret string) (int, error) {
	key := AdjustKey(secret)
	block, err := aes.NewCipher(key)
	if err != nil {
		return 0, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return 0, err
	}

	ciphertext, err := hex.DecodeString(hash)
	if err != nil {
		return 0, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return 0, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return 0, err
	}

	id, err := strconv.Atoi(string(plaintext))
	if err != nil {
		return 0, err
	}

	return id, nil
}
