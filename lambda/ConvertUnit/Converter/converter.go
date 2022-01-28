package converter

import (
	"errors"
	// "fmt"
	// "mealmates.com/lamda/ConvertUnit/Converter"
)

type Source struct {
	val  float64
	unit Unit
}

type Destination struct {
	val    float64
	unit   Unit
	source Source
	err    error
}

func Convert(val float64) Source {
	// Pass along the value
	src := Source{val: val}
	return src
}

func (s Source) From(unitId string) Destination {
	var dest Destination

	if isVolumeUnit(unitId) {
		s.unit = VolumeUnit{unit: unitId}
	} else if isMassUnit(unitId) {
		s.unit = MassUnit{unit: unitId}
	} else {
		dest = Destination{err: errors.New("Error: invalid 'From' unit provided: " + unitId)}
	}

	if dest.err == nil {
		s.unit.setUnit(unitId)

		dest = Destination{val: s.val, source: s}
	}
	return dest
}

func (d Destination) To(unitId string) (float64, error) {
	var result float64 = 0

	if d.err == nil {
		if isVolumeUnit(unitId) {
			d.unit = VolumeUnit{unit: unitId}
		} else if isMassUnit(unitId) {
			d.unit = MassUnit{unit: unitId}
		} else {
			d.err = errors.New("Error: invalid 'To' unit provided: " + unitId)
		}

		if d.err == nil {
			d.unit.setUnit(unitId)

			destUnitType := d.unit.getUnitType()
			srcUnitType := d.source.unit.getUnitType()
			if srcUnitType != destUnitType {
				result = 0
				d.err = errors.New("Error: mismatching unit types. Cannot convert from " + srcUnitType + " to " + destUnitType)
			}
		}

		if d.err == nil {
			temp := d.val * GetConvRate()[d.source.unit.getUnit()]
			result = temp / GetConvRate()[d.unit.getUnit()]
		}
	}

	return result, d.err
}
