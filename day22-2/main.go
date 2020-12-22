package main

import (
	c "adventofcode2020/common"
	"fmt"
	"strings"
)

func seenVal(data []int) string {
	res := ""
	for _, v := range data {
		res += fmt.Sprintf("%d,", v)
	}
	return res
}
func play(player1, player2 []int) (bool, []int, []int) {

	seen1 := map[string]bool{}
	seen2 := map[string]bool{}

	for len(player1) > 0 && len(player2) > 0 {
		if _, exists := seen1[seenVal(player1)]; exists {
			return true, player1, player2
		}
		if _, exists := seen2[seenVal(player2)]; exists {
			return true, player1, player2
		}
		seen1[seenVal(player1)] = true
		seen2[seenVal(player2)] = true
		player1winner := true
		if player2[0] > player1[0] {
			player1winner = false
		}
		if len(player1)-1 >= player1[0] && len(player2)-1 >= player2[0] {
			p1n := []int{}
			for i := 1; i <= player1[0]; i++ {
				p1n = append(p1n, player1[i])
			}
			p2n := []int{}
			for i := 1; i <= player2[0]; i++ {
				p2n = append(p2n, player2[i])
			}
			player1winner, _, _ = play(p1n, p2n)
		}
		if player1winner {
			player1 = append(player1, player1[0])
			player1 = append(player1, player2[0])
		} else {
			player2 = append(player2, player2[0])
			player2 = append(player2, player1[0])
		}
		player1 = player1[1:]
		player2 = player2[1:]
	}
	return len(player1) > len(player2), player1, player2
}

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
	player1won, player1, player2 := play(player1, player2)
	win := []int{}
	if player1won {
		win = player1
	} else {
		win = player2
	}
	for i := range win {
		sum += (len(win) - i) * win[i]
	}
	c.Print(sum)

}
