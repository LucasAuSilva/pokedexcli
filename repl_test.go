package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "charmander  ",
			expected: []string{"charmander"},
		},
		{
			input: "bulbasaur \n PIKACHU  charmander \n charmeleon",
			expected: []string{"bulbasaur", "pikachu", "charmander", "charmeleon"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("The lists don't match in length\nEXPECTED: %v\nRESULT: %v", c.expected, actual)
			return
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("%s doesn't match %s", word, expectedWord)
				return
			}
		}
	}
}
