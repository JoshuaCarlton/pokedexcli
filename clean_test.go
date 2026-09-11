package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "What Do you Do?",
			expected: []string{"what", "do", "you", "do?"},
		},
		{
			input:    "a    bvd c d fFf loop",
			expected: []string{"a", "bvd", "c", "d", "fff", "loop"},
		},
		{
			input:    "these	words	are	sePerated	by	TABS",
			expected: []string{"these", "words", "are", "seperated", "by", "tabs"},
		},
		{
			input:    "ThIs IS hOw NoRMaL PeoPlE tALk RIghT?",
			expected: []string{"this", "is", "how", "normal", "people", "talk", "right?"},
		},
	}
	for cn, c := range cases {
		actual := cleanInput(c.input)
		expected := c.expected
		if len(actual) != len(expected) {
			t.Errorf("test number %d failed, expected length %d actual length %d", cn, len(expected), len(actual))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := expected[i]
			if word != expectedWord {
				t.Errorf("test number %d failed, expected %v actual %v", cn, expectedWord, word)
				t.Fail()
			}
		}
	}
}
