package main

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"slices"
	"testing"
)

func TestDetectTypeAndPrint(t *testing.T) {
	tests := []struct {
		name  string
		value any
		want  string
	}{
		{"dec", 12, "Тип переменной dec: int\n"},
		{"str", "hello", "Тип переменной str: string\n"},
		{"fl", 122.4, "Тип переменной fl: float64\n"},
		{"isTrue", false, "Тип переменной isTrue: bool\n"},
		{"cmplx", complex64(321 + 5i), "Тип переменной cmplx: complex64\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output bytes.Buffer

			detectTypeAndPrint(&output, tt.name, tt.value)

			if got := output.String(); got != tt.want {
				t.Errorf("got = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestParamsString(t *testing.T) {
	tests := []struct {
		name        string
		inputParams map[string]any
		inputKeys   []string
		wantString  string
	}{
		{
			name: "Best way",
			inputParams: map[string]any{
				"dec":    12,
				"hex":    0x355,
				"oct":    0o14312,
				"fl":     122.4,
				"str":    "mfrgngnkv",
				"isTrue": false,
				"cmplx":  321 + 5i,
			},
			inputKeys: []string{
				"dec",
				"hex",
				"oct",
				"fl",
				"str",
				"isTrue",
				"cmplx",
			},
			wantString: "128536346122.4mfrgngnkvfalse(321+5i)",
		},
		{
			name: "Inverted keys",
			inputParams: map[string]any{
				"dec":    12,
				"hex":    0x355,
				"oct":    0o14312,
				"fl":     122.4,
				"str":    "mfrgngnkv",
				"isTrue": false,
				"cmplx":  321 + 5i,
			},
			inputKeys: []string{
				"cmplx",
				"isTrue",
				"str",
				"fl",
				"oct",
				"hex",
				"dec",
			},
			wantString: "(321+5i)falsemfrgngnkv122.4634685312",
		},
		{
			name:       "Empty keys and map",
			wantString: "",
		},
		{
			name: "Empty keys",
			inputParams: map[string]any{
				"dec":    12,
				"hex":    0x355,
				"oct":    0o14312,
				"fl":     122.4,
				"str":    "mfrgngnkv",
				"isTrue": false,
				"cmplx":  321 + 5i,
			},
			wantString: "",
		},
		{
			name: "Empty map",
			inputKeys: []string{
				"cmplx",
				"isTrue",
				"str",
				"fl",
				"oct",
				"hex",
				"dec",
			},
			wantString: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			paramsStr := paramsString(tt.inputParams, tt.inputKeys)

			if paramsStr != tt.wantString {
				t.Errorf("TestParamsString: got: %s, want: %s", paramsStr, tt.wantString)
			}
		})
	}
}

func TestCreateHashAndGet(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		salt    string
		salted  string
		wantErr error
	}{
		{
			name:   "Even length",
			input:  "abcd",
			salt:   "go-2024",
			salted: "abgo-2024cd",
		},
		{
			name:   "Odd length",
			input:  "abcde",
			salt:   "go-2024",
			salted: "abgo-2024cde",
		},
		{
			name:   "Unicode",
			input:  "Я🙂你",
			salt:   "go-2024",
			salted: "Яgo-2024🙂你",
		},
		{
			name:   "Empty salt",
			input:  "abcd",
			salted: "abcd",
		},
		{
			name:    "Empty input",
			salt:    "go-2024",
			wantErr: errRunesIsEmpty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runes := make([]rune, len([]rune(tt.input)), 64)
			copy(runes, []rune(tt.input))
			before := slices.Clone(runes)

			got, err := createHashAndGet(runes, tt.salt)

			if !slices.Equal(runes, before) {
				t.Errorf("slice changed: %q → %q",
					string(before), string(runes))
			}

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("error = %v, want error: %v",
					err, tt.wantErr)
			}
			if tt.wantErr != nil {
				return
			}

			want := sha256.Sum256([]byte(tt.salted))
			if got != want {
				t.Errorf("hash = %x, want = %x", got, want)
			}
		})
	}
}

func TestRuneSliceFromString(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    []rune
		wantErr error
	}{
		{"ASCII", "hello", []rune{'h', 'e', 'l', 'l', 'o'}, nil},
		{"Cyrillic", "Привет", []rune{'П', 'р', 'и', 'в', 'е', 'т'}, nil},
		{"Mixed Unicode", "Я🙂你", []rune{'Я', '🙂', '你'}, nil},
		{"Single emoji", "🙂", []rune{'🙂'}, nil},
		{"Whitespace", " \t\n", []rune{' ', '\t', '\n'}, nil},
		{"Combining character", "e\u0301", []rune{'e', '\u0301'}, nil},
		{"Empty input", "", []rune{}, errStrIsEmpty},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := runeSliceFromString(tt.value)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("error = %v, want error: %v", err, tt.wantErr)
			}

			if !slices.Equal(got, tt.want) {
				t.Errorf("got = %q, want = %q", got, tt.want)
			}
		})
	}
}
