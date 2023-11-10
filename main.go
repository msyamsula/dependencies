package dependencies

import (
	"rsc.io/quote/v3"
)

func Hello() string {
	return "this is beta in master"
}

func Proverb() string {
	return quote.Concurrency()
}
