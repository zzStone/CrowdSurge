package main

import (
  "fmt"
  "./go_sort"
  "os"
  "log"
)

func main() {
  // open the input file
  file, err := os.Open("pennyinput")
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
  count := 1
  for i := range data {
    tmp := make([]byte, 100)
    var err error
    count, err = file.Read(tmp)
    if err != nil {
      log.Fatal(err)
    } else {
      data[i] = tmp
    }
  }

  // call HeapSort() to sort
  go_sort.HeapSort(data, 0, len(data))

  fmt.Printf("read %d bytes: %q\n", count, data[0])
}