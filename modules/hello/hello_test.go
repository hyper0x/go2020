package hello

import "testing"

func TestHello(t *testing.T) {
	want1 := "Hello, world."
	want2 := "你好，世界。"
	if got := Hello(); got != want1 && got != want2 {
		t.Errorf("Hello() = %q, want %q or %q", got, want1, want2)
	}
}

func TestProverb(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := Proverb(); got != want {
		t.Errorf("Proverb() = %q, want %q", got, want)
	}
}
