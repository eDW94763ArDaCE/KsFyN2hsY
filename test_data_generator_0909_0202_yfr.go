// 代码生成时间: 2025-09-09 02:02:38
package main

import (
    "fmt"
    "math/rand"
    "time"
)

// TestDataGenerator is a struct that represents the test data generator.
type TestDataGenerator struct {
    // No additional fields are needed for this basic generator.
}

// NewTestDataGenerator creates a new instance of TestDataGenerator.
func NewTestDataGenerator() *TestDataGenerator {
    return &TestDataGenerator{}
}

// GenerateRandomString generates a random alphanumeric string of a given length.
func (g *TestDataGenerator) GenerateRandomString(length int) string {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
    b := make([]byte, length)
    for i := range b {
        b[i] = charset[seededRand.Intn(len(charset))]
    }
    return string(b)
}

// GenerateRandomNumber generates a random number within the given range.
func (g *TestDataGenerator) GenerateRandomNumber(min, max int) int {
    if min > max {
        return -1 // Error: min cannot be greater than max.
    }
    return rand.Intn(max-min) + min
}

// Main function to demonstrate the usage of TestDataGenerator.
func main() {
    // Initialize the test data generator.
    generator := NewTestDataGenerator()

    // Generate and print a random string of length 10.
    randomString := generator.GenerateRandomString(10)
    fmt.Println("Random String: ", randomString)

    // Generate and print a random number between 1 and 100.
    randomNumber := generator.GenerateRandomNumber(1, 100)
    fmt.Println("Random Number: ", randomNumber)
}
