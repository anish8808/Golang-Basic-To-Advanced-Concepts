package square //must the name of immediate dir

import (
	"fmt" // there are no access specifires ,
	"math/rand"
	"time"
)

// only a concept exported or not exported
// any memeber/type start with uppercase exportedþ
// /any memeber/type start with lowercase is unexported(local only to package)
func Greet() string {
	return "Hello World"
}

func ShowTime() {
	fmt.Println(time.Now())
}
func Area(s float32) float32 {
	return s * s
}

func Perimeter(l, b float32) float32 {
	return 2 * (l + b)
}

// this is not exported function beacuse it start with lowercase
func sayHey() {
	fmt.Println("Hey how are you doing  --> This is your token number-->", rand.Intn(2011))
}
