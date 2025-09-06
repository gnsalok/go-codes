// package main

// import "testing"

// func testCalculate(t *testing.T){
// 	if calculate(2) != 4 {
// 		t.Error("Expected 2+2 is to 4")
// 	}

// }




//testing using Testify 

package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
  assert.Equal(t, calculate(2), 4)
  assert.Equal(t,calculate(5),7)
}