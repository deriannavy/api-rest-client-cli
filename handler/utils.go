package handler

import (
	"fmt"
	"strconv"
	"strings"
)

func Ternary(expresion bool, value, coditionalValue string) string {
	if expresion {
		return value
	}
	return coditionalValue
}

func TernaryNumber(expresion bool, value, coditionalValue int) int {
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

func FillCenter(text string, length int) string {
	if len(text) >= length {
		return text
	}

	spaces := length - len(text)
	borderSpace := strings.Repeat(" ", spaces/2)

	return text + strconv.Itoa(length) + borderSpace
}
