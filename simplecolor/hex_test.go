package simplecolor

import "testing"

func TestFromHexStringFallbacks(t *testing.T) {
	const fallback = "#66042D"

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "hashless long", input: "00FF00", want: "#00FF00"},
		{name: "hashless short", input: "0F0", want: "#00FF00"},
		{name: "lowercase", input: "#abcdef", want: "#ABCDEF"},
		{name: "too short", input: "#12", want: fallback},
		{name: "too long", input: "#1234567", want: fallback},
		{name: "invalid digit", input: "#GGGGGG", want: fallback},
		{name: "empty", input: "", want: fallback},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := FromHexString(test.input).ToHex()
			if got != test.want {
				t.Fatalf("FromHexString(%q) = %s, want %s", test.input, got, test.want)
			}
		})
	}
}
