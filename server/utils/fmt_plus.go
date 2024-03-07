package utils

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

// @function: RoundNumber
// @description: Rounds a given number to the specified number of decimal places.
//               If decimals is less than or equal to 0, it defaults to 2 decimal places.
// @param: num float64 - The number to be rounded.
// @param: decimals int - The number of decimal places to round to.
// @return: float64 - The rounded number.

func RoundNumber(num float64, decimals int) float64 {
	const defaultDecimals = 2

	if decimals <= 0 {
		decimals = defaultDecimals
	}

	// Calculate the multiplier using math.Pow10 for better precision
	multiplier := math.Pow10(decimals)

	// Round the number to the specified decimal places
	return math.Round(num*multiplier) / multiplier
}

// ToSnakeCase converts a string from CamelCase or PascalCase to snake_case.
func ToSnakeCase(s string) string {
	var result strings.Builder
	for i, char := range s {
		if i > 0 && char >= 'A' && char <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(char)
	}
	return strings.ToLower(result.String())
}

//@function: StructToMap
//@description: Convert a structure to a map using reflection
//@param: obj interface{}
//@return: map[string]interface{}

func StructToMap(obj interface{}) map[string]interface{} {
	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < objType.NumField(); i++ {
		if objType.Field(i).Tag.Get("mapstructure") != "" {
			data[objType.Field(i).Tag.Get("mapstructure")] = objValue.Field(i).Interface()
		} else {
			data[objType.Field(i).Name] = objValue.Field(i).Interface()
		}
	}
	return data
}

//@function: ArrayToString
//@description: Format an array into a string
//@param: array []interface{}
//@return: string

func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

func Pointer[T any](in T) (out *T) {
	return &in
}

func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func IndexOf(el uint, data []uint) int {
	for k, v := range data {
		if el == v {
			return k
		}
	}
	return -1 //not found.
}
