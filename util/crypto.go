package util

import (
	"crypto/aes"
	"crypto/rand"
	"encoding/base32"
)

const randomBase = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Salt ==> generate a random salt
func Salt(len uint) string {
	b := make([]byte, len)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return base32.HexEncoding.EncodeToString(b)
	// return Byte2Str(b...)
}

// Encrypt ==> encrypt
func Encrypt(key []byte, src []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	dst := make([]byte, 32)
	block.Encrypt(dst, src)
	return string(dst)
}

// Encrypt16 ==> encrypt 16
func Encrypt16(key []byte, src []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	dst := make([]byte, 16)
	block.Encrypt(dst, src)
	return string(dst)
}

// Encrypt32 ==> encrypt 32
func Encrypt32(key []byte, src []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	dst := make([]byte, 32)
	block.Encrypt(dst, src)
	return string(dst)
}

// Encrypt64 ==> encrypt 64
func Encrypt64(key []byte, src []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	dst := make([]byte, 64)
	block.Encrypt(dst, src)
	return string(dst)
}

// Encrypt128 ==> encrypt 128
func Encrypt128(key []byte, src []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	dst := make([]byte, 128)
	block.Encrypt(dst, src)
	return string(dst)
}
