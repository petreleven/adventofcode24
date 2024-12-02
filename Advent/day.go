package advent

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func checkErr(err error) {
	panic(err)
}

func Entry() {
	loadFileSortLocations()
}

func mergeSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}
	var pivotIndex uint = uint(len(list) / 2)
	pivot := list[pivotIndex]
	lower := []int{}
	higher := []int{}
	sorted_list := []int{}
	for i := 0; i < len(list); i++ {
		if i == int(pivotIndex) {
			continue
		}
		if list[i] > pivot {
			higher = append(higher, list[i])
		} else {
			lower = append(lower, list[i])
		}
	}
	sorted_list = mergeSort(lower)
	sorted_list = append(sorted_list, pivot)
	sorted_list = append(sorted_list, mergeSort(higher)...)
	return sorted_list
}

func loadFileSortLocations() {
	bytes, err := os.ReadFile("./Advent/loc.txt")
	if err != nil {
		checkErr(err)
	}
	string_data := string(bytes)
	//free bytes buffer
	bytes = nil
	string_buffer := strings.Split(string_data, "\n")
	list_1, list_2 := []int{}, []int{}
	for _, line := range string_buffer {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		pair := strings.Fields(line)
		first, _ := strconv.Atoi(pair[0])
		second, _ := strconv.Atoi(pair[1])
		list_1 = append(list_1, first)
		list_2 = append(list_2, second)
	}
	//free strings buffer
	string_buffer = nil
	list_1 = mergeSort(list_1)
	list_2 = mergeSort(list_2)
	findTotalDistance(&list_1, &list_2)
}

func findTotalDistance(list1, list2 *[]int) {
	totalDistance := 0
	for i := 0; i < len(*list1); i++ {
		totalDistance += int(math.Abs(float64((*list1)[i] - (*list2)[i])))
	}
	fmt.Println(totalDistance)
}
