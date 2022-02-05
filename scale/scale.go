// Package scale implements well known scaling alrogithms for pixel art images
// Algorithms taken from https://en.wikipedia.org/wiki/Pixel-art_scaling_algorithms
package scale

import (
	"image"
)

// Scale2X performs the Scale2X algorithm on an input NRGBA image and returns the new scaled image
func Scale2X(img *image.NRGBA) *image.NRGBA {
	rescale_factor := 2

	total_rows := img.Bounds().Max.X
	total_columns := img.Bounds().Max.Y

	// ri stands for rescaled_image
	ri := image.NewNRGBA(image.Rect(0, 0, total_rows*rescale_factor, total_columns*rescale_factor))

	/*
	   For Scale2x algorithm (or EPX), the idea is as follows:

	   One pixel P is turned into four pixels, called 1, 2 3 and 4

	   Imagine the follow group of pixels
	             A
	           C P B
	             D

	   Pixel P becomes then
	       1 2
	       3 4

	   Whit following values
	   1=P; 2=P; 3=P; 4=P;
	   IF C==A AND C!=D AND A!=B => 1=A
	   IF A==B AND A!=C AND B!=D => 2=B
	   IF D==C AND D!=B AND C!=A => 3=C
	   IF B==D AND B!=A AND D!=C => 4=D
	*/

	for row := 0; row < total_rows; row++ {
		for column := 0; column < total_columns; column++ {

			// Get color of original image
			color := img.At(row, column)

			//1=P; 2=P; 3=P; 4=P;
			ri.Set(row*2, column*2, color)     //1
			ri.Set(row*2, column*2+1, color)   //2
			ri.Set(row*2+1, column*2, color)   //3
			ri.Set(row*2+1, column*2+1, color) //4

			// Get "letters"
			pixel_A := img.At(row-1, column)
			pixel_B := img.At(row, column+1)
			pixel_C := img.At(row, column-1)
			pixel_D := img.At(row+1, column)

			//IF C==A AND C!=D AND A!=B => 1=A
			if pixel_C == pixel_A && pixel_C != pixel_D && pixel_A != pixel_B {
				ri.Set(row*2, column*2, pixel_A)
			}

			//IF A==B AND A!=C AND B!=D => 2=B
			if pixel_A == pixel_B && pixel_A != pixel_C && pixel_B != pixel_D {
				ri.Set(row*2, column*2+1, pixel_B)
			}

			//IF D==C AND D!=B AND C!=A => 3=C
			if pixel_D == pixel_C && pixel_D != pixel_B && pixel_C != pixel_A {
				ri.Set(row*2+1, column*2, pixel_C)
			}

			//IF B==D AND B!=A AND D!=C => 4=D
			if pixel_B == pixel_D && pixel_B != pixel_A && pixel_D != pixel_C {
				ri.Set(row*2+1, column*2+1, pixel_D)
			}

		}
	}

	return ri
}
