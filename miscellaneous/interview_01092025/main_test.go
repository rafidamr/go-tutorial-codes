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
