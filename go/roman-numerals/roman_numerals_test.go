package romannumerals

import (
	// "fmt"
	"testing"

	// "github.com/stretchr/testify/require"
)

// func TestOurRomanNumeral(t *testing.T) {
// 	type tc struct {
// 		name     string
// 		input    int
// 		expected []int
// 	}

// 	test1 := tc{
// 		name:     "Cindy's Test!",
// 		input:    2345,
// 		expected: []int{2000, 300, 40, 5},
// 	}

// 	test := tc{
// 		name:     "Cindy's Test!",
// 		input:    2345,
// 		expected: nil,
// 	}

// 	// test2 := tc{
// 	// 	name:     "Jess's Test!",
// 	// 	input:    99,
// 	// 	expected: []int{90, 9},
// 	// }

// 	rn, err := RomanTest(test.input)
// 	require.Nil(t, err)
// 	require.Equal(t, test1.expected, rn)
// 	fmt.Println(rn)
// }

func TestRomanNumerals(t *testing.T) {
	tc := []romanNumeralTest{
		{0, "", true},
		{-1, "", true},
		{3001, "", true},
	}
	tc = append(tc, romanNumeralTests...)

	for _, test := range tc {
		actual, err := ToRomanNumeral(test.arabic)
		if err == nil && test.hasError {
			t.Errorf("ToRomanNumeral(%d) should return an error.", test.arabic)
			continue
		}
		if err != nil && !test.hasError {
			var _ error = err
			t.Errorf("ToRomanNumeral(%d) should not return an error.", test.arabic)
			continue
		}
		if !test.hasError && actual != test.roman {
			t.Errorf("ToRomanNumeral(%d): %s, expected %s", test.arabic, actual, test.roman)
		}
	}
}

func BenchmarkRomanNumerals(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range romanNumeralTests {
			ToRomanNumeral(test.arabic)
		}
	}
}
