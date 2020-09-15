package base62_test

import (
	"testing"

	"github.com/d4n13l-4lf4/restful-web-services-with-go/chapter07/base62Example/base62"
)

func TestEncodeBase62(t *testing.T) {
	expectedEncode := "1C"
	numberToEncode := 100
	result := base62.ToBase62(numberToEncode)

	if result != expectedEncode {
		t.Errorf("Expected encode of %d was %s, it does not match %s", numberToEncode, result, expectedEncode)
	}
}

func TestDecodeBase62(t *testing.T) {
	expectedDecode := 100
	encodeToDecode := "1C"
	result := base62.ToBase10(encodeToDecode)

	if result != expectedDecode {
		t.Errorf("Expected decode of %s was %d, it does not match %d", encodeToDecode, result, expectedDecode)
	}
}
