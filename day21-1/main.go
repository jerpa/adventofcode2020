package main

import (
	c "adventofcode2020/common"
	"strings"
)

type ingredient struct {
	allergens        map[string]bool
	possibleAllergen map[string]bool
	count            int
}

func main() {
	sum := 0
	g := c.Graph{}
	for _, v := range c.ReadInputFile() {
		v = strings.TrimRight(v, ")")
		data := strings.Split(v, " (contains ")
		rec, _ := g.AddNode(c.GraphNode{Type: "Food"})
		for _, in := range strings.Split(data[0], " ") {
			inNode, err := g.GetNodeByID(in)
			if err != nil {
				inNode, _ = g.AddNode(c.GraphNode{ID: in, Type: "Ingredient"})
			}
			rec.AddEdge("Ingredient", inNode)
			for _, a := range strings.Split(data[1], ", ") {
				all, err := g.GetNodeByID(a)
				if err != nil {
					all, _ = g.AddNode(c.GraphNode{ID: a, Type: "Allergy"})
				}
				inNode.AddEdge("Possible", all)
			}
		}
		for _, a := range strings.Split(data[1], ", ") {
			all, err := g.GetNodeByID(a)
			if err != nil {
				all, _ = g.AddNode(c.GraphNode{ID: a, Type: "Allergy"})
			}
			rec.AddEdge("Allergy", all)
		}
	}

	// Ingredient=allergy if it exists every time that allergy is mentioned
	for _, a := range g.GetNodesByType("Allergy") {
		num := len(a.GetEdgesFrom("Allergy"))
		c := map[*c.GraphNode]int{}
		for _, in := range a.GetEdgesFrom("Possible") {
			c[in.From]++
		}
		for k, v := range c {
			if v == num {
				g.RemoveNode(k)
			}
		}
	}

	for _, in := range g.GetNodesByType("Ingredient") {
		sum += len(in.GetEdgesFrom("Ingredient"))
	}
	c.Print(sum)
}
