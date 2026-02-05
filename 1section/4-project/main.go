package main

import "fmt"

type LogLevel int

const (
	LogInfo LogLevel = iota
	LogDebug
	LogFatal
	LevelWarning
	LevelError
)

var levelNames = []string{"TRACE", "DEBUG", "INFO", "WARNING", "ERROR"}

func main() {
	printLogLevel(LogInfo)
	printLogLevel(LogDebug)
	printLogLevel(LogFatal)
	printLogLevel(LevelWarning)
	printLogLevel(LevelError)
	printLogLevel(10)
}

func (l LogLevel) String() string {
	if l < LogInfo || l > LevelError {
		return "Unknown"
	}

	return levelNames[l]
}

func printLogLevel(level LogLevel) {
	fmt.Printf("Log level: %d %s\n", level, level.String())
}
