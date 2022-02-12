package convert

import (
	"errors"
	// "fmt"
	units "mealmates.com/lambda/ConvertUnit/Units"
)

type Source struct {
	// Ex: "convert(2).From("teaspoons").To("tablespoons")
	//		val = 2
	//		unit = "teaspoons"

	val  float64
	unit units.Unit
}

type Destination struct {
	// Ex: "convert(2).From("teaspoons").To("tablespoons")
	//		unit = "tablespoons"
	//		source = Source{val:2,unit:"teaspoons"}
	//		err = errors.New()

	unit   units.Unit
	source Source
	err    error
}

func Convert(val float64) Source {
	src := Source{val: val}
	return src
}

func (s Source) From(unit string) Destination {
	dest := Destination{}

	var err error
	s.unit, err = units.GetUnitObject(unit)
	if err != nil {
		dest.err = err
	}
	// if units.IsVolumeUnit(unit) {
	// 	s.unit = units.Volume{unit: unit}
	// } else if units.IsMassUnit(unit) {
	// 	s.unit = units.Mass{unit: unit}
	// } else {
	// 	// Invalid Source Unit -> error exit
	// 	dest = Destination{err: errors.New("Error: invalid 'From' unit provided: " + unit)}
	// }

	if dest.err == nil {
		dest.err = s.unit.SetUnit(unit)

		if dest.err == nil {
			dest.source = s
		}
	}
	return dest
}

func (d Destination) To(unit string) (float64, error) {
	var result float64 = 0

	if d.err == nil {
		d.unit, d.err = units.GetUnitObject(unit)

		// if units.IsVolumeUnit(unit) {
		// 	d.unit = units.Volume{unit: unitId}
		// } else if units.IsMassUnit(unitId) {
		// 	d.unit = units.Mass{unit: unitId}
		// } else {
		// 	d.err = errors.New("Error: invalid 'To' unit provided: " + unitId)
		// }

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
