package capitalist

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"strconv"
)

func CalculateSignature(timestamp string, body []byte, apiSecret string) string {
	hash := sha256.New()
	hash.Write([]byte(timestamp))
	hash.Write(body)
	hash.Write([]byte(apiSecret))
	return hex.EncodeToString(hash.Sum(nil))
}

func CalculateSignatureMillis(timestamp int64, body []byte, apiSecret string) string {
	return CalculateSignature(strconv.FormatInt(timestamp, 10), body, apiSecret)
}

func VerifySignature(timestamp string, body []byte, apiSecret string, signature string) bool {
	expected, err := hex.DecodeString(CalculateSignature(timestamp, body, apiSecret))
	if err != nil {
		return false
	}

	actual, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}

	return len(expected) == len(actual) && subtle.ConstantTimeCompare(expected, actual) == 1
}
