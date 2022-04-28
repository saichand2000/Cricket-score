package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hi")

	reader := bufio.NewReader(os.Stdin)
	var score int
	wickets := 0
	src := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"}
	indiva := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	for ball := 0; ball < 50; ball++ {

		if wickets == 10 {

			fmt.Println("Game Over", wickets)
			fmt.Println("Score =", score)
			break

		} else {
			guess := 0
			fmt.Println("----------------------")
			fmt.Println("Enter ball , Left ball", 50-ball)
			fmt.Println("----------------------")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			switch os := input; os {

			//No-balls
			case "w", "n":
				fmt.Println("wide/noball")
				score = 1 + score
				ball--

			//Wickets
			case "l", "c", "r", "b":
				fmt.Println("lbw/catch/runout/bowled")
				wickets++
				fmt.Println("Wickets Felldown", wickets)
				playersOut(src, wickets)
				eachplayer(indiva, wickets, score)

			//Runs
			case "2", "4", "6", "0":
				guess, _ = strconv.Atoi(input)
				score = guess + score
				fmt.Println("Score =", score)

			case "1", "3":
				guess, _ = strconv.Atoi(input)
				score = guess + score
				fmt.Println("Score =", score)
				changePlayer(src)

			default:
				fmt.Println("Enter Correct Command/Runs")
				ball--

			}
		}

	}
	fmt.Println("Target=", score+1)
	fmt.Println(indiva)
}

func playersOut(arr []string, wick int) {

	out := arr[0]
	swapF := reflect.Swapper(arr)
	swapF(0, wick+1)
	fmt.Println(out, "is OUT")
	fmt.Println("Current Players:", arr[0], arr[1])

}

func changePlayer(arr []string) {

	swapF := reflect.Swapper(arr)
	swapF(0, 1)
	fmt.Println("Current Players:", arr[0], arr[1])
}

func eachplayer(arr1 []int, wick int, sco int) {
	sco1 := 0
	arr1[wick] = sco - sco1
	fmt.Println(arr1[0])
}
