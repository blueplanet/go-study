package main

import "golang.org/x/tour/pic"

func Pic(dx int, dy int) [][]uint8 {
	var pic [][]uint8

	for y := 0; y <  dy; y++ {
		var picRow []uint8

		for x := 0; x < dx; x++ {
			picRow = append(picRow, uint8(x ^ y))
		}

		pic = append(pic, picRow)
	}

	return pic
}

func main() {
	pic.Show(Pic)
}
