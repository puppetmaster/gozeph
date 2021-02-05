package main

import "testing"

func TestAdditionOK(t *testing.T) {
   var param1 int
   var param2 int
   var funcReturn int

   param1 = 3
   param2 = 4

   funcReturn = monAddition(param1, param2)
   if funcReturn <= 0 {
      t.Error("monAddintion return is <= 0")
   }
}

func TestAdditionFnOK(t *testing.T) {
   var param1 int
   var param2 int
   var funcReturn int

   param1 = 5
   param2 = 5

   funcReturn = monAddition(param1, param2)
   if funcReturn != 10 {
      t.Error("Functional ERROR monAddition don't return 10")
   }
}
