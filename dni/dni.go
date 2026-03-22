package dni

import (
	"fmt"
	"math/rand/v2"
)

const lletresControl = "TRWAGMYFPDXBNJZSQVHLCKE"

func GeneraNumero() int {
	return rand.IntN(100000000)
}

func CalculaLletra(numero int) byte {
	return lletresControl[numero%len(lletresControl)]
}

func FormatDNI(numero int) string {
	return fmt.Sprintf("%08d%c", numero, CalculaLletra(numero))
}

func GeneraDNI() string {
	return FormatDNI(GeneraNumero())
}