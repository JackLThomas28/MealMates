package converter


import (
	// "errors"
)

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

func GetConvRate() map[string]float64 {
	return map[string]float64{
		// Volume Units based on liters
		"teaspoon":0.004928931739603077,
		"tablespoon":0.014786800000070619843,
		"fluid ounce":0.029573600000141239685,
		"gill":0.12,
		"cup":0.24,
		"pint":0.473176,
		"quart":0.946353,
		"gallon":3.78541,
		"milliliter":0.001,
		"liter":1,
		// Mass Units based on grams
		"pound":453.592,
		"ounce":28.3495,
		"milligram":0.001,
		"gram":1,
		"kilogram":1000,
	}
}

func getMassUnits() [5]string {
	return [5]string{
		"pound",
		"ounce",
		"milligram",
		"gram",
		"kilogram",
	}
}

type Unit interface {
	getUnitType() string
	setUnit(unitId string)
	getUnit() string
}

type VolumeUnit struct {
	unit string
}

type MassUnit struct {
	unit string
}
func (VolumeUnit) getUnitType() string {
	return "VOLUME"
}

func (MassUnit) getUnitType() string {
	return "MASS"
}

func (v VolumeUnit) setUnit(unitId string) {
	v.unit = unitId
}

func (m MassUnit) setUnit(unitId string) {
	m.unit = unitId
}

func (v VolumeUnit) getUnit() string {
	return v.unit
}

func (m MassUnit) getUnit() string {
	return m.unit
}

func isVolumeUnit(unit string) bool {
	for _, vu := range getVolumeUnits() {
		if vu == unit {
			return true
		}
	}
	return false
}

func isMassUnit(unit string) bool {
	for _, mu := range getMassUnits() {
		if mu == unit {
			return true
		}
	}
	return false
} 
