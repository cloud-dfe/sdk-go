package sdk_cloud_dfe

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"
)

func IsValid(token string, payload map[string]interface{}) (bool, error) {

	// Verifica se o payload não está vazio
	signature, ok := payload["signature"].(string)
	if !ok || signature == "" {
		return false, errors.New("payload incorreto não contém a assinatura")
	}

	// Verifica se o token não está vazio
	if token == "" {
		return false, errors.New("token vazio")
	}

	// Converte a chave
	key := convertKey(token)

	// Decodifica a assinatura da base64
	c, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, errors.New("falha ao decodificar a assinatura")
	}

	ivlen := 16

	iv := c[:ivlen]
	hmacOriginal := c[ivlen : ivlen+32]
	ciphertextRaw := c[48:]

	originalTime, err := decryptTime(ciphertextRaw, key, iv)
	if err != nil {
		return false, err
	}

	mac := hmac.New(sha256.New, []byte(token))
	mac.Write(ciphertextRaw)
	calcmac := mac.Sum(nil)

	if subtle.ConstantTimeCompare(hmacOriginal, calcmac) == 1 {

		currentTime := time.Now().Unix()
		if currentTime-originalTime < 300 {
			return true, nil
		}
		return false, errors.New("assinatura expirou")
	}

	return false, errors.New("token ou assinatura incorreta")
}

func convertKey(token string) []byte {

	key := token[:16]
	if len(key) < 16 {
		key = key + strings.Repeat("0", 16-len(key))
	}
	return []byte(key)
}

func decryptTime(ciphertextRaw, key, iv []byte) (int64, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return 0, err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(ciphertextRaw))
	mode.CryptBlocks(decrypted, ciphertextRaw)

	decrypted = bytes.TrimRight(decrypted, "\x00")
	decryptedStr := string(decrypted)

	var decryptedTime int64
	fmt.Sscanf(decryptedStr, "%d", &decryptedTime)

	return decryptedTime, nil
}
