package globals

func GetUnitsMap() map[string]string {
	return map[string]string {
		// Volume Units
		"teaspoon":"teaspoon",
		"tsp":"teaspoon",
		"t":"teaspoon",
		"tablespoon":"tablespoon",
		"T":"tablespoon",
		"tbs":"tablespoon",
		"tbsp":"tablespoon",
		"fluid ounce":"fluid ounce",
		"fl oz":"fluid ounce",
		"gill":"gill",
		"cup":"cup",
		"c":"cup",
		"pint":"pint",
		"p":"pint",
		"pt":"pint",
		"fl pt":"pint",
		"quart":"quart",
		"q":"quart",
		"qt":"quart",
		"fl qt":"quart",
		"gallon":"gallon",
		"gal":"gallon",
		"milliliter":"milliliter", 
		"mL":"milliliter",
		"ml":"milliliter",
		"cc":"milliliter",
		"liter":"liter",
		"L":"liter",
		"l":"liter",
		"litre":"liter",
		"deciliter":"deciliter",
		"dL":"deciliter",
		"dl":"deciliter",
		"decilitre":"deciliter",
		// Mass Units
		"pound":"pound", 
		"lb":"pound", 
		"#":"pound",
		"ounce":"ounce", 
		"oz":"ounce",
		"milligram":"milligram", 
		"mg":"milligram",
		"gram":"gram", 
		"g":"gram",
		"kilogram":"kilogram", 
		"kg":"kilogram",
	}
}

func GetFoodContainers() [4]string {
	return [4]string{
		"jar",
		"can",
		"package",
		"container",
	}
}

func GetFracsInDecs() map[string]float64 {
	return map[string]float64{
		"¼": 0.25,
		"½": 0.5,
		"¾": 0.75,
		"⅐": 0.142857142857143,
		"⅑": 0.111111111111111,
		"⅒": 0.1,
		"⅓": 0.333333333333333,
		"⅔": 0.666666666666667,
		"⅕": 0.2,
		"⅖": 0.4,
		"⅗": 0.6,
		"⅘": 0.8,
		"⅙": 0.166666666666667,
		"⅚": 0.833333333333333,
		"⅛": 0.125,
		"⅜": 0.375,
		"⅝": 0.625,
		"⅞": 0.875,
	}
}
