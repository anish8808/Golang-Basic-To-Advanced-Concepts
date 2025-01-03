package main

//to create a package
//  1- GOPATH
//  < Go 1.11 {if you create a dir then you have to maintain this particular strcture bin/pkg/src and add to go path like export}
// example - -->export GOPATH=C:/Users/Anish/Workspace/MyProject
//also we use go deps

//then from version 1.11 onwords GO111MODULE=on

import (
	"fmt"
	cube "shapes/shape/cube"
	"shapes/shape/rect"
	sq "shapes/shape/sqaure"
)

func main() {

	// 1- standard packages
	//2- user defined packages
	//3-third party packages
	fmt.Println(rect.Greet())

	a1 := rect.Area(10.34, 13.67)
	fmt.Println("Area of rect:", a1)

	a2 := rect.Perimeter(10.34, 13.67)
	fmt.Println("Area of rect:", a2)

	rect.ShowTime()

	//square packages
	a3 := sq.Area(2)
	fmt.Println("area of sqaure :", a3)

	b3 := sq.Perimeter(2, 4)
	fmt.Println("area of Permitere :", b3)

	//cube packages

	a3 = cube.Area(1.0, 2.0, 30)
	fmt.Println("Area of a cube is :", a3)

	a3 = cube.Permitere(1.0, 2.0, 30)
	fmt.Println("Permeter of a cube :", a3)
}
