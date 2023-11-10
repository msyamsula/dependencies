package v2

import (
	"rsc.io/quote/v3"
)

func Hello() string {
	return "this is v2"
}

func Proverb() string {
	return quote.Concurrency()
}
