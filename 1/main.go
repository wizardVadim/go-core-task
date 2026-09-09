package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var dec int = 1
	var hex int = 0xff
	var oct int = 0o32
	var fl float64 = 1.0
	var str string = "some string"
	var isTrue bool = false
	var cmplx complex64 = 3 + 4i

	params := map[string]any{
		"dec":    dec,
		"hex":    hex,
		"oct":    oct,
		"fl":     fl,
		"str":    str,
		"isTrue": isTrue,
		"cmplx":  cmplx,
	}

	keys := []string{
		"dec",
		"hex",
		"oct",
		"fl",
		"str",
		"isTrue",
		"cmplx",
	}

	for _, k := range keys {
		detectTypeAndPrint(os.Stdout, k, params[k])
	}

	paramsString := paramsString(params, keys)
	fmt.Println(paramsString)

	runeSlice, err := runeSliceFromString(paramsString)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(runeSlice)

	hash, err := createHashAndGet(runeSlice, "go-2024")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("new hash: %x\n", hash)
}

func createHashAndGet(runes []rune, salt string) ([32]byte, error) {
	if len(runes) == 0 {
		return [32]byte{}, errRunesIsEmpty
	}

	mid := len(runes) / 2
	saltRunes := []rune(salt)
	var finalRunes []rune
	finalRunes = append(finalRunes, runes[:mid]...)
	finalRunes = append(finalRunes, saltRunes...)
	finalRunes = append(finalRunes, runes[mid:]...)

	hash := sha256.Sum256([]byte(string(finalRunes)))
	return hash, nil
}

func runeSliceFromString(str string) ([]rune, error) {
	if len(str) == 0 {
		return []rune{}, errStrIsEmpty
	}
	return []rune(str), nil
}

func detectTypeAndPrint(w io.Writer, paramName string, param any) {
	fmt.Fprintf(w, "Тип переменной %s: %T\n", paramName, param)
}

func paramsString(params map[string]any, keys []string) string {
	if len(params) == 0 {
		return ""
	}

	var builder strings.Builder

	for _, k := range keys {
		fmt.Fprintf(&builder, "%v", params[k])
	}
	return builder.String()
}
