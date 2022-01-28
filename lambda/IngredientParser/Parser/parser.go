package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	globals "mealmates.com/lambda/IngredientParser/Globals"
)

const RE_AMT_PATTERN = "[0-9¼½¾⅐⅑⅒⅓⅔⅕⅖⅗⅘⅙⅚⅛⅜⅝⅞.]+"
const RE_PAR_PATTERN = "[()]"

type Ingredient struct {
	Name string
	Amount float64
	Unit string
	Raw string
}

func printIngredient(i Ingredient) {
	fmt.Println("Name:  ", i.Name)
	fmt.Println("Amount:", i.Amount)
	fmt.Println("Unit:  ", i.Unit)
	fmt.Println("-------------------")
}

func convertStringToFloat(strAmount string) (float64, error) {
	var err error = nil

	fracToDecMap := globals.GetFracsInDecs()
	amt, found := fracToDecMap[strAmount]

	// Try to parse a decimal num if not a fraction
	if !found {
		amt, err = strconv.ParseFloat(strAmount, 64)
	}

	return amt, err
}

func calculateAmount(rawIngredient string) (float64, error) {
	amount := strings.Split(rawIngredient, " ")[0]

	// Check for decimals, numbers, fractions
	// Ex: 10.75, ½, 1, 5 ½
	reAmt, _ := regexp.Compile(RE_AMT_PATTERN)
	resultAmt := reAmt.FindAllIndex([]byte(amount), -1)

	// Get any string num in decimal
	var amt float64 = 0
	for i := 0; i < len(resultAmt); i++ {
		// Get the slice containing the string num
		strAmount := amount[resultAmt[i][0]:resultAmt[i][1]]
		// Convert from string to float
		flt, err := convertStringToFloat(strAmount)
		if err != nil {
			return 0, err
		}

		amt = amt + flt
	}

	// Find any parentheses
	rePar, _ := regexp.Compile(RE_PAR_PATTERN)
	resultPar := rePar.FindAllIndex([]byte(rawIngredient), 2)

	if len(resultAmt) != 0 && len(resultPar) != 0 {
		lastIndexOfAmt := resultAmt[0][1]
		firstIndexOfParen := resultPar[0][0]
		// Check to make sure that the parentheses are immediately after amount
		// Good Ex: 1 (3 pound) chickent
		// Bad  Ex: 3 cups cheese (shredded)
		if lastIndexOfAmt == firstIndexOfParen - 1 {
			// Get the text from inside the parentheses
			splitParText := strings.Split(rawIngredient[resultPar[0][1]:resultPar[1][0]], " ")
			// The amount should be the first thing from parentheses
			// Ex: (3 pound)
			//      ^
			parAmt, err := strconv.ParseFloat(splitParText[0], 64)
			if err != nil {
				return parAmt, err
			}
			amt = amt * parAmt
		}
	}
	if amt == 0 {
		amt = 1
	}
	return amt, nil
}

func getUnit(rawIngredient string) (string, error) {
	for word, unit := range globals.GetUnitsMap() {
		// Find the unit. May be plural
		// Ex: pounds
		re, _ := regexp.Compile(`(\b` + word + `)(s\b|\b)`)
		result := re.FindAllIndex([]byte(rawIngredient), -1)
	
		if len(result) != 0 {
			return unit, nil
		}
	}
	// If we didn't find it, return generic
	return "unit", nil
}

func determineName(rawIngredient string) (string, error) {
	re, _ := regexp.Compile(",")
	commas := re.FindAllIndex([]byte(rawIngredient), -1)

	name := rawIngredient
	if len(commas) > 0 {
		// Remove everything after the last comma
		// Ex: 2 pounds skinless, boneless chicken breast halves, cut into 1/2-inch cubes
		//                                                      ^^^^^^^^^^^^^^^^^^^^^^^^^
		name = rawIngredient[:commas[len(commas) - 1][0]]
	}

	unitFound := false
	for word, _ := range globals.GetUnitsMap() {
		re, _ = regexp.Compile(`(\b` + word + `)(s\b|\b)`)
		units := re.FindAllIndex([]byte(name), 1)

		if len(units) != 0 {
			if name[units[0][1] + 1] == byte(' ') {
				name = name[units[0][1] + 2:]
			} else {
				name = name[units[0][1] + 1:]
			}
			unitFound = true
			break
		}
	}

	// If there was no unit, remove the amount (if it is there)
	if !unitFound {
		amount := strings.Split(rawIngredient, " ")[0]
		re, _ = regexp.Compile(RE_AMT_PATTERN)
		amts := re.FindAllIndex([]byte(amount), -1)

		// Only remove the amount if it was there
		if len(amts) > 0 {
			// Find the first space and remove everything before it
			re, _ = regexp.Compile(" ")
			spaces := re.FindAllIndex([]byte(name), 1)
	
			name = name[spaces[0][1]:]
		}
	}

	// Remove any food containers
	for _, container := range globals.GetFoodContainers() {
		re, _ = regexp.Compile(`(\b` + container + `)(s\b|\b)`)
		containers := re.FindAllIndex([]byte(name), -1)

		for _, indeces := range containers {
			if name[indeces[1] + 1] == byte(' ') {
				name = name[indeces[1] + 2:]
			} else {
				name = name[indeces[1] + 1:]
			}
		}
	}
	return name, nil
}

func ParseIngredients(ingredients []string) ([]Ingredient, error) {
	// 1. Store ingredients in Ingredient objects
	var ingredientList []Ingredient
	for i := 0; i < len(ingredients); i++ {
		curIngr := Ingredient{
			Raw: ingredients[i],
		}
		ingredientList = append(ingredientList, curIngr)
	}

	var err error
	for i := 0; i < len(ingredientList); i++ {
		// 2. Set the amount
		ingredientList[i].Amount, err = calculateAmount(ingredientList[i].Raw)
		if err != nil {
			return nil, err
		}

		// 3. Set the unit
		ingredientList[i].Unit, err = getUnit(ingredientList[i].Raw)
		if err != nil {
			return nil, err
		}

		// 4. Set the name
		ingredientList[i].Name, err = determineName(ingredientList[i].Raw)
		if err != nil {
			return nil, err
		}

		// printIngredient(ingredientList[i])
	}

	return ingredientList, err
}