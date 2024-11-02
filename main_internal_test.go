package main

import "testing"
/*
func ExampleMain() {
	main()
	// Output:
	// Hello World!
}
*/

func TestGreet(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}

	var tests = map[string] testCase{
		"English": {
			lang: "en",
			want: "Hello World",
		},

		"French": {
			lang: "fr",
			want: "Bonjour le monde",
		},

		"Greek": {
			lang: "el",
			want: "ηελλο ςορλδ",
		},

		"Yorube": {
			lang: "yr",
			want: "Eku Ojumo",
		},

		"Akkadian, not supported": {
			lang: "akk",
			want: `unsupported language: "akk"`,
		},

		"Empty": {
			lang: "",
			want: `unsupported language: ""`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)

			if got != tc.want {
				t.Errorf("expected: %q, got: %q", tc.want, got)
			}
		})
	}

}
