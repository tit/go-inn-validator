package innValidator

import (
  "fmt"
  "regexp"
  "strconv"
)

func IsLegalPersonInnValid(inn string) (isValid bool, err error) {
  weights := [9]int{2, 4, 10, 3, 5, 9, 4, 6, 8}

  if !regexp.MustCompile("^[0-9]{10}$").MatchString(inn) {
    return false, fmt.Errorf("%s is not legal person inn", inn)
  }

  array, err := intStringToIntArray(inn)
  if err != nil {
    return false, fmt.Errorf("%s is not legal person inn", inn)
  }

  var result int
  for index, weight := range weights {
    result += array[index] * weight
  }

  if result%11%10 == array[9] {
    return true, err
  } else {
    return false, fmt.Errorf("%s is not legal person inn", inn)
  }
}

func IsPrivatePersonInnValid(inn string) (isValid bool, err error) {
  var weights1 = [10]int{7, 2, 4, 10, 3, 5, 9, 4, 6, 8}
  var weights2 = [11]int{3, 7, 2, 4, 10, 3, 5, 9, 4, 6, 8}

  if !regexp.MustCompile("^[0-9]{12}$").MatchString(inn) {
    return false, fmt.Errorf("%s is not legal person inn", inn)
  }

  array, err := intStringToIntArray(inn)
  if err != nil {
    return false, fmt.Errorf("%s is not legal person inn", inn)
  }

  var result int
  for index, weight := range weights1 {
    result += array[index] * weight
  }

  if result%11%10 != array[10] {
    return false, fmt.Errorf("%s is not legal person inn", inn)
  }

  result = 0
  for index, weight := range weights2 {
    result += array[index] * weight
  }

  if result%11%10 != array[11] {
    return false, fmt.Errorf("%s is not legal person inn", inn)
  }

  return true, err
}

func intStringToIntArray(inputString string) (array []int, err error) {
  if inputString == "" {
    return []int{}, fmt.Errorf("%s is empty", inputString)
  }

  for _, value := range inputString {
    char := string(value)

    // strconv.Atoi return integer and nil error
    // if all chars in string is integer
    digit, err := strconv.Atoi(char)
    if err != nil {
      return []int{}, fmt.Errorf("%s in %s is not int", char, inputString)
    }

    array = append(array, digit)
  }

  return
}
