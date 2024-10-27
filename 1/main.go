package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

func main() {
	var numDecimal int = 42
	var numOctal int8 = 42
	var numHex int16 = 42
	var floatVal float64 = 42
	var stringVal string = "42"
	var boolVal bool = true
	var complexVal complex64= complex(1, 2)

	fmt.Printf("Type of numDecimal: %T\n", numDecimal)
	fmt.Printf("Type of numOctal: %T\n", numOctal)
	fmt.Printf("Type of numHex: %T\n", numHex)
	fmt.Printf("Type of floatVal: %T\n", floatVal)
	fmt.Printf("Type of stringVal: %T\n", stringVal)
	fmt.Printf("Type of boolVal: %T\n", boolVal)
	fmt.Printf("Type of complexVal: %T\n", complexVal)

	// ну ваще удобнее было бы string builder было бы использвать, но для задачи пойдет
	combinedStr := strconv.Itoa(numDecimal) +
		strconv.FormatInt(int64(numOctal), 8) +
		strconv.FormatInt(int64(numHex), 16) +
		strconv.FormatFloat(floatVal, 'f', -1, 64) +
		stringVal +
		strconv.FormatBool(boolVal) +
		fmt.Sprintf("%v", complexVal)

	fmt.Println(combinedStr)

	runes := []rune(combinedStr)

	mid := len(runes) / 2
	runes = append(runes[:mid], append([]rune("go-2024"), runes[mid:]...)...)

	hash := sha256.Sum256([]byte(string(runes)))
	fmt.Printf("%x\n", hash)
}