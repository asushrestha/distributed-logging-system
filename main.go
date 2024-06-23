package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//stores the logs
var logs = Logs{}

/*
* Function to handle logs
*/
func handleLogs(c *gin.Context) {
	switch c.Request.Method {
    case http.MethodPost:
        handlePostLog(c)
    case http.MethodGet:
        handleGetLogs(c)
    default:
        c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
    }
}

/*
* Function to handle post operation of logs.
*/
func handlePostLog(c *gin.Context){
	var log Log
    if err := c.ShouldBindJSON(&log); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    id := logs.AddLog(log)
    c.JSON(http.StatusCreated, gin.H{"id": id})
}

/*
* Function to handle fetch operation of logs.
*/
func handleGetLogs(c *gin.Context){
	var filter LogFilter
    serviceName := c.Query("serviceName")
    severity := c.Query("severity")
    startTimeString := c.Query("startTime")
    endTimeString := c.Query("endTime")

	//function to parse time
    parseTime := func(timeString string) (time.Time, error) {
        if timeString == "" {
            return time.Time{}, nil // Return zero value if string is empty
        }
        return time.Parse(time.RFC3339, timeString)
    }

    var err error
    filter.StartTime, err = parseTime(startTimeString)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start time format"})
        return
    }

    filter.EndTime, err = parseTime(endTimeString)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end time format"})
        return
    }

    if severity != "" {
        filter.Severity = severity
    }

    if serviceName != "" {
        filter.ServiceName = serviceName
    }

    logs := logs.QueryLogs(filter)
    c.JSON(http.StatusOK, logs)
}

func service(serviceName string, url string) {

	for {
		severity := []string{"INFO", "WARN", "ERROR"}[rand.Intn(3)]
		log := Log{
			ServiceName: serviceName,
			Severity:    severity,
			Message:     fmt.Sprintf("This is a %s message.", severity),
		}
		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(log)
		_, err := http.Post(url, "application/json", buf)
		if err != nil {
			fmt.Println("Error logging", err)
		}
		//define the sleep interval in seconds
		seconds := 5

		//sets the sleep time
		time.Sleep(time.Duration(seconds) * time.Second)
	}
}

func main() {
	r := gin.Default()

	//endpoint to post log.
	r.POST("/logs", handleLogs)

	//endpoint to get log.
	r.GET("/logs", handleLogs)

	//URL to post log for the service.
	postLogUrl := "http://localhost:8081/logs"

	//go routine call with a fixed time interval for log generation.
	go service("LoggerServiceA", postLogUrl)
	go service("LoggerServiceB", postLogUrl)
	go service("LoggerServiceC", postLogUrl)

	log.Fatal(r.Run(":8081"))
}
