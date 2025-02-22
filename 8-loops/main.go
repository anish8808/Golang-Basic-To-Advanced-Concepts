package main

//anish

func main() {
	for i := 0; i < 10; i++ {
		print(i, " ")
	}

	j := 0
	for j < 10 {
		print(j, " ")
		j++
	}

	k := 0
	for {
		if k > 9 {
			break
		}
		print(k, " ")
		k++
	}
	println()
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			print(i, j, " ")
		}
		println()
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == j {
				break // it will break the outer loop
			}
			print(i, j, " ")
		}
		println()
	}

outer:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == j {
				break outer // this way we can jump from any to this outer statements
			}
			print(i, j, " ")
		}
		println()
	}
}
