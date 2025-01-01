package config

import (
	"fmt"
)

var (
	logger *Logger
)

// Init initializes the MongoDB connection
func Init() error {
	var err error
	// Initialize MongoDB
	err = InitializeMongoDB()
	if err != nil {
		return fmt.Errorf("error initializing mongoDB: %v", err)
	}
	return nil
}

// GetLogger - Returns the Logger instance
func GetLogger(p string) *Logger {
	// Initialize logger
	logger = NewLogger(p)
	return logger
}
