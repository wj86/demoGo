package mytype

import (
    math "github.com/j123kaishichufa/goTestDemo/mymath"
)

type Circle struct  {
    Radium float32
    Vertex2  //center point of a cricle
}

type Square struct  {
    Edge float32
    Vertex2 //center point of a square
}

type Vertex2 struct{
    X float32
    Y float32
}

func (c *Circle) Length () float32 {
    return math.Multiply( math.Pi, c.Radium * 2)
}

func (s *Square) Length () float32 {
    return math.Multiply( float32(math.FourNumber), s.Edge )
}
