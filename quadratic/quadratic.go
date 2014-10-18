package main

import (
  "fmt"
  "net/http"
  "math"
  "log"
  "strconv"
)

const (
  pageTop = `<!DOCTYPE HTML><html><head><style>.error{color:#FF0000;}</style></head><title>Quadratic Equation Solver</title>
    <body><h3>Please Enter a Quadratic Equation</h3>`
  form = `<form action="/" method="POST">
    <input type="text" name="a">x<sup>2</sup> + <input type="text" name="b">x + <input type="text" name="c">
    <input type="submit" value="Solve">
    </form>`
    pageBottom = `</body></html>`
    anError = `<p class="error">%s</p>`
)

type quadEq struct {
  a float64
  b float64
  c float64
  discriminant float64
  real bool
  solutions []float64
}

func main() {
  http.HandleFunc("/", homePage)
  if err := http.ListenAndServe(":9001", nil); err != nil {
    log.Fatal("failed to start server", err)
  }
}

func homePage(writer http.ResponseWriter, request *http.Request) {
  fmt.Fprint(writer, pageTop, form)
  if request.Method == "POST" {
    err := request.ParseForm()
    if err != nil {
      fmt.Fprintf(writer, anError, err)
    } else {
      if numbers, message, ok := processRequest(request); ok {
        equation := solveEq(numbers)
        fmt.Println(equation)
        fmt.Fprint(writer, formatSolution(equation))
      } else if message != "" {
        fmt.Fprintf(writer, anError, message)
      }
    }
  }
  fmt.Fprint(writer, pageBottom)
}

func processRequest(request *http.Request) ([]float64, string, bool) {
  var numbers []float64
  for _, coefficient := range []string{"a", "b", "c"} {
    if formVal, found := request.Form[coefficient]; found {
      if formVal[0] == "" {
        formVal[0] = "1"
      }
      if x, err := strconv.ParseFloat(formVal[0], 64); err != nil {
        return numbers, "One or more fields are invalid", false
      } else {
        numbers = append(numbers, x)
      }
    }
  }
  return numbers, "", true
}

func formatSolution(equation quadEq) string {
  return fmt.Sprintf(`<p>%fx<sup>2</sup> + %fx + %f -> x = %f or x = %f</p>`, equation.a, equation.b, equation.c, equation.solutions[0], equation.solutions[1])
}

func solveEq(numbers []float64) (equation quadEq) {
  equation.a = numbers[0]
  equation.b = numbers[1]
  equation.c = numbers[2]
  if disc, ok := discriminant(equation); ok {
      equation.discriminant, equation.real = disc, ok
      equation.solutions = findRealSolutions(equation)
    } else {
      equation.solutions =findComplexSolutions(equation)
    }
    return equation
}

func findRealSolutions(equation quadEq) (solutions []float64) {
  leftRoot := ((0 - equation.b) + math.Sqrt(equation.discriminant)) / (2 * equation.a)
    solutions = append(solutions, leftRoot)
  rightRoot := ((0 - equation.b) - math.Sqrt(equation.discriminant)) / (2 * equation.a)
  solutions = append(solutions, rightRoot)
  return solutions
}

func findComplexSolutions(equation quadEq) (solutions []float64]) {
  
}

func discriminant(equation quadEq) (float64, bool) {
  if disc := (math.Pow(equation.b, 2) - (4 * equation.a * equation.c)); disc > 0 {
    return disc, true
  } else {
    return -1, false
  }
}