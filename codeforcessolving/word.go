package main

import (
	"fmt"
	"strings"
	"unicode"
)
// https://codeforces.com/problemset/problem/59/A

func main() {
	var theStr string;
	fmt.Scan(&theStr);

	lowerCount := 0;
	upperCount := 0

	for _, ch := range theStr {
		if (unicode.IsLower(ch)) {
			lowerCount++;
		} else {
			upperCount++;
		}
	}

	if (lowerCount == upperCount || lowerCount > upperCount) {
		ans := strings.ToLower(theStr);
		fmt.Println(ans);
	} else {
		ans := strings.ToUpper(theStr);
		fmt.Println(ans);
	}
}