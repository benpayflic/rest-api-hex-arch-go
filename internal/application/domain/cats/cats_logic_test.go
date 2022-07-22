package cats

import (
	"fmt"
	"testing"
)

func TestHumanYearsCalculator(t *testing.T) {
	catOne := NewCat("smudge", "british shorthair", "12/13/2020")
	catTwo := NewCat("button", "british shorthair", "03/09/2019")
	catThree := NewCat("ragnar", "ragdoll", "03/06/2022")

	var tests = []struct {
		c    *Cat
		want int
	}{
		{catOne, 7},
		{catTwo, 20},
		{catThree, 1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.c, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := tt.c.HumanYears()
			if ans != int64(tt.want) {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

}
