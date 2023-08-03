package security

import (
	"bot/lib/utils"
	"encoding/base64"
	"fmt"
)

func XOR(Input string) (Out string) {
	kL := len(utils.EncryptionKey)

	for i := range Input {
		Out += string(Input[i] ^ utils.EncryptionKey[i%kL])
	}

	return Out
}

func EncryptStr(Content string) string {
	Content = XOR(Content)
	return fmt.Sprintf("%s\n", base64.StdEncoding.EncodeToString([]byte(Content)))
}

func DecryptStr(Content string) string {
	Bytes, Err := base64.StdEncoding.DecodeString(Content)

	if utils.HandleError(Err) {
		return ""
	}

	return XOR(string(Bytes))
}
