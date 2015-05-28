package main

import (
	"github.com/h2object/pb"
	"time"
)

func main() {
	count := 5000
	bar := pb.New(count)

	bar.Prefix("Markdown files: ")
	// show percents (by default already true)
	bar.ShowPercent = true

	// show bar (by default already true)
	bar.ShowPercent = true

	// no need counters
	bar.ShowCounters = true

	bar.ShowTimeLeft = true

	// and start
	bar.Start()
	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.FinishPrint("The End!")
}
