package convert

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	inputString := "MzIzNDM5MzQzI3NvdXJjZQ=="
	encodedString := base64.StdEncoding.EncodeToString([]byte(inputString))
	fmt.Println("pre_" + encodedString)
	fmt.Println(len("pre_" + encodedString))
}
