package cats

import "github.com/benpayflic/rest-api-hex-arch-go/pkg/utils"

type CatService struct {
}

func NewCatService() *CatService {
	return &CatService{}
}

// Calculate the age of the cat in human years
func (CatService) CalculateHumanYears(c *Cat) {
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
	c.HumanYears = humanYears
}
