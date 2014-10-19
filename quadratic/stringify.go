package main

import (
  "fmt"
  "math"
  "strconv"
)

type radical struct {
  val int
  coefficient int
  radicand int
  denominator int
  multiplier int
  prime bool
  perfect bool
  factors []int
}

func (rad radical) SimplifyFraction() string {
  fmt.Println(rad)
  numerator := rad.multiplier * rad.coefficient
  finalNum := numerator
  finalDenom := rad.denominator
  var wholeNum int
  if rad.denominator == 0 {
    finalDenom = 1
    return strconv.Itoa(finalNum / finalDenom)
  } else if numerator % rad.denominator == 0 {
    return strconv.Itoa(finalNum / finalDenom)
  } else if rad.denominator % numerator == 0 {
    finalNum = 1
    finalDenom = rad.denominator / numerator
  } else {
    if numerator > rad.denominator {
      wholeNum = numerator / rad.denominator
      numerator = numerator % rad.denominator
      finalNum = numerator
      fmt.Println(numerator)
    }
    numFactors, _, _ := findPrimeFactors(numerator)
    denomFactors, denomIsPrime, _ := findPrimeFactors(rad.denominator)
    if !denomIsPrime { 
      for _, numFactor := range numFactors {
        for i, denomFactor := range denomFactors {
          if numFactor == denomFactor {
            finalNum = finalNum / numFactor
            finalDenom = finalDenom / denomFactor
            denomFactors = append(denomFactors[:i], denomFactors[i+1:]...)
            break
          }
        }
      }      
    }
  }
  if wholeNum > 0 {
    return fmt.Sprintf("<mfrac><mn>%d</mn><mn>%d</mn><mn>%d</mn></mfrac>", wholeNum, finalNum, finalDenom)
  } else {
    return fmt.Sprintf("<mfrac><mi>%d</mi><mi>%d</mi></mfrac>", finalNum, finalDenom)
  }
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
    factors = append(factors, squareRootFloored, squareRootFloored)
    return factors, true, false
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

func main() {
  var rad radical
  rad.val = 4500
  rad.coefficient = 6
  rad.multiplier = 1
  rad.denominator = 30
  fmt.Println(rad.SimplifyFraction())
}