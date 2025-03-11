package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"

	"github.com/WanderningMaster/kv/internal/assert"
)

func GenKey() []byte {
	key := make([]byte, 32)

	_, err := rand.Read(key)
	assert.Assert(err)

	return key
}

func Enc(input string, keyStr string) []byte {
	key, _ := hex.DecodeString(keyStr)
	b := []byte(input)

	block, err := aes.NewCipher(key)
	assert.Assert(err)

	aesGCM, err := cipher.NewGCM(block)
	assert.Assert(err)

	nonce := make([]byte, aesGCM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	assert.Assert(err)

	return aesGCM.Seal(nonce, nonce, b, nil)
}

func Dec(input string, keyStr string) []byte {
	key, _ := hex.DecodeString(keyStr)
	b := []byte(input)

	block, err := aes.NewCipher(key)
	assert.Assert(err)

	aesGCM, err := cipher.NewGCM(block)
	assert.Assert(err)

	nonceSize := aesGCM.NonceSize()
	nonce, cipherText := b[:nonceSize], b[nonceSize:]

	plain, err := aesGCM.Open(nil, nonce, cipherText, nil)
	assert.Assert(err)

	return plain
}
