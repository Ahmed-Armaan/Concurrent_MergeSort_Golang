package main

import (
	"fmt"
	"sync"
)

func merge(arr []int, l, mid, r int) {
	tempArr1 := append([]int{}, arr[l:mid+1]...)
	tempArr2 := append([]int{}, arr[mid+1:r+1]...)

	currIndex, currIndexArr1, currIndexArr2 := l, 0, 0

	for currIndexArr1 < len(tempArr1) && currIndexArr2 < len(tempArr2) {
		if tempArr1[currIndexArr1] < tempArr2[currIndexArr2] {
			arr[currIndex] = tempArr1[currIndexArr1]
			currIndexArr1++
		} else {
			arr[currIndex] = tempArr2[currIndexArr2]
			currIndexArr2++
		}
		currIndex++
	}

	for currIndexArr1 < len(tempArr1) {
		arr[currIndex] = tempArr1[currIndexArr1]
		currIndexArr1++
		currIndex++
	}
	for currIndexArr2 < len(tempArr2) {
		arr[currIndex] = tempArr2[currIndexArr2]
		currIndexArr2++
		currIndex++
	}
}

func mergeSort(arr []int, l, r int, wg *sync.WaitGroup) {
	defer wg.Done()

	if l >= r {
		return
	}

	mid := l + (r-l)/2

	var childWg sync.WaitGroup
	childWg.Add(2)

	go mergeSort(arr, l, mid, &childWg)
	go mergeSort(arr, mid+1, r, &childWg)

	childWg.Wait()
	merge(arr, l, mid, r)
}

func main() {
	arr := []int{5, 1, 4, 2, 3, 44, 3, 665654, 8}

	var wg sync.WaitGroup
	wg.Add(1)
	go mergeSort(arr, 0, len(arr)-1, &wg)
	wg.Wait()

	fmt.Println(arr)
}
