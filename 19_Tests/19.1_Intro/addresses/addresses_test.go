package addresses

import (
	"testing"
)

type testCases struct {
	input    string
	expected string
}

func TestAddressType(t *testing.T) {
	testCases := []testCases{
		{"Avenida Paulista", "avenida"},
		{"Rua Antônio Costa Rosendo", "rua"},
		{"Estrada de Barro", "estrada"},
		{"Rodovia dos Tamoios", "rodovia"},
		{"Praça Londres", "Invalid type"},
	}

	for _, test := range testCases {
		if test.expected != AddressType(test.input) {
			t.Errorf("Address %s is not %s", test.input, test.expected)
		}
	}
}
