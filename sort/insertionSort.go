package sort

func InsertionSort(data []int, a, b int) {

  for i := a + 1; i < b; i++ {
    for j := i; j > a; j-- {
      if data[j] < data[j-1] {
        // this swap method is defined in heapSort.go
        swap(data, j, j-1)
      } else {
        break
      }
    }
  }

}