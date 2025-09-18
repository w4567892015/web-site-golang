package logger

import (
	"fmt"
	"time"
)

func Log(serviceName, message string) {
	fmt.Printf("[%s] %s: %s\n", time.Now().Format(time.RFC3339), serviceName, message)
}
