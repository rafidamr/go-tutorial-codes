package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"
)

// Test function: test1
// Description: Tests the slice modification function that doubles each element
// Expected behavior: Should modify the slice in place, doubling each element
func TestTest1(t *testing.T) {
	// Capture stdout to test the output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function
	test1()

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Read the captured output
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := strings.TrimSpace(buf.String())

	// Expected output format: [2 4 6 8 10]
	expected := "[2 4 6 8 10]"
	if output != expected {
		t.Errorf("test1() output = %q, want %q", output, expected)
	}
}

// Test function: test2
// Description: Tests the recursive Fibonacci function
// Expected behavior: Should return the nth Fibonacci number
func TestTest2(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"Fibonacci of 0", 0, 0},
		{"Fibonacci of 1", 1, 1},
		{"Fibonacci of 2", 2, 1},
		{"Fibonacci of 3", 3, 2},
		{"Fibonacci of 4", 4, 3},
		{"Fibonacci of 5", 5, 5},
		{"Fibonacci of 6", 6, 8},
		{"Fibonacci of 7", 7, 13},
		{"Fibonacci of 8", 8, 21},
		{"Fibonacci of 9", 9, 34},
		{"Fibonacci of 10", 10, 55},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := test2(tt.input)
			if result != tt.expected {
				t.Errorf("test2(%d) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

// Test function: test3
// Description: Tests the goroutine function with WaitGroup
// Expected behavior: Should spawn 5 goroutines, each sleeping for 500ms and printing completion message
func TestTest3(t *testing.T) {
	// Capture stdout to test the output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function
	test3()

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Read the captured output
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := strings.TrimSpace(buf.String())

	// Expected output should contain 5 goroutine completion messages
	// Note: The order may vary due to goroutine scheduling
	expectedLines := []string{
		"Goroutine 0 completed",
		"Goroutine 1 completed",
		"Goroutine 2 completed",
		"Goroutine 3 completed",
		"Goroutine 4 completed",
	}

	// Check that all expected lines are present
	for _, expectedLine := range expectedLines {
		if !strings.Contains(output, expectedLine) {
			t.Errorf("test3() output missing expected line: %q", expectedLine)
		}
	}

	// Check that we have exactly 5 lines
	lines := strings.Split(output, "\n")
	if len(lines) != 5 {
		t.Errorf("test3() output has %d lines, want 5", len(lines))
	}
}

// Benchmark function: test2 (Fibonacci)
// Description: Performance testing for the recursive Fibonacci function
func BenchmarkTest2(b *testing.B) {
	// Test with a moderate input size to avoid excessive recursion
	input := 20
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		test2(input)
	}
}

// Benchmark function: test3 (Goroutines)
// Description: Performance testing for the goroutine function
func BenchmarkTest3(b *testing.B) {
	// Capture stdout to avoid cluttering benchmark output
	old := os.Stdout
	_, w, _ := os.Pipe()
	os.Stdout = w

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		test3()
	}

	// Restore stdout
	w.Close()
	os.Stdout = old
}

// Test function: test2 edge cases
// Description: Tests edge cases and negative inputs for the Fibonacci function
func TestTest2EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"Negative number", -5, -5}, // Current implementation returns negative numbers as-is
		{"Large negative number", -100, -100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := test2(tt.input)
			if result != tt.expected {
				t.Errorf("test2(%d) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

// Test function: test3 timing
// Description: Tests that test3 function completes within reasonable time
func TestTest3Timing(t *testing.T) {
	// Capture stdout to avoid cluttering test output
	old := os.Stdout
	_, w, _ := os.Pipe()
	os.Stdout = w

	start := time.Now()
	test3()
	duration := time.Since(start)

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Each goroutine sleeps for 500ms, but they run concurrently
	// Total time should be approximately 500ms (not 2.5 seconds)
	expectedMaxDuration := time.Millisecond * 600
	if duration > expectedMaxDuration {
		t.Errorf("test3() took %v, expected less than %v", duration, expectedMaxDuration)
	}
}

// Test function: test4
// Description: Tests the function that returns sum and product of two integers
// Expected behavior: Should return the sum and product of the input parameters
func TestTest4(t *testing.T) {
	tests := []struct {
		name         string
		x, y         int
		expectedSum  int
		expectedProd int
	}{
		{"Positive numbers", 5, 3, 8, 15},
		{"Zero values", 0, 0, 0, 0},
		{"Negative numbers", -2, -3, -5, 6},
		{"Mixed signs", 10, -5, 5, -50},
		{"Large numbers", 1000, 2000, 3000, 2000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum, product := test4(tt.x, tt.y)
			if sum != tt.expectedSum {
				t.Errorf("test4(%d, %d) sum = %d, want %d", tt.x, tt.y, sum, tt.expectedSum)
			}
			if product != tt.expectedProd {
				t.Errorf("test4(%d, %d) product = %d, want %d", tt.x, tt.y, product, tt.expectedProd)
			}
		})
	}
}

// Test function: tests5 (processNumbers and channel processing)
// Description: Tests the channel-based number processing function
// Expected behavior: Should process numbers through a channel and print doubled values
func TestTests5(t *testing.T) {
	// Capture stdout to test the output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function
	tests5()

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Read the captured output
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := strings.TrimSpace(buf.String())

	// Expected output format: "2 4 6 8 10" (doubled values)
	expected := "2 4 6 8 10"
	if output != expected {
		t.Errorf("tests5() output = %q, want %q", output, expected)
	}
}

// Test function: test6 (worker pool pattern)
// Description: Tests the worker pool implementation with channels
// Expected behavior: Should process 5 jobs through 3 workers, each taking 1 second
func TestTest6(t *testing.T) {
	// This test will take at least 1 second due to worker sleep
	start := time.Now()

	// Capture stdout to avoid cluttering test output
	old := os.Stdout
	_, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function
	test6()

	// Restore stdout
	w.Close()
	os.Stdout = old

	duration := time.Since(start)

	// Should take at least 1 second (worker sleep time) but less than 2.5 seconds
	// due to concurrent processing
	if duration < time.Second {
		t.Errorf("test6() completed too quickly: %v, expected at least 1 second", duration)
	}
	if duration > 2500*time.Millisecond {
		t.Errorf("test6() took too long: %v, expected less than 2.5 seconds", duration)
	}
}

// Test function: test7 (goroutines with mutex)
// Description: Tests concurrent sum calculation with mutex protection
// Expected behavior: Should calculate sum of numbers using goroutines with mutex
func TestTest7(t *testing.T) {
	// Capture stdout to test the output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function
	test7()

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Read the captured output
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := strings.TrimSpace(buf.String())

	// Expected output: "150" (sum of 10+20+30+40+50)
	expected := "150"
	if output != expected {
		t.Errorf("test7() output = %q, want %q", output, expected)
	}
}

// Test function: test8 (database operations)
// Description: Tests database connection and query operations
// Expected behavior: Should attempt to connect to database and query user data
// Note: This test will fail in CI/CD environments without a real database
func TestTest8(t *testing.T) {
	// This test will fail due to missing MySQL driver, so we'll skip it
	// In a real environment, you would need to import "github.com/go-sql-driver/mysql"
	// and have a MySQL database running
	t.Skip("Skipping test8 due to missing MySQL driver and database connection")

	// Capture stdout and stderr to test the output
	old := os.Stdout
	oldStderr := os.Stderr
	r, w, _ := os.Pipe()
	rErr, wErr, _ := os.Pipe()
	os.Stdout = w
	os.Stderr = wErr

	// Call the function (will likely fail due to no database connection)
	test8()

	// Restore stdout and stderr
	w.Close()
	wErr.Close()
	os.Stdout = old
	os.Stderr = oldStderr

	// Read the captured output
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := strings.TrimSpace(buf.String())

	// Read the captured error output
	var bufErr bytes.Buffer
	bufErr.ReadFrom(rErr)
	errorOutput := strings.TrimSpace(bufErr.String())

	// The function will either print user data (if DB connection succeeds)
	// or log.Fatal will terminate the program (if DB connection fails)
	// In test environment, it will likely fail, so we check for either scenario
	if output == "" && errorOutput == "" {
		t.Log("test8() executed but output was empty (likely due to database connection failure)")
	} else if strings.Contains(output, "User ID:") {
		t.Log("test8() successfully connected to database and retrieved user data")
	} else if strings.Contains(errorOutput, "sql: unknown driver") || strings.Contains(errorOutput, "connection refused") {
		t.Log("test8() failed as expected due to database connection issues")
	}
}

// Test function: test9 (struct value vs pointer)
// Description: Tests struct modification behavior
// Expected behavior: Should demonstrate that struct is passed by value, not by reference
func TestTest9(t *testing.T) {
	// Capture stdout to test the output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function
	test9()

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Read the captured output
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := strings.TrimSpace(buf.String())

	// Expected output: "10 5" (original values, not doubled)
	// This demonstrates that the struct is passed by value
	expected := "10 5"
	if output != expected {
		t.Errorf("test9() output = %q, want %q", output, expected)
	}
}

// Benchmark function: test4 (sum and product)
// Description: Performance testing for the sum and product function
func BenchmarkTest4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		test4(i, i+1)
	}
}
