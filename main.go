package main

import (
  "./sort"
  "os"
  "log"
)

func main() {
  // open the input file
  file, err := os.Open("input")
  if err != nil {
    log.Fatal(err)
  }
  // get the file info
  fi, err1 := file.Stat()
  if err1 != nil {
    log.Fatal(err1)
  }
  // generate original data
  data := make([][]byte, fi.Size()/100)
  for i := range data {
    tmp := make([]byte, 100)
    _, err := file.Read(tmp)
    if err != nil {
      log.Fatal(err)
    } else {
      data[i] = tmp
    }
  }

  // call HeapSort() to sort
  sort.HeapSort(data, 0, len(data))
  // write in output file
  file, err = os.Create("heapsort_output")
  if err != nil {
    log.Fatal(err)
  }
  for i := range data {
    _, err := file.Write(data[i])
    if err != nil {
      log.Fatal(err)
    }
  }

  // regenerate original data
  file, err = os.Open("input")
  if err != nil {
    log.Fatal(err)
  }
  for i := range data {
    tmp := make([]byte, 100)
    _, err := file.Read(tmp)
    if err != nil {
      log.Fatal(err)
    } else {
      data[i] = tmp
    }
  }

  // call InsertionSort() to sort
  sort.InsertionSort(data, 0, len(data))
  // write in output file
  file, err = os.Create("insertionsort_output")
  if err != nil {
    log.Fatal(err)
  }
  for i := range data {
    _, err := file.Write(data[i])
    if err != nil {
      log.Fatal(err)
    }
  }
}