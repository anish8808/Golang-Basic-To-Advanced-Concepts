package rect

import (
	"fmt"
	"math/rand"
	"time"
)

// this is not exported function beacuse it start with lowercase
func sayHey() {
	fmt.Println("Hey How are you doing :", rand.Intn(2011))
}

func Greet() string {
	return "Hello World"
}

func ShowTime() {
	fmt.Println(time.Now())
}
