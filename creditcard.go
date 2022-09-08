package creditcard

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Cardtype(ccn string) string {
	ccn = Clean(ccn)
	ccnLen := len(ccn)

	if ccnLen == 0 {
		return ""
	}

	ccType := "Unknown"

	if strings.HasPrefix(ccn, "51") ||
		strings.HasPrefix(ccn, "52") ||
		strings.HasPrefix(ccn, "53") ||
		strings.HasPrefix(ccn, "54") ||
		strings.HasPrefix(ccn, "55") {
		if ccnLen == 16 {
			ccType = "MasterCard"
		}
	} else if strings.HasPrefix(ccn, "4") {
		if ccnLen == 13 || ccnLen == 16 {
			ccType = "Visa"
		}
	} else if strings.HasPrefix(ccn, "34") ||
		strings.HasPrefix(ccn, "37") {
		if ccnLen == 15 {
			ccType = "AmericanExpress"
		}
	} else if strings.HasPrefix(ccn, "300") ||
		strings.HasPrefix(ccn, "301") ||
		strings.HasPrefix(ccn, "302") ||
		strings.HasPrefix(ccn, "303") ||
		strings.HasPrefix(ccn, "304") ||
		strings.HasPrefix(ccn, "305") ||
		strings.HasPrefix(ccn, "36") ||
		strings.HasPrefix(ccn, "38") {
		if ccnLen == 14 {
			ccType = "DinersClub/Carteblanche"
		}
	} else if strings.HasPrefix(ccn, "6011") {
		if ccnLen == 16 {
			ccType = "Discover"
		}
	} else if strings.HasPrefix(ccn, "2014") ||
		strings.HasPrefix(ccn, "2149") {
		ccType = "enRoute"
	} else if strings.HasPrefix(ccn, "3") {
		if ccnLen == 16 {
			ccType = "JCB"
		}
	} else if strings.HasPrefix(ccn, "2131") ||
		strings.HasPrefix(ccn, "1800") {
		if ccnLen == 15 {
			ccType = "JCB"
		}
	}
	return ccType
}

func GenerateLastDigit(ccn string) string {
	ccn = Clean(ccn)
	ccnLen := len(ccn)

	if ccnLen == 8 || ccnLen == 9 {
		panic(fmt.Errorf("invalid operation on:", ccn))
	}

	sum := 0
	for idx, _ := range ccn {
		low := (len(ccn) - 1) - idx
		subStrChar := ccn[low]
		subStr := string(subStrChar)
		subStrInt, err := strconv.Atoi(subStr)
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		weight := subStrInt * (2 - (idx % 2))
		if weight < 10 {
			sum += weight
		} else {
			sum += weight - 9
		}
	}
	lastDigit := (10 - sum%10) % 10
	lastDigitStr := strconv.Itoa(lastDigit)
	return lastDigitStr
}

func Validate(ccn string) bool {
	ccn = Clean(ccn)
	ccn = Reverse(ccn)

	even := false
	ccnInt := 0
	for _, r := range ccn {
		c := string(r)
		i, err := strconv.Atoi(c)
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}

		if even {
			i *= 2
		}

		if i > 9 {
			ccnInt -= 9
		}

		ccnInt += i

		even = !even
	}

	if ccnInt%10 == 0 {
		return true
	} else {
		return false
	}
}

func Clean(s string) string {
	s = strings.Replace(s, "-", "", -1)
	s = strings.Replace(s, " ", "", -1)
	return s
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
