package dependencies

import (
	"rsc.io/quote/v3"
)

func Hello() string {
	return "this is stable version"
}

func Proverb() string {
	return quote.Concurrency()
}
