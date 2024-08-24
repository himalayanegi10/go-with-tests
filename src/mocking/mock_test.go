package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountDown(t *testing.T) {

	t.Run("print 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		SpySleeper := &SpySleeper{}
		Countdown(buffer, SpySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if SpySleeper.Calls != 3 {
			t.Errorf("not enough calls to sleeper, want 3 got %d", SpySleeper.Calls)
		}			
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)
		want := []string{write, sleep, write, sleep, write, sleep, write}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}