package goutils

import (
	"testing"
)

type mt struct{}

func TestDirDirWalk(t *testing.T) {
	tests := []struct {
		path        string
		expected    map[string]mt
		expectedErr string
	}{
		{
			path:        "invalid",
			expected:    nil,
			expectedErr: "invalid does not exist",
		},
		{
			path: "test/pixies",
			expected: map[string]mt{
				"I-cant-forget.txt":            {},
				"Ive-been-waiting-for-you.txt": {},
				"born-in-chicago.txt":          {},
				"debaser-I.txt":                {},
				"debaser-II.txt":               {},
				"gigantic.txt":                 {},
				"something-against-you.txt":    {},
				"where-is-my-mind.txt":         {},
				"surfer-rosa":                  {},
				"doolittle":                    {},
			},
			expectedErr: "",
		},
		{
			path: "test/pink-floyd",
			expected: map[string]mt{
				"echos.txt":                    {},
				"one-of-these-days.txt":        {},
				"san-tropez.txt":               {},
				"seamus.txt":                   {},
				"I-V.txt":                      {},
				"VI-IX.txt":                    {},
				"have-a-cigar.txt":             {},
				"wish-you-were-here-part1.txt": {},
				"wish-you-were-here-part2.txt": {},
				"wish-you-were-here-part3.txt": {},
				"wish-you-were-here":           {},
				"meddle":                       {},
				"shine-on-you-crazy-diamond":   {},
			},
			expectedErr: "",
		},
		{
			path: "test",
			expected: map[string]mt{
				"I-cant-forget.txt":                      {},
				"Ive-been-waiting-for-you.txt":           {},
				"born-in-chicago.txt":                    {},
				"debaser-I.txt":                          {},
				"debaser-II.txt":                         {},
				"gigantic.txt":                           {},
				"something-against-you.txt":              {},
				"where-is-my-mind.txt":                   {},
				"echos.txt":                              {},
				"one-of-these-days.txt":                  {},
				"san-tropez.txt":                         {},
				"seamus.txt":                             {},
				"I-V.txt":                                {},
				"VI-IX.txt":                              {},
				"have-a-cigar.txt":                       {},
				"wish-you-were-here-part1.txt":           {},
				"wish-you-were-here-part2.txt":           {},
				"wish-you-were-here-part3.txt":           {},
				"tmbg-ana-ng.txt":                        {},
				"tmbg-particle-man.txt":                  {},
				"tmbg-sapphire-bullets-of-pure-love.txt": {},
				"pixies":                     {},
				"pink-floyd":                 {},
				"surfer-rosa":                {},
				"doolittle":                  {},
				"wish-you-were-here":         {},
				"meddle":                     {},
				"shine-on-you-crazy-diamond": {},
			},
			expectedErr: "",
		},
	}

	for _, test := range tests {
		d := &Dir{Files: []file{}}
		err := d.Walk(test.path)
		if err != nil {
			if err.Error() != test.expectedErr {
				t.Errorf("Expected %q got %q", test.expectedErr, err)
			}
		} else {
			if test.expectedErr != "" {
				t.Errorf("Expected error %s", err)
			} else {
				for _, f := range d.Files {
					if _, ok := test.expected[f.Info.Name()]; !ok {
						t.Errorf("%s was indexed but not found in the expected filename list", f.Info.Name())
					}
				}
			}
		}
	}
}

func TestPathExists(t *testing.T) {
	tests := []struct {
		path     string
		expected bool
		errS     string
	}{
		{"", false, ""},
		{"../test", true, ""},
		{"../test/pixies/born-in-chicago.txt", true, ""},
		{"../tst", false, ""},
		{"../test/pink-floyd/animals", false, ""},
	}

	for _, test := range tests {
		exists, err := PathExists(test.path)
		if err != nil {
			if test.errS == "" {
				t.Errorf("An unexpected error was encountered while checking the existence of %s: %s", test.path, err.Error())
			} else {
				if err.Error() != test.errS {
					t.Errorf("Expected %q, got %q", test.errS, err.Error())
				}
			}
		} else {
			if test.errS != "" {
				t.Errorf("%s was expected, but no error was encountered.", test.errS)
			} else {
				if exists != test.expected {
					t.Errorf("Expected %v got %v for %s", test.expected, exists, test.path)
				}
			}
		}
	}
}
