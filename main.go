package main

import (
  "bufio"
  "flag"
  // "io/ioutil"
  "os"
  // "strconv"
)

func main() {
  flag.Parse()

  scanner := bufio.NewScanner(os.Stdin)

  for scanner.Scan() {
    os.Stdout.WriteString(scanner.Text() + "\n")
  }
}
