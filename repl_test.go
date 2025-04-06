package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct{
		input		string
		expected	[]string
	}{
		{
			input: "    Hello World    ",
			expected: []string{"hello", "world"},
		},
		// {
		// 	input: "    hello    world    ",
		// 	expected: []string{"hello", "world"},
		// },
		{
			input: "    heLLo worLd     ",
			expected: []string{"hello", "world"},
		},
		{
			input: "    HELLO world",
			expected: []string{"hello", "world"},
		},
		{
			input: "HELLO world    ",
			expected: []string{"hello", "world"},
		},
		{
			input: "hellO-world",
			expected: []string{"hello-world"},
		},
	}

	for _,c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) returned %d words, expected %d", c.input, len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("cleanInput(%q)[%d] = %q, expected %q", c.input, i, word, expectedWord)
			}
		}
	}
}