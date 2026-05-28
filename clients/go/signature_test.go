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

func TestVerifySignatureRejectsChangedCallbackBody(t *testing.T) {
	timestamp := "1678901234567"
	body := []byte(`{"state":"EXECUTED","fee":1.12,"documentId":123,"comment":"Own funds","amount":100,"currency":"USD","type":"PAYONEER","accountFrom":"U0123504","userRequestId":"9876543219","callbackUrl":"https://some-domain.com/for-callbacks"}`)
	secret := "somePrivateApiSecret"
	signature := CalculateSignature(timestamp, body, secret)

	if !VerifySignature(timestamp, body, secret, signature) {
		t.Fatal("VerifySignature() = false, want true")
	}

	changedBody := append([]byte{}, body...)
	changedBody = append(changedBody, '\n')
	if VerifySignature(timestamp, changedBody, secret, signature) {
		t.Fatal("VerifySignature() = true, want false")
	}
}
