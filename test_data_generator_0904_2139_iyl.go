// 代码生成时间: 2025-09-04 21:39:06
package main

import (
    "fmt"
    "log"
    "math/rand"
    "time"
)

// TestData contains attributes that simulate test data
type TestData struct {
    ID        int    "json:"id""
    Name      string "json:"name""
    Email     string "json:"email""
    Age       int    "json:"age""
    IsActive  bool   "json:"is_active""
}

// GenerateTestData generates a slice of TestData with the specified count
func GenerateTestData(count int) ([]TestData, error) {
    if count <= 0 {
        return nil, fmt.Errorf("count must be greater than 0")
    }

    testDataList := make([]TestData, count)
    rand.Seed(time.Now().UnixNano()) // Seed with current time

    for i := 0; i < count; i++ {
        testDataList[i] = TestData{
            ID:        i + 1,
            Name:      fmt.Sprintf("TestUser%d", i+1),
            Email:     fmt.Sprintf("testuser%d@example.com", i+1),
            Age:       rand.Intn(100) + 1, // Random age between 1 and 100
            IsActive:  true,
        }
    }
    return testDataList, nil
}

func main() {
    count := 10 // Number of test data entries to generate
    testData, err := GenerateTestData(count)
    if err != nil {
        log.Fatalf("Error generating test data: %v", err)
    }

    // Output the generated test data
    for _, data := range testData {
        fmt.Printf("ID: %d, Name: %s, Email: %s, Age: %d, IsActive: %t
",
            data.ID, data.Name, data.Email, data.Age, data.IsActive)
    }
}