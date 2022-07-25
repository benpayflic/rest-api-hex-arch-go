package cats

import "encoding/json"

type Cat struct {
	Name       string `json:"name,omitempty"`
	Breed      string `json:"breed,omitempty"`
	DOB        string `json:"dob,omitempty"`
	HumanYears int64  `json:"humanYears,omitempty"`
}

func UnmarshalCat(data []byte) (Cat, error) {
	var r Cat
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Cat) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func NewCat() *Cat {
	return &Cat{}
}
