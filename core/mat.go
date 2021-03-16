package core

import (
	"image"
)

// encapsulate the image.Image interface.
// See https://golang.org/pkg/image/ for details of the image package
type Mat struct {
	image.Image
}

func (m *Mat) AsImage() image.Image {
	return m
}

func (m *Mat) Shape() (int, int) {
	return m.Bounds().Max.X, m.Bounds().Max.Y
}


// IMAGES! There is a standard image package!!!
//my_image := image.NewRGBA(image.Rect(0, 0, 100, 100)) // need to pass it a rectangle
//fmt.Println(my_image.Bounds())         							 // returns top_left & bottom_right pixel coordinates
//fmt.Println(my_image.At(0, 0))                        // pixel at location (0,0)
//fmt.Printf("%T/n", my_image.At(0, 0))                 // pixel type is called color.RGBA
//fmt.Println(my_image.At(0, 0).RGBA())                 // RGBA values of specific pixel
