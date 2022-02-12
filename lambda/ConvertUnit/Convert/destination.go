package convert

import (
	"errors"

	// Local Packages
	units "mealmates.com/lambda/ConvertUnit/Units"
)

type Destination struct {
	// Ex: "convert(2).From("teaspoons").To("tablespoons")
	//		unit = "tablespoons"
	//		source = Source{val:2,unit:"teaspoons"}
	//		err = errors.New()

	unit   units.Unit
	source Source
	err    error
}

func (d Destination) To(unit string) (float64, error) {
	var result float64 = 0

	if d.err == nil {
		d.unit, d.err = units.GetUnitObject(unit)

		if d.err == nil {
			d.err = d.unit.SetUnit(unit)

			if d.err == nil {
				destUnitType := d.unit.UnitType()
				srcUnitType := d.source.unit.UnitType()
				if srcUnitType != destUnitType {
					// Mismatching Unit Types (Ex: Vol to Mass) -> error exit
					d.err = errors.New("Error: mismatching unit types. Cannot convert from " + srcUnitType + " to " + destUnitType)
				}
			}
		}

		// Perform the conversion calculation
		if d.err == nil {
			temp := d.source.val * units.GetConvRate()[d.source.unit.Unit()]
			result = temp / units.GetConvRate()[d.unit.Unit()]
		}
	}

	return result, d.err
}
