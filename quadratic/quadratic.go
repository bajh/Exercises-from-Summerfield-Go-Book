package main

import (
  "fmt"
  "net/http"
  "math"
  "log"
  "strconv"
)

const (
  pageTop = `<!DOCTYPE HTML><html><head><style>.error{color:#FF0000;}</style><script type="text/javascript" src="http://cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML"></script></head><title>Quadratic Equation Solver</title>
    <body><h3>Please Enter a Quadratic Equation</h3>`
  form = `<form action="/" method="POST">
    <input type="text" name="a">x<sup>2</sup> + <input type="text" name="b">x + <input type="text" name="c">
    <input type="submit" value="Solve">
    </form>`
    pageBottom = `</body></html>`
    anError = `<p class="error">%s</p>`
)

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
        equation, leftSolution, rightSolution := newEquation(numbers)
        fmt.Fprint(writer, formatSolutions(equation, leftSolution, rightSolution))
      } else if message != "" {
        fmt.Fprintf(writer, anError, message)
      }
    }
  }
  fmt.Fprint(writer, pageBottom)
}

func processRequest(request *http.Request) ([]int, string, bool) {
  var numbers []int
  for _, coefficient := range []string{"a", "b", "c"} {
    if formVal, found := request.Form[coefficient]; found {
      if formVal[0] == "" {
        formVal[0] = "1"
      }
      if x, err := strconv.ParseInt(formVal[0], 10, 8); err != nil {
        return numbers, "One or more fields are invalid", false
      } else {
        numbers = append(numbers, int(x))
      }
    }
  }
  return numbers, "", true
}

func formatSolutions(equation quadEq, leftSolution string, rightSolution string) string {
  if leftSolution == rightSolution {
    return "<p>" + strconv.Itoa(equation.a) + "x<sup>2</sup>" + strconv.Itoa(equation.b) + " x " + strconv.Itoa(equation.c) + " -> x = <math mode='display'>" + leftSolution + "</math></p>"
  } else {
    return "<p>" + strconv.Itoa(equation.a) + "x<sup>2</sup>" + strconv.Itoa(equation.b) + " x " + strconv.Itoa(equation.c) + " -> x = <math mode='display'>" + leftSolution + "</math> and x = <math mode='display'>" + rightSolution + "</math></p>"
  }
}


/*******************************/

type quadEq struct {
  a int
  b int
  c int
  discriminant int
  complex bool
  perfectDisc bool
  primeDisc bool
}

type fraction struct {
  wholeNum int
  num int
  denom int
}

type radical struct {
  coefficient int
  radicand int
}

func newEquation(numbers []int) (equation quadEq, leftSolution string, rightSolution string) {
  equation.a = numbers[0]
  equation.b = numbers[1]
  equation.c = numbers[2]
  leftSolution, rightSolution = equation.solve()
  return equation, leftSolution, rightSolution
}

func (eq *quadEq) solve() (leftResult string, rightResult string) {
  eq.discriminant = (int(math.Pow(float64(eq.b), 2)) - (4 * eq.a * eq.c))
  if eq.discriminant < 0 {
    eq.discriminant = 0 - eq.discriminant
    eq.complex = true
  }
  squareRootFloored := int(math.Sqrt(float64(eq.discriminant)))
  //Discriminant is a pefect square...
  denominator := 2 * eq.a
  if float64(eq.discriminant) / float64(squareRootFloored) == float64(squareRootFloored) {
    eq.perfectDisc = true
    leftNumerator := -eq.b - squareRootFloored
    rightNumerator := -eq.b + squareRootFloored
    if leftNumerator % denominator == 0 {
      leftResult = "<mi>" + strconv.Itoa(leftNumerator / denominator) + "</mi>"
    } else {
      leftResult = simplifyFraction(leftNumerator, denominator).String()
    }
    if rightNumerator % denominator == 0 {
      rightResult = "<mi>" + strconv.Itoa(rightNumerator / denominator) + "</mi>"
    } else {
      rightResult = simplifyFraction(rightNumerator, denominator).String()
    }
  } else {
    radical := simplifyRadical(eq.discriminant)
    if eq.b % denominator == 0 {
      leftResult += "<mi>" + strconv.Itoa(-eq.b / denominator) +
                    "</mi><mo>+</mo><mfrac><mrow><mn>" + strconv.Itoa(radical.coefficient) +
                    "</mn><msqrt><mn>" + strconv.Itoa(radical.radicand) + 
                    "</mn></msqrt></mrow><mrow><mn>" + strconv.Itoa(denominator) +
                    "</mn></mrow></mfrac>"
      rightResult += "<mi>" + strconv.Itoa(-eq.b / denominator) +
                    "</mi><mo>-</mo><mfrac><mrow><mn>" + strconv.Itoa(radical.coefficient) +
                    "</mn><msqrt><mn>" + strconv.Itoa(radical.radicand) + 
                    "</mn></msqrt></mrow><mrow><mn>" + strconv.Itoa(denominator) +
                    "</mn></mrow></mfrac>"
    } else {
      leftResult += "<mfrac><mrow><mi>" + strconv.Itoa(-eq.b) +
                    "</mi><mo>+</mo><mn>" + strconv.Itoa(radical.coefficient) +
                    "</mn><msqrt><mn>" + strconv.Itoa(radical.radicand) +
                    "</mn></msqrt></mrow><mrow><mn>" + strconv.Itoa(denominator) +
                    "</mn></mrow></mfrac>"
      rightResult += "<mfrac><mrow><mi>" + strconv.Itoa(-eq.b) +
                    "</mi><mo>-</mo><mn>" + strconv.Itoa(radical.coefficient) +
                    "</mn><msqrt><mn>" + strconv.Itoa(radical.radicand) +
                    "</mn></msqrt></mrow><mrow><mn>" + strconv.Itoa(denominator) +
                    "</mn></mrow></mfrac>"
    }
  }
  return leftResult, rightResult
}

func simplifyFraction(numerator int, denominator int) fraction {
  frac := fraction{num: numerator, denom: denominator}
  negativeNum := false
  if numerator < 0 {
    negativeNum = true
    numerator = int(math.Abs(float64(numerator)))
  }
  if numerator > denominator {
    frac.wholeNum = numerator / denominator
    numerator = numerator - (frac.wholeNum * denominator)
  }
  numeratorFactors, _, _ := findPrimeFactors(numerator)
  denominatorFactors, _, _ := findPrimeFactors(denominator)
  for _, numeratorFactor := range numeratorFactors {
    for i, denominatorFactor := range denominatorFactors {
      if numeratorFactor == denominatorFactor {
        numerator = numerator / numeratorFactor
        denominator = denominator / denominatorFactor
        denominatorFactors = append(denominatorFactors[:i], denominatorFactors[i+1:]...)
        break
      }
    }
  }
  if negativeNum {
    frac.num = -numerator
  } else {
    frac.num = numerator
  }
  frac.denom = denominator
  return frac
}

func findPrimeFactors(number int) (factors []int, prime bool, perfect bool) {
  if number == 1 || number == 2 || number == 3 {
    factors = append(factors, number)
    if number == 1 {
      return factors, false, true
    } else {
      return factors, true, false
    }
  }
  squareRootFloored := int(math.Sqrt(float64(number)))
  if float64(number) / float64(squareRootFloored) == float64(squareRootFloored) {
    //Number is a perfect square
    factors = append(factors, squareRootFloored, squareRootFloored)
    return factors, false, true
  }
  for i := 2; i <= squareRootFloored; i++ {
    if number % i == 0 {
      factors = append(factors, i)
      otherFactors, _, _ := findPrimeFactors(number / i)
      for _, element := range otherFactors {
        factors = append(factors, element)
      }
      return factors, false, false
    }
  }
  factors = append(factors, number)
  return factors, true, false
}

func simplifyRadical(number int) (rad radical) {
  coefficient := 1
  radicand := 1
  var factors []int
  factors, _, _ = findPrimeFactors(number)
  tallies := mapOfFactors(factors)
  for factor, tally := range tallies {
    if tally >= 2 {
      coefficient *= factor * (tally / 2)
      if tally % 2 != 0 {
        radicand *= factor * (tally % 2)
      }
    } else {
      radicand *= factor
    }
  }
  rad.coefficient = coefficient
  rad.radicand = radicand
  return rad
}

func mapOfFactors(factors []int) map[int]int {
  tallies := make(map[int]int)
  for _, factor := range factors {
    if _, found := tallies[factor]; found {
      tallies[factor] += 1
    } else {
      tallies[factor] = 1
    }
  }
  return tallies
}

func (frac fraction) String() (result string) {
  if frac.wholeNum !=0 && frac.num < 0 {
    frac.wholeNum = -frac.wholeNum
    frac.num = -frac.num
  }
  if frac.num == 0 {
    result += "<mi>" + strconv.Itoa(frac.wholeNum) + "</mi>"
  } else {
    if frac.wholeNum != 0 {
      result += "<mi>" + strconv.Itoa(frac.wholeNum) + "</mi>"
    }
    result += "<mfrac><mrow><mn>" + strconv.Itoa(frac.num) + "</mn></mrow>"
    result += "<mrow><mn>" + strconv.Itoa(frac.denom) + "</mn>"
    result += "</mrow></mfrac>"
  }
  return result
}