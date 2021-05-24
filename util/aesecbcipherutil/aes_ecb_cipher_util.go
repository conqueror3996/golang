package aesecbcipherutil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

// PadByPkcs5 ...
func PadByPkcs5(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// UnPadByPkcs5 ...
func UnPadByPkcs5(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

// DecryptAESECBModeV1 ...
func DecryptAESECBModeV1(crypted, key []byte) (origData []byte, errres error) {
	defer func() {
		if rec := recover(); rec != nil {
			errres = fmt.Errorf("panic recovered: %v", rec)
			return
		}
	}()

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := NewECBDecrypter(block)
	origData = make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = UnPadByPkcs5(origData)
	return origData, nil
}

// EncryptAESECBModeV1 ...
func EncryptAESECBModeV1(src, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	ecb := NewECBEncrypter(block)
	content := []byte(src)
	content = PadByPkcs5(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)
	return crypted, nil
}

// EncryptAESECBModeV2 ...
func EncryptAESECBModeV2(src, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ecb := NewECBEncrypter(block)
	content := src
	content = PadByPkcs5(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)
	return crypted, nil
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

// NewECBEncrypter ...
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}

func (x *ecbEncrypter) BlockSize() int {
	return x.blockSize
}

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto / cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto / cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}

func (x *ecbDecrypter) BlockSize() int {
	return x.blockSize
}

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto / cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto / cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
