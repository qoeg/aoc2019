package day08

type layer struct {
	data []int
	count0 int
	count1 int
	count2 int
}

type layersByPixelNum struct {
	layers []layer
	pixelNum int
}

func (l layersByPixelNum) Len() int {
	return len(l.layers)
}

func (l layersByPixelNum) Swap(i, j int) {
	l.layers[i], l.layers[j] = l.layers[j], l.layers[i]
}

func (l layersByPixelNum) Less(i, j int) bool {
	var result bool
	switch l.pixelNum {
	case 0:
		result = l.layers[i].count0 < l.layers[j].count0
	case 1:
		result = l.layers[i].count1 < l.layers[j].count1
	case 2:
		result = l.layers[i].count2 < l.layers[j].count2
	}
	return result
}
