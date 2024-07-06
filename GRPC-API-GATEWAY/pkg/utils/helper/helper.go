package helper

import (
	"errors"
	"strconv"
	"strings"
)

func ConvertStringToArray(inputs []string) ([]int64, error) {
	var intArray []int64
	for _, input := range inputs {
		cleanedInput := strings.ReplaceAll(input, " ", "")
		strElements := strings.Split(cleanedInput, ",")
		for _, strElement := range strElements {
			if strElement == "" {
				continue
			}
			intValue, err := strconv.ParseInt(strElement, 10, 64)
			if err != nil {
				return nil, errors.New("invalid integer value: " + strElement)
			}
			intArray = append(intArray, intValue)
		}
	}
	return intArray, nil
}
