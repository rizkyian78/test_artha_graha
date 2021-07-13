package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	player := 3
	dice := 4
	var obj = map[string]int{}
	var i int
	for i = 1; i <= player; i++ {
		obj[fmt.Sprintf("%d", i)] = 0
	}
	var k int = 1
	rand.Seed(time.Now().UnixNano())
	var diceCounter int = dice
	for k <= player {
		if diceCounter > 0 {
			var j int
			for j = 1; j <= dice; j++ {
				var randomDice = rand.Intn(7)
				if randomDice == 1 {
					diceCounter--
				} else if randomDice == 6 {
					obj[fmt.Sprintf("%d", k)] += 1
					diceCounter--
				}
			}
		} else {
			k++
			diceCounter += dice
		}
	}
	var maxValue int = 0
	for key, value := range obj {
		if value > maxValue {
			value, _ = strconv.Atoi(key)
			maxValue = value
		}
	}
	fmt.Println(maxValue)
}
