package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	mapper := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	//r := []rune(s)
	result := 0
	for i := 0; i < len(s); {
		value := mapper[s[i]]
		if i+1 < len(s) {
			nextValue := mapper[(s[i+1])]
			if nextValue > value {
				result += nextValue - value
				i += 2
				continue
			}
		}

		value = mapper[s[i]]
		result += value
		i++
	}
	return result
}

//func romanValue(s string) int {
//	switch s {
//	case "I":
//		return 1
//	case "V":
//		return 5
//	case "X":
//		return 10
//	case "L":
//		return 50
//	case "C":
//		return 100
//	case "D":
//		return 500
//	case "M":
//		return 1000
//	case "IV":
//		return 4
//	case "IX":
//		return 9
//	case "XL":
//		return 40
//	case "XC":
//		return 90
//	case "CD":
//		return 400
//	case "CM":
//		return 900
//	default:
//		return 0
//	}
//}
