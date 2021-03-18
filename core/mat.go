//package core
package main

import (
	"image"
	"fmt"
)

// encapsulate the image.Image interface.
// See https://golang.org/pkg/image/ for details of the image package
type MatImage struct {
	image.Image
}


func (m *MatImage) AsImage() image.Image {
	return m
}

func (m *MatImage) Shape() (int, int) {
	return m.Bounds().Max.X, m.Bounds().Max.Y
}

// change of plans. Don't wrap image.Image
// use regular nested arrays and array slices
// only use the image module for compatibility and image loading / display
// eventually consider building custom codecs
type Mat struct {
	data [][][]int // row-column depth(aka channel) format
}
// IMAGES! There is a standard image package!!!
//my_image := image.NewRGBA(image.Rect(0, 0, 100, 100)) // need to pass it a rectangle
//fmt.Println(my_image.Bounds())         							 // returns top_left & bottom_right pixel coordinates
//fmt.Println(my_image.At(0, 0))                        // pixel at location (0,0)
//fmt.Printf("%T/n", my_image.At(0, 0))                 // pixel type is called color.RGBA
//fmt.Println(my_image.At(0, 0).RGBA())                 // RGBA values of specific pixel

//constructor functions

//one of these ways of constructing is probably bad, but I don't know which
func Eye(size int) *Mat {
	return new(Mat)
}

func Eye2(size int) Mat {
	m := new(Mat)
	m.data = make([][][]int, size)
	return *m
}


func Ones(rows int, cols int) Mat {
	return *new(Mat)
}


func main() {
	a := Eye(3)
	fmt.Printf("'t' type for *Mat: %t \n", a)
	fmt.Printf("'T' type for *Mat: %T \n", a)


	b := Eye2(3)
	fmt.Printf("'t' type for Mat: %t \n", b)
	fmt.Printf("'T' type for Mat: %T \n", b)
}