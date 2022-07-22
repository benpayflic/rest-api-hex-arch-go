package cats

import "github.com/benpayflic/rest-api-hex-arch-go/pkg/utils"

type Cat struct {
	Name  string
	Breed string
	DOB   string
}

func NewCat(name string, breed string, dob string) *Cat {
	return &Cat{
		Name:  name,
		Breed: breed,
		DOB:   dob,
	}
}

// Calculate the age of the cat in human years
func (c *Cat) HumanYears() int64 {
	age := utils.Age(c.DOB)
	var humanYears int64

	if age > 6 {
		age = age - 6
		age = age * 4
		humanYears = int64(age * 40)
	} else {
		age = age * 19
		age = age / 3
		humanYears = int64(age + 1)
	}
	return humanYears
}
