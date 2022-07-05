package helpers

import "strings"

type String struct {
}

func NewString() *String {
	return &String{}
}

func (s String) Compare(str1 string, str2 string) bool {
	if result := strings.Compare(str1, str2); result < 0 {

	}
	return false
}
