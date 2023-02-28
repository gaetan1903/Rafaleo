package utils

import (
	"os"
	"time"
)

func TypePrint(text string) {
	// print like a typewriter
	for index := range text {
		os.Stdout.WriteString(string(text[index]))
		time.Sleep(50 * time.Millisecond)
	}

	os.Stdout.WriteString("\n")
}
