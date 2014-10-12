package main
import (
  "fmt"
  "log"
  "os"
  "path/filepath"
  "strings"
)

var bigDigits = [][]string{
    {"  000  ",
     " 0   0 ",
     "0     0",
     "0     0",
     "0     0",
     " 0   0 ",
     "  000  "},
    {" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
    {" 222 ", "2   2", "   2 ", "  2  ", " 2   ", "2    ", "22222"},
    {" 333 ", "3   3", "    3", "  33 ", "    3", "3   3", " 333 "},
    {"   4  ", "  44  ", " 4 4  ", "4  4  ", "444444", "   4  ",
        "   4  "},
    {"55555", "5    ", "5    ", " 555 ", "    5", "5   5", " 555 "},
    {" 666 ", "6    ", "6    ", "6666 ", "6   6", "6   6", " 666 "},
    {"77777", "    7", "   7 ", "  7  ", " 7   ", "7    ", "7    "},
    {" 888 ", "8   8", "8   8", " 888 ", "8   8", "8   8", " 888 "},
    {" 9999", "9   9", "9   9", " 9999", "    9", "    9", "    9"},
}

func main() {
  arguments := os.Args[1:]
  bar := false
  help := false
  stringOfDigits := ""
  for i := range arguments {
    if arguments[i] == "-b" || arguments[i] == "--bar" {
      bar = true
    } else if arguments[i] == "-h" || arguments[i] == "--help" {
      help = true
    } else {
      stringOfDigits = arguments[i]
    }
  }

  if len(os.Args) == 1 || help {
    fmt.Printf("usage: %s <whole-number>\n", filepath.Base(os.Args[0]))
    os.Exit(1)
  }

  testDigit := stringOfDigits[0] - '0'
  if 0 > testDigit || testDigit > 9 {
    log.Fatal("invalid whole number")
  }

  for row := range bigDigits[0] {
    line := ""
    for column := range stringOfDigits {
      digit := stringOfDigits[column] - '0'
      line += bigDigits[digit][row] + " "
    }
    if bar && row == 0 {
      fmt.Println(strings.Repeat("*", len(line)))
    }
    fmt.Println(line)
    if bar && row == (len(bigDigits[0]) - 1) {
      fmt.Println(strings.Repeat("*", len(line)))
    }
  }

}