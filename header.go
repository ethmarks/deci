package main

import (
	"fmt"
)

const (
	headerTextLeft  = "  deci 0.0.1"
	headerTextRight = "by @ethmarks  "
)

func getHeader(termWidth int) string {
	if len(headerTextLeft)+len(headerTextRight) >= termWidth {
		return headerTextLeft + "" + headerTextRight
	}

	raw := fmt.Sprintf("%s%*s", headerTextLeft, termWidth-len(headerTextLeft), headerTextRight)

	formatted := inverseStyle.Render(raw)

	return formatted
}
