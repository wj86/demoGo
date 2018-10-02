package main

import "github.com/j123kaishichufa/goTestDemo/mytype"
import "github.com/j123kaishichufa/goTestDemo/mymath"
import "fmt"

type Shape interface {
    Length() float32
}

type Graph interface {
    Shape
}

func printRadiumOfCircle(cir mytype.Circle) {
    fmt.Println(cir.Radium)
}

func generateASquare() (*mytype.Square) {
    var sq mytype.Square = mytype.Square {4.0, mytype.Vertex2{2.0,3.0}}
    return &sq
}

func main() {
    var square mytype.Square = mytype.Square{3.0, mytype.Vertex2{1.0,2.0}}
    var circle mytype.Circle = mytype.Circle{3.0, mytype.Vertex2{1.0,2.0}}

    var s Shape = &square
    fmt.Println(s.Length()) //3*4 =12

    s = &circle
    fmt.Println(s.Length()) // 3.14 * 2 * 3 = 18.84
    
    var tmpSquare *mytype.Square = generateASquare()
    fmt.Println(tmpSquare.Length())  // 4*4 = 16
    fmt.Println(mymath.Pi)
    printRadiumOfCircle(circle)
}
