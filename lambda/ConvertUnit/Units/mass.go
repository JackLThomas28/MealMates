package units

import "errors"

type Mass struct {
	unit string
}

// ******** METHODS *********
func (Mass) UnitType() string {
	return "MASS"
}

func (m Mass) SetUnit(unit string) error {
	if !IsMassUnit(unit) {
		// Invalid mass unit provided -> error exit
		return errors.New("Error: cannot set mass unit as " + unit)
	}

	m.unit = unit
	return nil
}

func (m Mass) Unit() string {
	return m.unit
}
// ******** ******* *********

func getMassUnits() [5]string {
	return [5]string{
		"pound",
		"ounce",
		"milligram",
		"gram",
		"kilogram",
	}
}

func IsMassUnit(unit string) bool {
	for _, mu := range getMassUnits() {
		if mu == unit {
			return true
		}
	}
	return false
}