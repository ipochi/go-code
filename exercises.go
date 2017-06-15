package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//isLeap()
	//hammingDistance()
	//raindrops()
	//squareAList()
	//acronym()
	englishPangram()
}

func englishPangram() {
	phrases := []string{
		"",
		"The quick brown fox jumps over the lazy dog",
		"a quick movement of the enemy will jeopardize five gunboats",
		"the quick brown fish jumps over the lazy dog",
		"the 1 quick brown fox jumps over the 2 lazy dogs",
		"7h3 qu1ck brown fox jumps ov3r 7h3 lazy dog",
	}
	for _, val := range phrases {
		lval := strings.ToLower(val)

		fmt.Println(val)
		var arr [26]bool

		for _, c := range lval {
			if c >= 'a' && c <= 'z' {
				arr[c-'a'] = true
			}
		}

		isPanagram := true
		for _, ans := range arr {
			if ans == false {
				isPanagram = false
				break
			}
		}

		if isPanagram {
			fmt.Println(val, " is a Pangram")
		} else {
			fmt.Println(val, " is NOT a Pangram")

		}
	}
}

func acronym() {
	fmt.Println("Enter a Long Phrase")

	in := bufio.NewReader(os.Stdin)
	phrase, _ := in.ReadString('\n')

	x := strings.Fields(phrase)

	var abb string
	for _, val := range x {
		abb = abb + val[:1]
	}

	fmt.Println("Abbreviation is ::: ", strings.ToUpper(abb))
}

func squareAList() {
	f := func(x int64) int64 {
		return x * x
	}

	a := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, val := range a {
		a[i] = f(val)
	}

	fmt.Println("Squared list ::: ", a)
}

func raindrops() {
	var num int

	fmt.Println("Enter a number")
	fmt.Scanf("%d", &num)

	switch true {
	case num%3 == 0:
		fmt.Println("Pling")

	case num%5 == 0:
		fmt.Println("Plang")

	case num%7 == 0:
		fmt.Println("Plong", num%7)
	default:
		fmt.Println(num)
	}
}

func hammingDistance() {
	var s1, s2 string

	fmt.Println("Enter the first string")
	fmt.Scanln(&s1)

	fmt.Println("Enter the second string")
	fmt.Scanln(&s2)

	if len(s1) != len(s2) {
		fmt.Println("Strings length not same")
	} else {
		count := 0
		for i, _ := range s1 {
			if s1[i] != s2[i] {
				count++
			}
		}

		fmt.Println("Hamming ditance  ::: ", count)
	}
}

func isLeap() {
	var year string
	fmt.Println("Enter the number : ")
	fmt.Scanf("%s", &year)

	if _, err := strconv.Atoi(year); err != nil {
		fmt.Println(year + "Doesn't look like a number")
		fmt.Println(err)
	} else {
		y, _ := strconv.Atoi(year)
		if y%4 == 0 {
			fmt.Println("Leap Year !!!")
		} else {
			fmt.Println("Not a Leap Year.")
		}
	}
}
