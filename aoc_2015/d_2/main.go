package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	buf, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	str := string(buf)

	total_fabric := 0
	total_ribbon := 0

	scanner := bufio.NewScanner(strings.NewReader(str))

	for scanner.Scan() {
		num_fabric, num_ribbon := calc_required_fabric(scanner.Text())

		total_fabric += num_fabric
		total_ribbon += num_ribbon
	}

	fmt.Printf("total_fabric: %v\ntotal_ribbon: %v\n", total_fabric, total_ribbon)
}

func calc_required_fabric(line string) (int, int) {

	vals := strings.Split(line, "x")

	length, _ := strconv.Atoi(vals[0])
	width, _ := strconv.Atoi(vals[1])
	height, _ := strconv.Atoi(vals[2])

	return calc_surface_area(length, width, height) + calc_lowest_val(length, width, height), calc_ribbon(length, width, height)
}

func calc_surface_area(length int, width int, height int) int {

	surface_area := (2*(length*width) + 2*(width*height) + 2*(height*length))
	return surface_area
}

func calc_lowest_val(val1 int, val2 int, val3 int) int {

	min1_2 := math.Min(float64(val1*val2), float64(val2*val3))

	return int(math.Min(min1_2, float64(val1*val3)))
}

func calc_ribbon(length int, width int, height int) int {

	perimeter1 := math.Min(float64(calc_perimeter(length, width)), float64(calc_perimeter(length, height)))

	smallest_perimeter := int(math.Min(float64(calc_perimeter(width, height)), perimeter1))

	return smallest_perimeter + (length * width * height)
}

func calc_perimeter(val1 int, val2 int) int {

	return 2*val1 + 2*val2
}
