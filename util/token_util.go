package util

// TokenV1 ...
func TokenV1() string {
	return UUIDV2()
}

// TokenV2 ...
func TokenV2(keystr string) string {
	uuidString := UUIDV2()
	signatureString := URLSafeHmacSha256SignatureV1(uuidString, keystr)
	uiidSignatureString := uuidString + "." + signatureString
	return uiidSignatureString
}

// TokenV3 ...
func TokenV3(cipherIV string, cipherKey string, cipherText string) (string, error) {
	plainText, err := DecryptAESCBCModeV7(cipherIV, cipherKey, cipherText)
	if err != nil {
		return "", err
	}

	return TokenV2(plainText), nil
}

// TokenV4 ...
func TokenV4(cipherIV string, cipherKey string, cipherText string) (string, error) {
	plainText, err := DecryptAESCBCModeV9(cipherIV, cipherKey, cipherText)
	if err != nil {
		return "", err
	}

	return TokenV2(plainText), nil
}
