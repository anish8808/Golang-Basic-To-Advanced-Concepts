package main

import "math/rand"

func main() {
	//if you want to achive something that other programiing wihtout break it will run till default case
	var day uint8 = 5
	switch day {
	case 1:
		println("Sunday") //break is default in each case
	case 2:
		println("Monday")
	case 3:
		println("TuesDay")
	case 4:
		println("thursday")
	case 5:
		println("wednesday")
	case 6:
		println("friday")
	case 7:
		println("Saturday")
	default:
		println("Noday")
	}

	//or
	switch day := 3; day {
	case 1:
		println("Sunday") //break is default in each case
	case 2:
		println("Monday")
	case 3:
		println("TuesDay")
	case 4:
		println("thursday")
	case 5:
		println("wednesday")
	case 6:
		println("friday")
	case 7:
		println("Saturday")
	default:
		println("Noday")
	}
	for i := 0; i < 10; i++ {
		num := rand.Intn(2000)

		switch {
		case num >= 50 && num < 100:
			println("The number is grater then 50 and less then 100")
		case num >= 100 && num < 200:
			println("The number is grater then 100 and less then 200")
		default:
			println("num is greatrer then 200 or less then 50")
		}
	}
	str := "Hello Worlds"
	count := 0
	for _, char := range str {
		switch char {
		case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
			count++
			println(string(char), "is an Wovel")
		default:
			println(string(char), "is a constent")
		}
	}

	println("Number of wovels in a given String is ", count)

	var ch int32 = 65
	if ch == 'A' {
		println("YES")
	} else {
		println("NO")
	}
}
