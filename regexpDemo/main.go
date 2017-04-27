package main

import (
	"fmt"
	"regexp"
	"strings"
)

const langRawStrExp = `^([a-zA-Z]+-[a-zA-Z]+)(?:_[a-zA-Z]+$)?`

func main() {
	var src = []string{"en-US", "zh-TW", "zh-CN", "zh-CN_Hant", "zh-Hans"}
	re := regexp.MustCompile(langRawStrExp)

	for _, v := range src {
		fmt.Println(strings.ToLower(re.ReplaceAllString(v, "$1")))
	}

}
