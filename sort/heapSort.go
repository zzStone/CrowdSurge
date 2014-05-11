package sort

// maintain max-heap property from index
func siftDown(data []int, lo, hi, first int) {
  root := lo
  for {
    child := 2*root + 1
    if child >= hi {
      break
    }
    if child+1 < hi && data[first+child] < data[first+child+1] {
      child++
    }
    if data[first+root] >= data[first+child] {
      return
    }
    swap(data, first+root, first+child)
    root = child
  }
}

// swap the index a and b of array
func swap(data []int, a, b int) {
  temp := data[a]
  data[a] =  data[b]
  data[b] = temp
}

func HeapSort(data []int, a, b int) {
  first := a
  lo := 0
  hi := b - a
  // build max-heap from index a to index b
  for i := (hi - 1) / 2; i >= 0; i-- {
    siftDown(data, i, hi, first)
  }
  // pop maximum to the end and rebuild the former elements
  for i := hi - 1; i >= 0; i-- {
    swap(data, first, first+i)
    siftDown(data, lo, i, first)
  }
}