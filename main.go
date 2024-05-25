package main
import (
  "fmt"
  "os"
  "bufio"
  "io"
)

func main() {
  fmt.Printlin(Count(os.stdin))
}

func Count(r io.Reader) {
  // Defining a scanner which reads input from Reader io
  scanner := bufio.NewScanner(r)

  // Spliting words of scanner
  scanner.split(bufio.ScanWords)

  // Defining a word counter
  wc := 0
  
  for scanner.Scan(){
    wc++
  }
  return wc
}
