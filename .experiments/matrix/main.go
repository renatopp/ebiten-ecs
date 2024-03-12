package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

func printMatrix(mat ebiten.GeoM) {
	fmt.Printf("[%01.1f %01.1f %01.1f]\n", mat.Element(0, 0), mat.Element(0, 1), mat.Element(0, 2))
	fmt.Printf("[%01.1f %01.1f %01.1f]\n", mat.Element(1, 0), mat.Element(1, 1), mat.Element(1, 2))
	fmt.Printf("[%01.1f %01.1f %01.1f]\n", 0., 0., 1.)
}

func main() {
	a := ebiten.GeoM{}
	a.Scale(2, 5)
	a.Translate(10, 20)

	// b := ebiten.GeoM{}
	// b.Scale(3, 4)
	// b.Translate(5, 7)

	printMatrix(a)
	// printMatrix(b)
	// a.Concat(b)

	// printMatrix(a)
}
