package main

import (
  "fmt"
  "./sort"
)

func main() {
  array := []int{4,9,5,7,1,8,2,6,3}
  fmt.Println("The origin array is: ", array)
  sort.HeapSort(array, 0, len(array))
  fmt.Println("The sorted array is: ", array)
}