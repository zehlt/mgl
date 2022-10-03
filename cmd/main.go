package main

import (
	"log"

	"github.com/zehlt/mgl"
)

func main() {
	log.Println("hello mgl")

	v1 := mgl.Vector3f{X: 10, Y: 20, Z: 30}

	log.Println(v1)
	log.Println(v1.Neg())
	log.Println(v1.Norm())
	log.Println(v1.Add(v1))
	log.Println(v1)

}
