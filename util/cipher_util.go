package util

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"strings"
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

// // PadByPkcs7 ...
// func PadByPkcs7(data []byte, blockSize int) []byte {
// 	padSize := aes.BlockSize
// 	if len(data)%aes.BlockSize != 0 {
// 		padSize = aes.BlockSize - (len(data))%aes.BlockSize
// 	}

// 	pad := bytes.Repeat([]byte{byte(padSize)}, padSize)
// 	return append(data, pad...)
// }

// // UnPadByPkcs7 ...
// func UnPadByPkcs7(data []byte) []byte {
// 	padSize := int(data[len(data)-1])
// 	return data[:len(data)-padSize]
// }

// EncryptAESCBCModeV1 ...
func EncryptAESCBCModeV1(iv []byte, key []byte, plainText []byte) ([]byte, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/cipher_util.go : EncryptAESCBCModeV1 : func EncryptAESCBCModeV1(iv []byte, key []byte, plainText []byte) ([]byte, error) start.")))

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	paddedPlainText := PadByPkcs5(plainText, block.BlockSize())

	cipherText := make([]byte, len(paddedPlainText))

	mode := cipher.NewCBCEncrypter(block, iv)

	mode.CryptBlocks(cipherText, paddedPlainText)

	return cipherText, nil
}

// DecryptAESCBCModeV1 ...
func DecryptAESCBCModeV1(iv []byte, key []byte, cipherText []byte) ([]byte, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/cipher_util.go : DecryptAESCBCModeV1 : func DecryptAESCBCModeV1(iv []byte, key []byte, cipherText []byte) ([]byte, error) start.")))

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// if len(cipherText) < aes.BlockSize {
	// 	return nil, errors.New("cipher text must be longer than blocksize")
	// } else if len(cipherText)%aes.BlockSize != 0 {
	// 	return nil, errors.New("cipher text must be multiple of blocksize")
	// }

	plainText := make([]byte, len(cipherText))

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(plainText, cipherText)

	unpaddedPlainText := UnPadByPkcs5(plainText)

	return unpaddedPlainText, nil
}

// EncryptAESCBCModeV2 ...
func EncryptAESCBCModeV2(key []byte, plainText []byte) ([]byte, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/cipher_util.go : EncryptAESCBCModeV2 : func EncryptAESCBCModeV2(key []byte, plainText []byte) ([]byte, error) start.")))

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	paddedPlainText := PadByPkcs5(plainText, block.BlockSize())

	cipherText := make([]byte, len(paddedPlainText))

	iv := make([]byte, aes.BlockSize)
	_, err = rand.Read(iv)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)

	mode.CryptBlocks(cipherText, paddedPlainText)

	cipherText = append(iv, cipherText...)

	return cipherText, nil
}

// DecryptAESCBCModeV2 ...
func DecryptAESCBCModeV2(key []byte, cipherText []byte) ([]byte, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/cipher_util.go : DecryptAESCBCModeV2 : func DecryptAESCBCModeV2(key []byte, cipherText []byte) ([]byte, error) start.")))

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// if len(cipherText) < aes.BlockSize {
	// 	return nil, errors.New("cipher text must be longer than blocksize")
	// } else if len(cipherText)%aes.BlockSize != 0 {
	// 	return nil, errors.New("cipher text must be multiple of blocksize")
	// }

	iv := cipherText[:aes.BlockSize]

	plainText := make([]byte, len(cipherText[aes.BlockSize:]))

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(plainText, cipherText[aes.BlockSize:])

	unpaddedPlainText := UnPadByPkcs5(plainText)

	return unpaddedPlainText, nil
}

// DecryptAESCBCModeV3 ...
func DecryptAESCBCModeV3(cipherText string) (string, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/cipher_util.go : DecryptAESCBCModeV3 : func DecryptAESCBCModeV3(cipherText string) (string, error) start.")))

	iv, err := HexDecodeStringV1("2a296367833c49e081a20501a3fb7226")
	if err != nil {
		return "", err
	}

	key, err := HexDecodeStringV1("54598ba3648847beb4dc4ca73ee60fcccfa372581c6f495fb9d740c0b76e2b93")
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	base64StdEncodingDecodeString, err := Base64StdEncodingDecodeStringV1(cipherText)
	if err != nil {
		return "", err
	}

	plainText := make([]byte, len(base64StdEncodingDecodeString))

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(plainText, base64StdEncodingDecodeString)

	unpaddedPlainText := UnPadByPkcs5(plainText)

	return string(unpaddedPlainText), nil
}

// DecryptAESCBCModeV4 ...
func DecryptAESCBCModeV4(cipherText string) ([]byte, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/cipher_util.go : DecryptAESCBCModeV4 : func DecryptAESCBCModeV4(cipherText string) ([]byte, error) start.")))

	decx, err := DecryptAESCBCModeV3(cipherText)
	if err != nil {
		return nil, err
	}

	decb, err := HexDecodeStringV1(decx)
	if err != nil {
		return nil, err
	}

	return decb, nil
}

// DecryptAESCBCModeV4p1 ...
func DecryptAESCBCModeV4p1(cipherText string) (string, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/cipher_util.go : DecryptAESCBCModeV4p1 : func DecryptAESCBCModeV4p1(cipherText string) (string, error) start.")))

	decx, err := DecryptAESCBCModeV3(cipherText)
	if err != nil {
		return "", err
	}

	return decx, nil
}

// DecryptAESCBCModeV5 ...
func DecryptAESCBCModeV5(cipherText1 string, cipherText2 string) ([]byte, []byte, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/cipher_util.go : DecryptAESCBCModeV5 : func DecryptAESCBCModeV5(cipherText1 string, cipherText2 string) ([]byte, []byte, error) start.")))

	dec1, err := DecryptAESCBCModeV4(cipherText1)
	if err != nil {
		return nil, nil, err
	}

	dec2, err := DecryptAESCBCModeV4(cipherText2)
	if err != nil {
		return nil, nil, err
	}

	return dec1, dec2, nil
}

// EncryptAESCBCModeV6 ...
func EncryptAESCBCModeV6(cipherIV string, cipherKey string, plainText string) ([]byte, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/cipher_util.go : EncryptAESCBCModeV6 : func EncryptAESCBCModeV6(cipherIV string, cipherKey string, plainText string) ([]byte, error) start.")))

	iv, key, err := DecryptAESCBCModeV5(cipherIV, cipherKey)
	if err != nil {
		return nil, err
	}

	dec, err := EncryptAESCBCModeV1(iv, key, []byte(plainText))
	if err != nil {
		return nil, err
	}

	return dec, nil
}

// DecryptAESCBCModeV6 ...
func DecryptAESCBCModeV6(cipherIV string, cipherKey string, cipherText string) ([]byte, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/cipher_util.go : DecryptAESCBCModeV6 : func DecryptAESCBCModeV6(cipherIV string, cipherKey string, cipherText string) ([]byte, error) start.")))

	iv, key, err := DecryptAESCBCModeV5(cipherIV, cipherKey)
	if err != nil {
		return nil, err
	}

	cipherb, err := Base64StdEncodingDecodeStringV1(cipherText)
	if err != nil {
		return nil, err
	}

	dec, err := DecryptAESCBCModeV1(iv, key, cipherb)
	if err != nil {
		return nil, err
	}

	return dec, nil
}

// EncryptAESCBCModeV7 ...
func EncryptAESCBCModeV7(cipherIV string, cipherKey string, plainText string) (string, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/cipher_util.go : EncryptAESCBCModeV7 : func EncryptAESCBCModeV7(cipherIV string, cipherKey string, plainText string) (string, error) start.")))

	enc, err := EncryptAESCBCModeV6(cipherIV, cipherKey, plainText)
	if err != nil {
		return "", err
	}

	return Base64StdEncodingEncodeToStringV1(enc), nil
}

// DecryptAESCBCModeV7 ...
func DecryptAESCBCModeV7(cipherIV string, cipherKey string, cipherText string) (string, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/cipher_util.go : DecryptAESCBCModeV7 : func DecryptAESCBCModeV7(cipherIV string, cipherKey string, cipherText string) (string, error) start.")))

	dec, err := DecryptAESCBCModeV6(cipherIV, cipherKey, cipherText)
	if err != nil {
		return "", err
	}

	return string(dec), nil
}

// EncryptAESCBCModeV8 ...
func EncryptAESCBCModeV8(cipherIV string, cipherKey string, plainText string) ([]byte, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/cipher_util.go : EncryptAESCBCModeV8 : func EncryptAESCBCModeV8(cipherIV string, cipherKey string, plainText string) ([]byte, error) start.")))

	iv, err := HexDecodeStringV1(cipherIV)
	if err != nil {
		return nil, err
	}

	key, err := HexDecodeStringV1(cipherKey)
	if err != nil {
		return nil, err
	}

	dec, err := EncryptAESCBCModeV1(iv, key, []byte(plainText))
	if err != nil {
		return nil, err
	}

	return dec, nil
}

// DecryptAESCBCModeV8 ...
func DecryptAESCBCModeV8(cipherIV string, cipherKey string, cipherText string) ([]byte, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/cipher_util.go : DecryptAESCBCModeV8 : func DecryptAESCBCModeV8(cipherIV string, cipherKey string, cipherText string) ([]byte, error) start.")))

	iv, err := HexDecodeStringV1(cipherIV)
	if err != nil {
		return nil, err
	}

	key, err := HexDecodeStringV1(cipherKey)
	if err != nil {
		return nil, err
	}

	cipherb, err := Base64StdEncodingDecodeStringV1(cipherText)
	if err != nil {
		return nil, err
	}

	dec, err := DecryptAESCBCModeV1(iv, key, cipherb)
	if err != nil {
		return nil, err
	}

	return dec, nil
}

// EncryptAESCBCModeV9 ...
func EncryptAESCBCModeV9(cipherIV string, cipherKey string, plainText string) (string, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/cipher_util.go : EncryptAESCBCModeV9 : func EncryptAESCBCModeV9(cipherIV string, cipherKey string, plainText string) (string, error) start.")))

	enc, err := EncryptAESCBCModeV8(cipherIV, cipherKey, plainText)
	if err != nil {
		return "", err
	}

	return Base64StdEncodingEncodeToStringV1(enc), nil
}

// DecryptAESCBCModeV9 ...
func DecryptAESCBCModeV9(cipherIV string, cipherKey string, cipherText string) (string, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/cipher_util.go : DecryptAESCBCModeV9[] : func DecryptAESCBCModeV9(cipherIV string, cipherKey string, cipherText string) (string, error) start.")))

	dec, err := DecryptAESCBCModeV8(cipherIV, cipherKey, cipherText)
	if err != nil {
		return "", err
	}

	return string(dec), nil
}

// EncryptAESCBCModeV10 ...
func EncryptAESCBCModeV10(cipherIV string, cipherKey string, plain []byte) ([]byte, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/cipher_util.go : EncryptAESCBCModeV10 : func EncryptAESCBCModeV10(cipherIV string, cipherKey string, plain []byte) ([]byte, error) start.")))

	iv, err := HexDecodeStringV1(cipherIV)
	if err != nil {
		return nil, err
	}

	key, err := HexDecodeStringV1(cipherKey)
	if err != nil {
		return nil, err
	}

	enc, err := EncryptAESCBCModeV1(iv, key, plain)
	if err != nil {
		return nil, err
	}

	return enc, nil
}

// EncryptAESCBCModeV11 ...
func EncryptAESCBCModeV11(cipherIV []byte, cipherKey string, plain []byte) (string, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/cipher_util.go : EncryptAESCBCModeV11 : func EncryptAESCBCModeV8(cipherIV string, cipherKey string, plain []byte) ([]byte, error) start.")))

	key, err := HexDecodeStringV1(cipherKey)
	if err != nil {
		return "", err
	}

	enc, err := EncryptAESCBCModeV1(cipherIV, key, plain)
	if err != nil {
		return "", err
	}

	ivEnc := append(cipherIV, enc...)

	return Base64StdEncodingEncodeToStringV1(ivEnc), nil
}

// DecryptAESCBCModeV11 ...
func DecryptAESCBCModeV11(cipherKey string, cipherText string) ([]byte, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/cipher_util.go : DecryptAESCBCModeV11 : func DecryptAESCBCModeV11(cipherKey string, cipherText string) ([]byte, error) start.")))

	key, err := HexDecodeStringV1(cipherKey)
	if err != nil {
		return nil, err
	}

	cipherb, err := Base64StdEncodingDecodeStringV1(cipherText)
	if err != nil {
		return nil, err
	}

	ivb := cipherb[0:16]
	encb := cipherb[16:]

	dec, err := DecryptAESCBCModeV1(ivb, key, encb)
	if err != nil {
		return nil, err
	}

	return dec, nil
}

// // DecryptRSAOAEPV1 ...
// func DecryptRSAOAEPV1(cipherText []byte, privKey rsa.PrivateKey, label []byte) ([]byte, error) {
// 	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/cipher_util.go : DecryptRSAOAEPV1 : func DecryptRSAOAEPV1(cipherText []byte, privKey rsa.PrivateKey, label []byte) ([]byte, error) start.")))

// 	// label := []byte("OAEP Encrypted")
// 	rng := rand.Reader
// 	plainText, err := rsa.DecryptOAEP(sha256.New(), rng, &privKey, cipherText, label)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return plainText, nil
// }

// // ReadRsaPrivateKeyV1 ...
// func ReadRsaPrivateKeyV1(pemFile string) (*rsa.PrivateKey, error) {
// 	bytes, err := ioutil.ReadFile(pemFile)
// 	if err != nil {
// 		return nil, err
// 	}

// 	block, _ := pem.Decode(bytes)
// 	if block == nil {
// 		return nil, errors.New("invalid private key data")
// 	}

// 	var key *rsa.PrivateKey
// 	if block.Type == "RSA PRIVATE KEY" {
// 		key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
// 		if err != nil {
// 			return nil, err
// 		}
// 	} else if block.Type == "PRIVATE KEY" {
// 		keyInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
// 		if err != nil {
// 			return nil, err
// 		}
// 		var ok bool
// 		key, ok = keyInterface.(*rsa.PrivateKey)
// 		if !ok {
// 			return nil, errors.New("not RSA private key")
// 		}
// 	} else {
// 		return nil, fmt.Errorf("invalid private key type : %s", block.Type)
// 	}

// 	key.Precompute()

// 	if err := key.Validate(); err != nil {
// 		return nil, err
// 	}

// 	return key, nil
// }

// // ReadRsaPublicKeyV1 ...
// func ReadRsaPublicKeyV1(path string) (*rsa.PublicKey, error) {
// 	bytes, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		return nil, err
// 	}

// 	block, _ := pem.Decode(bytes)
// 	if block == nil {
// 		return nil, errors.New("invalid public key data")
// 	}
// 	if block.Type != "PUBLIC KEY" {
// 		return nil, fmt.Errorf("invalid public key type : %s", block.Type)
// 	}

// 	keyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	key, ok := keyInterface.(*rsa.PublicKey)
// 	if !ok {
// 		return nil, errors.New("not RSA public key")
// 	}

// 	return key, nil
// }

// HmacSha256V1 ...
func HmacSha256V1(key []byte, plainText []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(plainText)
	return mac.Sum(nil)
}

// HmacSha256V2 ...
func HmacSha256V2(key string, plainText string) (string, error) {
	decKey, err := DecryptAESCBCModeV4p1(key)
	if err != nil {
		return "", err
	}
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/cipher_util.go : HmacSha256V2 : decKey=%s", decKey)))

	return Base64StdEncodingEncodeToStringV1(HmacSha256V1([]byte(decKey), []byte(plainText))), nil
}

// HmacSha256V3 ...
func HmacSha256V3(key string, plainText string) (string, error) {
	return Base64StdEncodingEncodeToStringV1(HmacSha256V1([]byte(key), []byte(plainText))), nil
}

// CompareHmacSha256V1 ...
func CompareHmacSha256V1(key []byte, plainText string, hmacSha256Base64Text string) bool {
	plain2HmacSha256Base64Text := Base64StdEncodingEncodeToStringV1(HmacSha256V1(key, []byte(plainText)))
	if hmacSha256Base64Text == plain2HmacSha256Base64Text {
		return true
	}
	return false
}

// URLSafeHmacSha256SignatureV1 ...
func URLSafeHmacSha256SignatureV1(message string, keystr string) string {
	plain2HmacSha256 := HmacSha256V1([]byte(keystr), []byte(message))
	plain2HmacSha256Base64Text := base64.StdEncoding.EncodeToString(plain2HmacSha256)
	strRep := strings.NewReplacer("=", "", "/", "_", "+", "-")
	plain2HmacSha256UrlSafeBase64Text := strRep.Replace(plain2HmacSha256Base64Text)
	return plain2HmacSha256UrlSafeBase64Text
}

// HmacSha512V1 ...
func HmacSha512V1(key []byte, plainText []byte) []byte {
	mac := hmac.New(sha512.New, key)
	mac.Write(plainText)
	return mac.Sum(nil)
}

// URLSafeHmacSha512SignatureV1 ...
func URLSafeHmacSha512SignatureV1(message string, keystr string) string {
	plain2HmacSha512 := HmacSha512V1([]byte(keystr), []byte(message))
	plain2HmacSha512Base64Text := base64.StdEncoding.EncodeToString(plain2HmacSha512)
	strRep := strings.NewReplacer("=", "", "/", "_", "+", "-")
	plain2HmacSha512UrlSafeBase64Text := strRep.Replace(plain2HmacSha512Base64Text)
	return plain2HmacSha512UrlSafeBase64Text
}

// URLSafeBase64V1 ...
func URLSafeBase64V1(plainText string) string {
	plain2Base64Text := base64.StdEncoding.EncodeToString([]byte(plainText))
	strRep := strings.NewReplacer("=", "", "/", "_", "+", "-")
	plain2UrlSafeBase64Text := strRep.Replace(plain2Base64Text)
	return plain2UrlSafeBase64Text
}

// URLSafeBase64DecodeV1 ...
func URLSafeBase64DecodeV1(encText string) ([]byte, error) {
	// return base64.URLEncoding.DecodeString(encText)
	return base64.RawURLEncoding.DecodeString(encText)
}

// HexEncodeToStringV1 ...
func HexEncodeToStringV1(src []byte) string {
	return hex.EncodeToString(src)
}

// HexDecodeStringV1 ...
func HexDecodeStringV1(s string) ([]byte, error) {
	return hex.DecodeString(s)
}

// Base64StdEncodingEncodeToStringV1 ...
func Base64StdEncodingEncodeToStringV1(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

// Base64StdEncodingDecodeStringV1 ...
func Base64StdEncodingDecodeStringV1(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

// VerifySignatureV1 ...
func VerifySignatureV1(message string, keystr string, signature string) bool {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/cipher_util.go : VerifySignatureV1 : func VerifySignatureV1(message string, keystr string, signature string) bool start.")))

	mac := hmac.New(sha256.New, []byte(keystr))
	mac.Write([]byte(message))
	plain2HmacSha256 := mac.Sum(nil)
	plain2HmacSha256Base64Text := base64.StdEncoding.EncodeToString(plain2HmacSha256)
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/cipher_util.go : VerifySignatureV1 : plain2HmacSha256Base64Text=%s", plain2HmacSha256Base64Text)))
	if signature == plain2HmacSha256Base64Text {
		return true
	}
	return false
}

// VerifySignatureV2 ...
func VerifySignatureV2(message string, keystr string, signature string) bool {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/cipher_util.go : VerifySignatureV2 : VerifySignatureV2(message string, keystr string, signature string) bool start.")))

	// PEMの中身はDERと同じASN.1のバイナリデータをBase64によってエンコーディングされたテキストなのでBase64でデコードする
	// ゆえにDERエンコード形式に変換
	keyBytes, err := base64.StdEncoding.DecodeString(keystr)
	if err != nil {
		return false
	}

	// DERでエンコードされた公開鍵を解析する
	// 成功すると、pubは* rsa.PublicKey、* dsa.PublicKey、または* ecdsa.PublicKey型になる
	pub, err := x509.ParsePKIXPublicKey(keyBytes)
	if err != nil {
		return false
	}

	// 署名文字列はBase64でエンコーディングされたテキストなのでBase64でデコードする
	signDataByte, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false
	}

	// SHA-256のハッシュ関数を使って受信データのハッシュ値を算出する
	h := crypto.Hash.New(crypto.SHA256)
	h.Write([]byte(message))
	hashed := h.Sum(nil)

	// 署名の検証、有効な署名はnilを返すことによって示される
	// ここで何をしているかというと、、
	// ①送信者のデータ（署名データ）を公開鍵で復号しハッシュ値を算出
	// ②受信側で算出したハッシュ値と、①のハッシュ値を比較し、一致すれば、「送信者が正しい」「データが改ざんされていない」ということを確認できる
	err = rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA256, hashed, signDataByte)
	if err != nil {
		return false
	}

	return true
}

// VerifySignatureV3 ...
func VerifySignatureV3(message string, keystr string, signature string) bool {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/cipher_util.go : VerifySignatureV3 : func VerifySignatureV3(message string, keystr string, signature string) bool start.")))

	plain2HmacSha512UrlSafeBase64Text := URLSafeHmacSha512SignatureV1(message, keystr)
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/cipher_util.go : VerifySignatureV3 : plain2HmacSha512UrlSafeBase64Text=%s", plain2HmacSha512UrlSafeBase64Text)))
	if signature == plain2HmacSha512UrlSafeBase64Text {
		return true
	}
	return false
}
