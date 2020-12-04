package main

import (
	c "adventofcode2020/common"
	"regexp"
	"strconv"
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

	byr, err := strconv.Atoi(data["byr"])
	if err != nil {
		return false
	}
	if byr < 1920 || byr > 2002 {
		return false
	}

	iyr, err := strconv.Atoi(data["iyr"])
	if err != nil {
		return false
	}
	if iyr < 2010 || iyr > 2020 {
		return false
	}

	eyr, err := strconv.Atoi(data["eyr"])
	if err != nil {
		return false
	}
	if eyr < 2020 || eyr > 2030 {
		return false
	}

	if len(data["hgt"]) < 4 {
		return false
	}
	if data["hgt"][len(data["hgt"])-2:] != "cm" && data["hgt"][len(data["hgt"])-2:] != "in" {
		return false
	}
	hgt, err := strconv.Atoi(data["hgt"][0 : len(data["hgt"])-2])
	if err != nil {
		return false
	}
	if data["hgt"][len(data["hgt"])-2:] == "cm" && (hgt < 150 || hgt > 193) {
		return false
	}
	if data["hgt"][len(data["hgt"])-2:] == "in" && (hgt < 59 || hgt > 76) {
		return false
	}

	if ok, _ := regexp.MatchString(`\A\#[0-9a-z]{6}$`, data["hcl"]); !ok {
		return false
	}

	if data["ecl"] != "amb" && data["ecl"] != "blu" && data["ecl"] != "brn" && data["ecl"] != "gry" && data["ecl"] != "grn" && data["ecl"] != "hzl" && data["ecl"] != "oth" {
		return false
	}

	if ok, _ := regexp.MatchString(`\A[0-9]{9}$`, data["pid"]); !ok {
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
