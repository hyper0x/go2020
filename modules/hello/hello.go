package hello

import "rsc.io/quote/v3"

// Hello 会返回一句问候语。
func Hello() string {
	return quote.HelloV3()
}

// Proverb 会返回一句格言。
func Proverb() string {
	return quote.Concurrency()
}
