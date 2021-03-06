package innValidator

import (
  "reflect"
  "testing"
)

func Test_intStringToIntArray(t *testing.T) {
  type args struct {
    inputString string
  }
  tests := []struct {
    name      string
    args      args
    wantArray []int
    wantErr   bool
  }{
    {name: "one digit", args: args{inputString: "0"}, wantArray: []int{0}, wantErr: false},
    {name: "lead zero", args: args{inputString: "01"}, wantArray: []int{0, 1}, wantErr: false},
    {name: "lead minus", args: args{inputString: "-1"}, wantArray: []int{}, wantErr: true},
    {name: "float", args: args{inputString: "42.2"}, wantArray: []int{}, wantErr: true},
    {name: "empty string", args: args{inputString: ""}, wantArray: []int{}, wantErr: true},
    {name: "chars only", args: args{inputString: "foo"}, wantArray: []int{}, wantErr: true},
    {name: "mix chars digits", args: args{inputString: "foo42"}, wantArray: []int{}, wantErr: true},
    {name: "mix lead digits chars", args: args{inputString: "42foo"}, wantArray: []int{}, wantErr: true},
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      gotArray, err := intStringToIntArray(tt.args.inputString)
      if (err != nil) != tt.wantErr {
        t.Errorf("intStringToIntArray() error = %v, wantErr %v", err, tt.wantErr)
        return
      }
      if !reflect.DeepEqual(gotArray, tt.wantArray) {
        t.Errorf("intStringToIntArray() = %v, want %v", gotArray, tt.wantArray)
      }
    })
  }
}

func TestIsLegalPersonInnValid(t *testing.T) {
  type args struct {
    inn string
  }
  tests := []struct {
    name        string
    args        args
    wantIsValid bool
    wantErr     bool
  }{
    {name: "valid inn", args: args{"7802565953"}, wantIsValid: true, wantErr: false},
    {name: "invalid inn", args: args{"1234567890"}, wantIsValid: false, wantErr: true},
    {name: "less count digits", args: args{"780256595"}, wantIsValid: false, wantErr: true},
    {name: "more count digits", args: args{"78025659530"}, wantIsValid: false, wantErr: true},
    {name: "text", args: args{"foo"}, wantIsValid: false, wantErr: true},
    {name: "mix lead valid inn text", args: args{"7802565953foo"}, wantIsValid: false, wantErr: true},
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      gotIsValid, err := IsLegalPersonInnValid(tt.args.inn)
      if (err != nil) != tt.wantErr {
        t.Errorf("IsLegalPersonInnValid() error = %v, wantErr %v", err, tt.wantErr)
        return
      }
      if gotIsValid != tt.wantIsValid {
        t.Errorf("IsLegalPersonInnValid() = %v, want %v", gotIsValid, tt.wantIsValid)
      }
    })
  }
}

func TestIsPrivatePersonInnValid(t *testing.T) {
  type args struct {
    inn string
  }
  tests := []struct {
    name        string
    args        args
    wantIsValid bool
    wantErr     bool
  }{
    {name: "valid inn", args: args{"614309291796"}, wantIsValid: true, wantErr: false},
    {name: "invalid inn", args: args{"614309291790"}, wantIsValid: false, wantErr: true},
    {name: "less count digits", args: args{"61430929179"}, wantIsValid: false, wantErr: true},
    {name: "more count digits", args: args{"6143092917960"}, wantIsValid: false, wantErr: true},
    {name: "text", args: args{"foo"}, wantIsValid: false, wantErr: true},
    {name: "mix lead valid inn text", args: args{"614309291796foo"}, wantIsValid: false, wantErr: true},
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      gotIsValid, err := IsPrivatePersonInnValid(tt.args.inn)
      if (err != nil) != tt.wantErr {
        t.Errorf("IsPrivatePersonInnValid() error = %v, wantErr %v", err, tt.wantErr)
        return
      }
      if gotIsValid != tt.wantIsValid {
        t.Errorf("IsPrivatePersonInnValid() = %v, want %v", gotIsValid, tt.wantIsValid)
      }
    })
  }
}
