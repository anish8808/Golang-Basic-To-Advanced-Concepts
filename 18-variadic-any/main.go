package main

import "fmt"

func main() {
	fmt.Println(sumEle(10, 20, 30, 60))
	fmt.Println(sumEle(10.0, 20.0, 30.0, 60.0))
	sum, err := sumEle(10, 22, 4.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The sum is ", sum)
	}

}

func sumEle(nums ...interface{}) (interface{}, error) {

	var sum1 int
	var sum2 float64
	for _, val := range nums {
		switch val.(type) {
		case int:
			sum1 = sum1 + val.(int)
		case int64:
			sum1 = sum1 + int(val.(int64))
		case float64:
			sum2 = sum2 + val.(float64)
		case float32:
			sum2 = sum2 + float64(val.(float64))

		default:
			fmt.Errorf("Invalid")
		}
	}
	if sum1 != 0 {
		return sum1, nil
	} else if sum2 != 0 {
		return sum2, nil
	} else if sum1 != 0 && sum2 != 0 {
		return sum2 + float64(sum1), nil
	}
	return nil, nil
}
