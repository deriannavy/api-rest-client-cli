package handler

import (
	"fmt"
	"strconv"
)

func Ternary(expresion bool, value, coditionalValue string) string {
	if expresion {
		return value
	}
	return coditionalValue
}

func Truncate(text, tail, align string, length int) string {
	var (
		alignFormat = Ternary(align == "left", "%-", "%")
		currentTail = Ternary(len(text) <= length, "", tail)
		sizedText   = text
	)

	if length < 0 {
		return ""
	}

	if len(text) > length {
		sizedText = text[:length-1] + currentTail
	}

	formater := alignFormat + strconv.Itoa(length) + "s"
	textFormated := fmt.Sprintf(formater, sizedText)

	return textFormated
}
