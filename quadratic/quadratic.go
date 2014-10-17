package main

import (
  "fmt"
  "math"
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
  solveEq([]float64{1, 4, -21})
}

func solveEq(numbers []float64) (equation quadEq) {
  equation.a = numbers[0]
  equation.b = numbers[1]
  equation.c = numbers[2]
  if disc, ok := discriminant(equation); ok {
      equation.discriminant, equation.real = disc, ok
      equation.solutions = findRealSolutions(equation)
    } else {
      fmt.Println("No real answers")
    }
    return equation
}

func findRealSolutions(equation quadEq) (solutions []float64) {
  fmt.Println(equation.a)
  fmt.Println(0 - equation.b)
  fmt.Println(equation.discriminant)
  leftRoot := ((0 - equation.b) + math.Sqrt(equation.discriminant)) / (2 * equation.a)
    solutions = append(solutions, leftRoot)
  rightRoot := ((0 - equation.b) - math.Sqrt(equation.discriminant)) / (2 * equation.a)
  solutions = append(solutions, rightRoot)
  return solutions
}

func discriminant(equation quadEq) (float64, bool) {
  if disc := (math.Pow(equation.b, 2) - (4 * equation.a * equation.c)); disc > 0 {
    return disc, true
  } else {
    return -1, false
  }
}