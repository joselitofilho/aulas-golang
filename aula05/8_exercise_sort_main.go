package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func MultiMergeSort(data []float64, res chan []float64) {
	if len(data) < 2 {
		res <- data
		return
	}

	leftChan := make(chan []float64)
	rightChan := make(chan []float64)
	middle := len(data) / 2

	go MultiMergeSort(data[:middle], leftChan)
	go MultiMergeSort(data[middle:], rightChan)
	ldata := <-leftChan
	rdata := <-rightChan

	close(leftChan)
	close(rightChan)
	res <- Merge(ldata, rdata)
	return
}

func RunMultiMergeSort(data []float64) (multiResult []float64) {
	res := make(chan []float64)
	go MultiMergeSort(data, res)
	multiResult = <-res
	return
}

func Merge(ldata []float64, rdata []float64) (result []float64) {
	result = make([]float64, len(ldata)+len(rdata))
	lidx, ridx := 0, 0

	for i := 0; i < cap(result); i++ {
		switch {
		case lidx >= len(ldata):
			result[i] = rdata[ridx]
			ridx++
		case ridx >= len(rdata):
			result[i] = ldata[lidx]
			lidx++
		case ldata[lidx] < rdata[ridx]:
			result[i] = ldata[lidx]
			lidx++
		default:
			result[i] = rdata[ridx]
			ridx++
		}
	}
	return
}

func generateArray(numOfElements int) []float64 {
	s := make([]float64, numOfElements)
	for i := 0; i < cap(s); i++ {
		s[i] = rand.Float64() * float64(numOfElements)
	}
	return s
}

func main() {
	var size int = 3

	s := generateArray(size)
	fmt.Println("generate "+strconv.Itoa(size)+" numbers...", s)

	fmt.Println("running multithread without limited number of threads")
	start := time.Now()
	multiResult := RunMultiMergeSort(s)
	fmt.Println(time.Since(start))

	fmt.Println("Verifying the answer")
	fmt.Println(multiResult)
}
