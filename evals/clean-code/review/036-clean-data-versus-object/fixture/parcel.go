package fixture

import "errors"

const maxWeightGrams = 30000

var ErrTooHeavy = errors.New("validate parcel: weight exceeds the carrier limit")

type Address struct {
	Street     string
	City       string
	PostalCode string
}

type Parcel struct {
	weightGrams int
	destination Address
}

func NewParcel(weightGrams int, destination Address) Parcel {
	return Parcel{weightGrams: weightGrams, destination: destination}
}

func (p Parcel) Validate() error {
	if p.weightGrams > maxWeightGrams {
		return ErrTooHeavy
	}
	return nil
}

func (p Parcel) IsBoundFor(city string) bool {
	return p.destination.City == city
}
