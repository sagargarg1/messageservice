package utils

import (
	"encoding/json"
	"io"
)

// ToJSON serializes the given interface into a string based JSON format
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(i)
}

//FromJSON deserializes the object from JSON string
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}

func Reverse(text string) (result string) {
   for _, v := range text {
      result = string(v) + result
   }
   return
}

func IsPalindrome(text string) bool {
   if text == Reverse(text) {
      return true
   }
   return false
}
