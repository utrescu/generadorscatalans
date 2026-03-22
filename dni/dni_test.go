package dni

import "testing"

func TestCalculaLletra(t *testing.T) {
	tests := map[int]byte{
		0:        'T',
		1:        'R',
		12345678: 'Z',
		99999999: 'R',
	}

	for numero, esperada := range tests {
		if obtinguda := CalculaLletra(numero); obtinguda != esperada {
			t.Fatalf("CalculaLletra(%d) = %c, volia %c", numero, obtinguda, esperada)
		}
	}
}

func TestFormatDNI(t *testing.T) {
	if dni := FormatDNI(12345678); dni != "12345678Z" {
		t.Fatalf("FormatDNI(12345678) = %s, volia 12345678Z", dni)
	}

	if dni := FormatDNI(42); dni != "00000042L" {
		t.Fatalf("FormatDNI(42) = %s, volia 00000042L", dni)
	}
}
