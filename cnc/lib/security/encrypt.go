package security

import (
	"cnc/lib/utils"
	"fmt"
	"strings"
)

func XOR(Input string) (Out string) {
	kL := len(utils.EncryptionKey)

	for i := range Input {
		Out += string(Input[i] ^ utils.EncryptionKey[i%kL])
	}

	return Out
}

func EncryptStr(Content string) string {
	return strings.ToUpper(fmt.Sprintf("%s\n", XOR(Content)))
}

func DecryptStr(Content string) string {
	return XOR(Content)
}
