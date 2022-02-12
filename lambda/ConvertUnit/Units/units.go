package units

import "errors"

func GetConvRate() map[string]float64 {
	return map[string]float64{
		// Volume Units based on liters
		"teaspoon":    0.004928931739603077,
		"tablespoon":  0.014786800000070619843,
		"fluid ounce": 0.029573600000141239685,
		"gill":        0.12,
		"cup":         0.24,
		"pint":        0.473176,
		"quart":       0.946353,
		"gallon":      3.78541,
		"milliliter":  0.001,
		"liter":       1,
		// Mass Units based on grams
		"pound":     453.592,
		"ounce":     28.3495,
		"milligram": 0.001,
		"gram":      1,
		"kilogram":  1000,
	}
}

type Unit interface {
	UnitType() string
	SetUnit(unitId string) error
	Unit() string
}

func GetUnitObject(unit string) (Unit, error) {
	if IsMassUnit(unit) {
		return Mass{unit: unit}, nil
	} else if IsVolumeUnit(unit) {
		return Volume{unit: unit}, nil
	} else {
		// Invalid Source Unit -> error exit
		return Mass{}, errors.New("Error: invalid 'From' unit provided: " + unit)
	}
}