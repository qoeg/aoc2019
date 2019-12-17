package main

import (
	"fmt"
	"strconv"
)

type dayFlags []day

func (df *dayFlags) String() string {
	var result string
	for _, d := range []day(*df) {
		result += fmt.Sprintf("Day %d\n", d.key)
	}
	return result
}

func (df *dayFlags) Set(number string) error {
	d, err := strconv.Atoi(number)
	if err != nil {
		return err
	}

	*df = append(*df, days[d])
    return nil
}
