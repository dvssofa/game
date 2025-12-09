package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)

	score1 := 0
	score2 := 0

	for {
		fmt.Print("Игрок 1, нажмите Enter")
		scanner.Scan()
		dice := rand.Intn(6) + 1
		score1 += dice
		fmt.Println("Выпало:", dice, "Всего:", score1)

		if score1 >= 100 {
			fmt.Println("Игрок 1 победил!")
			break
		}

		fmt.Print("Игрок 2, нажмите Enter")
		scanner.Scan()
		dice = rand.Intn(6) + 1
		score2 += dice
		fmt.Println("Выпало:", dice, "Всего:", score2)

		if score2 >= 100 {
			fmt.Println("Игрок 2 победил!")
			break
		}
	}
}
