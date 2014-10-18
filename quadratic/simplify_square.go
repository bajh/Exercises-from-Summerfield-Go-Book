package main

import (
  "fmt"
  "math"
)

func main() {
  fmt.Println(primeFactorize(44))
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
  if number / squareRootFloored == number {
    factors = append(factors, squareRootFloored, squareRootFloored)
    return factors, false, true
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