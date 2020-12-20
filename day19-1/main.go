package main

import (
	c "adventofcode2020/common"
	"strings"
)

type rule struct {
	ch    string
	rules [][]string
}

var rules map[string]rule

func ruleChecker(data, rule string, pos int) (bool, int) {
	if pos >= len(data) {
		return false, -1
	}
	r := rules[rule]
	if len(r.rules) == 0 {
		if string(data[pos]) == r.ch {
			return true, pos + 1
		}
	} else {
		for _, v := range r.rules {
			p := pos
			ok := false
			for _, t := range v {
				ok, p = ruleChecker(data, t, p)
				if !ok {
					break
				}
			}
			if ok {
				return true, p
			}
		}
	}
	return false, -1
}

func main() {
	sum := 0
	mode := "rules"
	rules = map[string]rule{}

	for _, v := range c.ReadInputFile() {
		if len(v) == 0 {
			mode = "check"
		} else {
			if mode == "rules" {
				ru := rule{}
				s := strings.Split(v, ": ")
				r := s[0]
				if s[1][0] == '"' {
					ru.ch = string(s[1][1])
				} else {
					for _, p := range strings.Split(s[1], "|") {
						l := []string{}
						for _, li := range strings.Split(strings.TrimSpace(p), " ") {
							l = append(l, li)
						}
						ru.rules = append(ru.rules, l)
					}
				}
				rules[string(r)] = ru
			} else {
				if ok, pos := ruleChecker(v, "0", 0); ok && pos == len(v) {
					sum++
				}
			}
		}
	}
	c.Print(sum)

}
