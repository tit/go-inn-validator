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
    {name: "one digit", args: args{inputString: "0",}, wantArray: []int{0}, wantErr: false,},
    {name: "lead zero", args: args{inputString: "01",}, wantArray: []int{0, 1}, wantErr: false,},
    {name: "lead minus", args: args{inputString: "-1",}, wantArray: []int{}, wantErr: true,},
    {name: "float", args: args{inputString: "42.2",}, wantArray: []int{}, wantErr: true,},
    {name: "empty string", args: args{inputString: "",}, wantArray: []int{}, wantErr: true,},
    {name: "chars only", args: args{inputString: "foo",}, wantArray: []int{}, wantErr: true,},
    {name: "mix chars digits", args: args{inputString: "foo42",}, wantArray: []int{}, wantErr: true,},
    {name: "mix lead digits chars", args: args{inputString: "42foo",}, wantArray: []int{}, wantErr: true,},
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
