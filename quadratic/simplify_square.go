package main

import (
  "fmt"
  "math"
  "strconv"
)

func main() {
  fmt.Println(simplifySquareRoot(4500))
}

func simplifySquareRoot(number int) string {
  if factors, prime, perfect := primeFactorize(number); prime {
    return ("sqr" + strconv.Itoa(number))
  } else if perfect {
    return string(factors[0])
  } else {
    return simplify(factors)
  }
}

func simplify(factors []int) string {
  coefficient := 1
  radicand := 1
  fmt.Println(factors)
  tallies := make(map[int]int)
  for _, factor := range factors {
    if _, found := tallies[factor]; found {
      tallies[factor] += 1
    } else {
      tallies[factor] = 1
    }
  }
  fmt.Println(tallies)
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
  fmt.Println(coefficient)
  fmt.Println(radicand)
  return "done"
}

func primeFactorize(number int) (factors []int, prime bool, perfect bool) {
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
    factors = append(factors, squareRootFloored, squareRootFloored)
    return factors, true, false
  }
  for i := 2; i <= squareRootFloored; i++ {
    if number % i == 0 {
      factors = append(factors, i)
      otherFactors, _, _ := primeFactorize(number / i)
      for _, element := range otherFactors {
        factors = append(factors, element)
      }
      return factors, false, false
    }
  }
  factors = append(factors, number)
  return factors, true, false
}