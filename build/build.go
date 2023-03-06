package build

import (
	"fmt"
	"strings"
)

func AppBanner(app, version, gitHash, buildTime string) {
	width := 90
	fmt.Println(banner(width))
	fmt.Println(center("<-----LICENCE------>", width))
	fmt.Println(center(app, width))
	fmt.Println(center(fmt.Sprintf("%s<%s>", version, gitHash), width))
	fmt.Println(center(buildTime, width))
	fmt.Println(banner(width))
}

func center(s string, w int) string {
	return fmt.Sprintf("%*s", -w, fmt.Sprintf("%*s", (w+len(s))/2, s))
}

func banner(w int) string {
	return fmt.Sprintf("<%s>", strings.Repeat("=", w-2))
}
