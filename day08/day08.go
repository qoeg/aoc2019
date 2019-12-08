package day08

import (
	"fmt"
	"sort"
	"strconv"
)

var width = 25
var height = 6

func getLayers(data []int) []layer {
	size := width*height
	num := len(data)/size

	result := make([]layer, num)

	for i := 0; i < num; i++ {
		pointer := i*size
		l := layer{
			data: data[pointer:pointer+size],
		}
		for _, pix := range l.data {
			switch pix {
			case 0:
				l.count0++
			case 1:
				l.count1++
			case 2:
				l.count2++
			}
		}
		result[i] = l
	}

	return result
}

func render(layers []layer) {
	print := func(img []string) {
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				fmt.Print(img[(y*width)+x])
			}
			fmt.Println()
		}
	}
	
	image := make([]string, width*height)
	for i := len(layers)-1; i >= 0; i-- {
		for j, pixel := range layers[i].data {
			switch pixel {
			case 0:
				image[j] = "█"
			case 1:
				image[j] = "|"
			case 2:
				if image[j] == "" {
					image[j] = " "
				}
			}
		}
	}

	print(image)
}

// Answer1 returns the first puzzle answer
func Answer1() string {
	layers := getLayers(Input)
	sort.Sort(layersByPixelNum{layers, 0})

	return strconv.Itoa(layers[0].count1*layers[0].count2)
}

// Answer2 returns the second puzzle answer
func Answer2() string {
	layers := getLayers(Input)
	render(layers)

	return "N/A (image)"
}
