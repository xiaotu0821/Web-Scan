package option

import (
	"regexp"
)

func FechRegex(Regepx string, body string) [][]string {
	reg, err := regexp.Compile(Regepx)
	if err != nil {
		MistakPrint("Compile Regepx failed")
		return nil
	}
	result := reg.FindAllStringSubmatch(string(body), -1)
	return result
}
