package goutils

import (
	"testing"
)

type  mt struct{}

func TestDirDirWalk(t *testing.T) {
	tests := []struct{
		path string
		expected map[string]mt
		expectedErr string
	}{
		{
			path:  "invalid", 
			expected: nil,
			expectedErr: "invalid does not exist",
		},
		{
                        path:  "test/pixies", 
                        expected: map[string]mt{
                                "I-cant-forget.txt": mt{},
				"Ive-been-waiting-for-you.txt": mt{},
				"born-in-chicago.txt": mt{},
                                "debaser-I.txt": mt{},
                                "debaser-II.txt": mt{},
                                "gigantic.txt": mt{},
				"something-against-you.txt": mt{},
				"where-is-my-mind.txt": mt{},
				"surfer-rosa": mt{},
                                "doolittle": mt{},
                        },      
                        expectedErr: "",
                },
              	{
                        path:  "test/pink-floyd", 
                        expected: map[string]mt{
                                "echos.txt": mt{},
				"one-of-these-days.txt": mt{},
                                "san-tropez.txt": mt{},
                                "seamus.txt": mt{},
                                "I-V.txt": mt{},
                                "VI-IX.txt": mt{},
                                "have-a-cigar.txt": mt{},
                                "wish-you-were-here-part1.txt": mt{},
                                "wish-you-were-here-part2.txt": mt{},
                                "wish-you-were-here-part3.txt": mt{},
                                "wish-you-were-here": mt{},
                                "meddle": mt{},
                                "shine-on-you-crazy-diamond": mt{},
                        },      
                        expectedErr: "",
                },
              	{
                        path:  "test", 
                       	expected: map[string]mt{
                                "I-cant-forget.txt": mt{},
                                "Ive-been-waiting-for-you.txt": mt{},
                                "born-in-chicago.txt": mt{},
                                "debaser-I.txt": mt{},
                                "debaser-II.txt": mt{},
                                "gigantic.txt": mt{},
                                "something-against-you.txt": mt{},
                                "where-is-my-mind.txt": mt{},
                                "echos.txt": mt{},
                                "one-of-these-days.txt": mt{},
                                "san-tropez.txt": mt{},
                                "seamus.txt": mt{},
                                "I-V.txt": mt{},
                                "VI-IX.txt": mt{},
                                "have-a-cigar.txt": mt{},
                                "wish-you-were-here-part1.txt": mt{},
                                "wish-you-were-here-part2.txt": mt{},
                                "wish-you-were-here-part3.txt": mt{},
				"tmbg-ana-ng.txt": mt{},
                                "tmbg-particle-man.txt": mt{},
                                "tmbg-sapphire-bullets-of-pure-love.txt": mt{},
				"pixies": mt{},
				"pink-floyd": mt{},
                                "surfer-rosa": mt{},
                                "doolittle": mt{},
                                "wish-you-were-here": mt{},
                                "meddle": mt{},
                                "shine-on-you-crazy-diamond": mt{},
                        },
                        expectedErr: "",
                },

	}

	
	for _, test := range tests {
		d := &Dir{Files: []file{}}
		err := d.Walk(test.path)
		if err != nil {
			if err.Error() != test.expectedErr {
				t.Errorf("Expected %s got %s", test.expectedErr, err)
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
	tests := []struct{
		path	 string
		expected bool
		errS	 string
	}{
		{"", false, ""},
		{"test", true, ""},
		{"test/pixies/born-in-chicago.txt", true, ""},
		{"tst", false, ""},
		{"test/pink-floyd/animals", false, ""},
	}

	for  _, test := range tests{
		exists, err := PathExists(test.path)
		if err != nil {
			if test.errS == "" {
				t.Errorf("An unexpected error was encountered while checking the existence of %s: %s", test.path, err.Error())
			} else {
				if err.Error() != test.errS {
					t.Errorf("Expected %s, got %s", test.errS, err.Error())
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
