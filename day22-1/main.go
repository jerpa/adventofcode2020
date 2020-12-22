package main

import (
	c "adventofcode2020/common"
	"strings"
)

func main() {
	sum := 0
	player1 := []int{}
	player2 := []int{}
	pl := 0
	for _, v := range c.ReadInputFile() {
		if len(v) == 0 {
			continue
		}
		if strings.HasPrefix(v, "Player") {
			pl++
		} else {
			if pl == 1 {
				player1 = append(player1, c.GetInt(v))
			} else {
				player2 = append(player2, c.GetInt(v))
			}
		}
	}
	for len(player1) > 0 && len(player2) > 0 {
		if player1[0] > player2[0] {
			player1 = append(player1, player1[0])
			player1 = append(player1, player2[0])
		} else {
			player2 = append(player2, player2[0])
			player2 = append(player2, player1[0])
		}
		player1 = player1[1:]
		player2 = player2[1:]
	}
	win := []int{}
	if len(player1) > len(player2) {
		win = player1
	} else {
		win = player2
	}
	for i := range win {
		sum += (len(win) - i) * win[i]
	}
	c.Print(sum)

}
