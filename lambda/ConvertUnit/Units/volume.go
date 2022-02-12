package units

import "errors"

type Volume struct {
	unit string
}

// ******** METHODS *********
func (Volume) UnitType() string {
	return "VOLUME"
}

func (v Volume) SetUnit(unit string) error {
	if !IsVolumeUnit(unit) {
		// Invalid volume unit provided -> error exit
		return errors.New("Error: cannot set volume unit as " + unit)
	}

	v.unit = unit
	return nil
}

func (v Volume) Unit() string {
	return v.unit
}
// ******** ******* *********

func getVolumeUnits() [11]string {
	return [11]string{
		"teaspoon",
		"tablespoon",
		"fluid ounce",
		"gill",
		"cup",
		"pint",
		"quart",
		"gallon",
		"milliliter",
		"liter",
	}
}

func IsVolumeUnit(unit string) bool {
	for _, vu := range getVolumeUnits() {
		if vu == unit {
			return true
		}
	}
	return false
}
