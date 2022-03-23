package hello

import (
	"rsc.io/quote/v3"
)

// hello
func Hello() string {
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}
