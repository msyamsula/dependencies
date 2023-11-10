package dependencies

import (
	"rsc.io/quote/v3"
)

func Hello() string {
	return "this is alpha"
}

func Proverb() string {
	return quote.Concurrency()
}
