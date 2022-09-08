// Copyright 2013 Doug Sparling. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"github.com/birchb1024/drakaea"
	"math/rand"
	"os"
	"unicode"
)

func init() {
	cd := []struct {
		in  string
		out int
	}{
		{in: "", out: 0},
		{in: "1111-22", out: 6},
		{in: "1111-2222", out: 8},
		{in: "111122", out: 6},
		{in: "11112222", out: 8},
		{in: "1111-22", out: 6},
		{in: "1111-2222", out: 8},
		{in: "1111-2222-3333-444", out: 15},
		{in: "1111-2222-3333-4444", out: 16},
		{in: "1111-2222-3333-44445", out: 17},
		{in: "12345678901234567890", out: 20},
	}
	for _, test := range cd {
		if countDigits(test.in) != test.out {
			panic(fmt.Errorf("unit test failed %v", test))
		}
	}
}

func init() {
	cd := []struct {
		in  string
		out string
	}{
		{in: "", out: ""},
		{in: "111@*$I*@#$(&)!(@&11-22", out: "1111122"},
		{in: "1111-2222-3333-444", out: "111122223333444"},
		{in: "1111 2222 3333 444", out: "111122223333444"},
		{in: "12345678901234567890", out: "12345678901234567890"},
	}
	for _, test := range cd {
		if stripLine(test.in) != test.out {
			panic(fmt.Errorf("unit test failed, wanted %v got %v", test.out, stripLine(test.in)))
		}
	}
}

func init() {
	if prettyCard("1111222233334444") != "1111 2222 3333 4444" {
		panic(fmt.Errorf("unit test failed, prettyCard(\"1111222233334444\") != \"1111 2222 3333 4444\""))
	}
}

func main() {
	if len(os.Args) > 1 {
		panic(fmt.Errorf("Too command line arguments: %v", os.Args))
	}

	rand.Seed(42)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		err := scanner.Err()
		if err != nil {
			panic(err)
		}
		if countDigits(line) > 15 {
			panic(fmt.Errorf("Too many digits in line: %v", line))
		}
		card := stripLine(line)
		for i := len(card); i < 15; i++ {
			card = card + "1" // strconv.Itoa(rand.Intn(9))
		}
		card = card + creditcard.GenerateLastDigit(card)
		if !creditcard.Validate(card) {
			panic(fmt.Errorf("ERROR: generated bad card: %s", card))
		}
		fmt.Println(prettyCard(card), "\t", creditcard.Cardtype(card))
	}
}

func prettyCard(card string) string {
	var p string
	for i, r := range card {
		if i > 0 && i%4 == 0 {
			p = p + " "
		}
		p = p + string(r)
	}

	return p
}

func stripLine(line string) string {
	var card = ""
	for _, r := range line {
		if !unicode.IsDigit(r) {
			continue
		}
		card = card + string(r)
	}
	return card
}

func countDigits(s string) int {
	var length = 0
	for _, v := range s {
		if !unicode.IsDigit(v) {
			continue
		}
		length += 1
	}
	return length
}
