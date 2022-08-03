package main

import (
	"reflect"
	"testing"
)

func TestStringToASCIIArray(t *testing.T) {
  var tests = []struct {
    input string
    expected []int32
  }{
    {string(" "), []int32{32}},
    {string("A"), []int32{65}},
    {string("a"), []int32{97}},
  } 

  for _, test := range tests {
    if output := StringToASCIIArray(test.input); !reflect.DeepEqual(output, test.expected) {
      t.Error("Test failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
    }
  }
}

func TestIntToBinary(t *testing.T) {
  var tests = []struct {
    input []rune
    expected string
  }{
    { []rune{0}, string("00000000")},
    { []rune{72, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 33}, 
    string("010010000110010101101100011011000110111100100000011101110110111101110010011011000110010000100001")},
  }

  for _, test := range tests {
    if output := IntToBinary(test.input); output != test.expected {
      t.Error("Test failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
    }
  }
}

func TestSeparateInSix(t *testing.T) {
  var tests = []struct {
    input string
    expectedStr []string
    expectedPadding int8
  }{
    {string("01100001"), []string{"011000", "010000"}, 2},
  }

  for _, test := range tests {
    if outputStr, outputPadding := SeparateInSix(test.input); !reflect.DeepEqual(outputStr, test.expectedStr) || outputPadding != test.expectedPadding {
      t.Error("Test failed: {} inputted, {}{} expected, received: {}{}", test.input, test.expectedStr, test.expectedPadding, outputStr, outputPadding)
    }
  }
 }
