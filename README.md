[![GoDoc](https://godoc.org/github.com/ttacon/libphonenumber?status.png)](https://godoc.org/github.com/tit/go-inn-validator)

# innValidator
Валидатор ИНН физических и юридических лиц

## Установка

```bash
$ go get github.com/tit/go-inn-validator
```

## Использование

```go
package main

import (
  "fmt"
  "github.com/tit/go-inn-validator"
)

func main() {
  legalPersonInn := "7802565953"
  isValid, err := innValidator.IsLegalPersonInnValid(legalPersonInn)
  if err != nil {
    return
  }
  fmt.Printf("%s is %t legal person inn\n", legalPersonInn, isValid)

  privatePersonInn := "614309291796"
  isValid, err = innValidator.IsPrivatePersonInnValid(privatePersonInn)
  if err != nil {
    return
  }
  fmt.Printf("%s is %t private person inn\n", privatePersonInn, isValid)
}
```