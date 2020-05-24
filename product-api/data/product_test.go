package data

import "testing"

func TestCheckValidation(t *testing.T)  {
  p := &Product{ Name: "test", Price: 12.0, }
  err := p.Validate()
  if err != nil {
     t.Fatal(err)
  }
}
