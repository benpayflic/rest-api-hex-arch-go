package cats

import (
	"fmt"
	"testing"
)

func TestHumanYearsCalculator(t *testing.T) {
	catOne := NewCat()
	catTwo := NewCat()
	catThree := NewCat()

	catOne.Name = "smudge"
	catOne.Breed = "british shorthair"
	catOne.DOB = "12/13/2020"

	catTwo.Name = "button"
	catTwo.Breed = "british shorthair"
	catTwo.DOB = "03/09/2019"

	catThree.Name = "ragnar"
	catThree.Breed = "ragdoll"
	catThree.DOB = "03/06/2022"

	var tests = []struct {
		c    *Cat
		want int
	}{
		{catOne, 7},
		{catTwo, 20},
		{catThree, 1},
	}

	cs := NewCatService()

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.c, tt.want)
		t.Run(testname, func(t *testing.T) {
			cs.CalculateHumanYears(tt.c)
			if tt.c.HumanYears != int64(tt.want) {
				t.Errorf("got %d, want %d", tt.c.HumanYears, tt.want)
			}
		})
	}

}
