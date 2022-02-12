package convert

import units "mealmates.com/lambda/ConvertUnit/Units"

type Source struct {
	// Ex: "convert(2).From("teaspoons").To("tablespoons")
	//		val = 2
	//		unit = "teaspoons"

	val  float64
	unit units.Unit
}

func (s Source) From(unit string) Destination {
	dest := Destination{}

	var err error
	s.unit, err = units.GetUnitObject(unit)
	if err != nil {
		dest.err = err
	}

	if dest.err == nil {
		dest.err = s.unit.SetUnit(unit)

		if dest.err == nil {
			dest.source = s
		}
	}
	return dest
}
