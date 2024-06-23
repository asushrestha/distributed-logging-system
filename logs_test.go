package main

import (
	"testing"
	"time"
)

// Test for QueryLogs method
func TestQueryLogs(t *testing.T) {
	//current datetimestamp
	now := time.Now()

	//Initial logs
	logs := &Logs{
		logs: []Log{
			{Timestamp: now.AddDate(0,0,1), Severity: "ERROR", ServiceName: "LoggerServiceA"},
			{Timestamp: now, Severity: "INFO", ServiceName: "LoggerServiceB"},
			{Timestamp: now.AddDate(0,0,1), Severity: "ERROR", ServiceName: "LoggerServiceA"},
		},
	}

	//test values for LogFilter and expected length.
	tests := []struct {
		filter   LogFilter
		expected int
	}{
		{LogFilter{StartTime: now.AddDate(0,0,-2), EndTime: now.AddDate(0,0,2)}, 3},
		{LogFilter{Severity: "ERROR"}, 2},
		{LogFilter{ServiceName: "LoggerServiceA"}, 2},
		{LogFilter{ServiceName: "LoggerServiceB"}, 1},
	}
	
	for _, tt := range tests {
		result := logs.QueryLogs(tt.filter)
		if len(result) != tt.expected {
			t.Errorf("QueryLogs() = %v, expected %v", len(result), tt.expected) // prints the expected and length of result
		}
	}
}

// Test for AddLog method
func TestAddLog(t *testing.T) {
	logs := &Logs{}

	// Adding first log
	log1 := Log{
		Severity:    "INFO",
		ServiceName: "LoggerTestServiceA",
		Message:     "This is First log entry",
	}
	id1 := logs.AddLog(log1)

	//check for id.
	if id1 != 0 {
		t.Errorf("Expected ID 0, got %d", id1)
	}

	//check for logs length.
	if len(logs.logs) != 1 {
		t.Errorf("Expected logs length 1, got %d", len(logs.logs))
	}

	//check for the severity, service,message.
	if logs.logs[0].Severity != log1.Severity || logs.logs[0].ServiceName != log1.ServiceName || logs.logs[0].Message != log1.Message {
		t.Errorf("Log content mismatch")
	}
	
	//check for timestamp
	if logs.logs[0].Timestamp.IsZero() {
		t.Errorf("Timestamp not set")
	}

	// Adding second log
	log2 := Log{
		Severity:    "ERROR",
		ServiceName: "LoggerTestServiceB",
		Message:     "This is Second log entry",
	}
	time.Sleep(1 * time.Second) // Ensure a different timestamp
	id2 := logs.AddLog(log2)

	//check for id.
	if id2 != 1 {
		t.Errorf("Expected ID 1, got %d", id2)
	}

	//check for logs length.
	if len(logs.logs) != 2 {
		t.Errorf("Expected logs length 2, got %d", len(logs.logs))
	}

	//check for the severity, service,message.
	if logs.logs[1].Severity != log2.Severity || logs.logs[1].ServiceName != log2.ServiceName || logs.logs[1].Message != log2.Message {
		t.Errorf("Log content mismatch")
	}

	//check for timestamp
	if logs.logs[1].Timestamp.IsZero() || logs.logs[1].Timestamp == logs.logs[0].Timestamp {
		t.Errorf("Timestamp not properly set or same as previous log")
	}
}