package main

import (
	"strings"
)

const (
	// https://patorjk.com/software/taag/#p=display&f=Colossal&t=deci&x=none&v=4&h=4&w=80&we=false
	helpText = `
        888                   d8b
        888                   Y8P
        888
    .d88888  .d88b.   .d8888b 888
   d88" 888 d8P  Y8b d88P"    888
   888  888 88888888 888      888
   Y88b 888 Y8b.     Y88b.    888
    "Y88888  "Y8888   "Y8888P 888

"your second-to-last next editor"
---------------------------------

deci is a text editor like nano.

source: https://github.com/ethmarks/deci/
site: https://ethmarks.github.io/deci/
`
)

func getHelpLines() []string {

	s := helpText

	return strings.Split(s, "\n")[1:]
}
