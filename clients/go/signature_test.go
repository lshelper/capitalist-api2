package capitalist

import "testing"

func TestCalculateSignatureDocumentationExample(t *testing.T) {
	timestamp := "1678901234567"
	body := []byte(`{"type":"PAYONEER","accountFrom":"U0123504","payload":{"account":"[email protected]"},"userRequestId":"1745844029437","callbackUrl":"https://some-domain.com/callbacksFromCap","amount":100}`)
	secret := "somePrivateApiSecret"
	expected := "57bd25d836c076c24a81f1f088f405b5fd371918ba96f9a28df72d39860396b6"

	if got := CalculateSignature(timestamp, body, secret); got != expected {
		t.Fatalf("CalculateSignature() = %s, want %s", got, expected)
	}

	if !VerifySignature(timestamp, body, secret, expected) {
		t.Fatal("VerifySignature() = false, want true")
	}
}
