package innValidator

import (
  "fmt"
  "strconv"
)

var (
// innLegalPersonRegexp   = regexp.MustCompile("^[0-9]{10}$")
// innPrivatePersonRegexp = regexp.MustCompile("^[0-9]{12}$")
)

// func IsLegalPersonInnValid(inn string) (isValid bool, err error) {
//   if !innLegalPersonRegexp.MatchString(inn) {
//     return false, fmt.Errorf("%s is not legal person inn", inn)
//   }
//
//   return
// }

// func legalPerson(inn string) (valid bool, err error) {
//   weights := []int{2, 4, 10, 3, 5, 9, 4, 6, 8}
//   var r = 0
//   for index, value := range weights {
//     i, _ := strconv.Atoi(string(inn[index]))
//     r += i * value
//
//   }
//   i, _ := strconv.Atoi(string(inn[9]))
//   valid = r%11%10 == i
//   return
// }

// convert string to slice int
// intStringToIntArray("42") // => {4, 2}
func intStringToIntArray(inputString string) (array []int, err error) {
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
