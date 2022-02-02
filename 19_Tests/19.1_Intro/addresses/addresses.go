package addresses

import (
	"strings"
)

func AddressType(address string) string {
	validTypes := []string{"rua", "avenida", "estrada", "rodovia"}
	firstWordOfAddress := strings.ToLower(strings.Split(address, " ")[0])
	validType := false

	for _, t := range validTypes {
		if t == firstWordOfAddress {
			validType = true
		}
	}

	if validType {
		return firstWordOfAddress
	}

	return "Invalid type"
}
