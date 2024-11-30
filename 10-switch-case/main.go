package main

func main() {
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
	switch day := 4; day {
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
}
