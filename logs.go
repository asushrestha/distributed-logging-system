package main

import (
	"sync"
	"time"
)

//struct for Log
type Log struct {
	ID          int64     `json:"id"`
	Timestamp   time.Time `json:"timestamp"`
	Severity    string    `json:"severity"`
	ServiceName string    `json:"serviceName"`
	Message     string    `json:"message"`
}

//struct for Logs
type Logs struct {
	logs   []Log
	mu     sync.Mutex
	nextId int64
}

//struct for LogFilter
type LogFilter struct {
	StartTime   time.Time
	EndTime   	time.Time
	Severity    string
	ServiceName string
}

//Add a new log
func (logs *Logs) AddLog(log Log) int64 {
	logs.mu.Lock()
	defer logs.mu.Unlock()
	log.ID = logs.nextId
	logs.nextId++
	log.Timestamp = time.Now().UTC()
	logs.logs = append(logs.logs, log)
	return log.ID
}

func (f LogFilter) Matches(log Log) bool {
	return ((f.StartTime.IsZero() || log.Timestamp.After(f.StartTime)) && (f.EndTime.IsZero() || log.Timestamp.Before(f.EndTime)) && (f.Severity == "" || f.Severity == log.Severity) && (f.ServiceName == "" || f.ServiceName == log.ServiceName))
}

func (logs *Logs) QueryLogs(filter LogFilter) []Log {
	logs.mu.Lock()
	defer logs.mu.Unlock()
	var result []Log
	for _, log := range logs.logs {
		if filter.Matches(log) {
			result = append(result, log)
		}
	}
	return result
}