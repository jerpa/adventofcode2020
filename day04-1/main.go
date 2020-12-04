package main

import (
	c "adventofcode2020/common"
	"strings"
)

func validate(data map[string]string) bool {
	if _, ok := data["byr"]; !ok {
		return false
	}
	if _, ok := data["iyr"]; !ok {
		return false
	}
	if _, ok := data["eyr"]; !ok {
		return false
	}
	if _, ok := data["hgt"]; !ok {
		return false
	}
	if _, ok := data["hcl"]; !ok {
		return false
	}
	if _, ok := data["ecl"]; !ok {
		return false
	}
	if _, ok := data["pid"]; !ok {
		return false
	}

	return true
}

func main() {
	sum := 0
	data := c.ReadInputFile()
	vals := map[string]string{}

	for _, v := range data {
		if len(strings.TrimSpace(v)) == 0 {
			if validate(vals) {
				sum++
			}
			vals = map[string]string{}
		} else {
			for _, w := range strings.Split(v, " ") {
				k := strings.Split(w, ":")
				vals[k[0]] = k[1]
			}
		}
	}
	if validate(vals) {
		sum++
	}

	c.Print(sum)

}
