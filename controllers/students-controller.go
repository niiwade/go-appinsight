package controllers

import (
	"appinsights/test/models"
	"appinsights/test/providers"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

var telemetryClient = appinsights.NewTelemetryClient(os.Getenv("InstrumentationKey"))

func GetStudents(c *gin.Context) {
	telemetryClient.TrackEvent("GetStudents_Start")

	startTime := time.Now()
	students := providers.GetAllStudents()
	duration := time.Since(startTime)

	requestTelemetry := appinsights.NewRequestTelemetry("GET", c.FullPath(), duration, "200")
	requestTelemetry.Properties["UserType"] = c.Request.Header.Get("UserType")
	requestTelemetry.Properties["OS"] = c.Request.Header.Get("User-Agent")

	c.JSON(http.StatusOK, students)
	telemetryClient.TrackEvent("GetStudents_Success")
}

func GetStudentByID(c *gin.Context) {
	telemetryClient.TrackEvent("GetStudentByID_Start")

	startTime := time.Now()
	id := c.Param("id")
	student, err := providers.GetStudentByID(id)
	duration := time.Since(startTime)

	if err != nil {
		telemetryClient.TrackException(err)
		requestTelemetry := appinsights.NewRequestTelemetry("GET", c.FullPath(), duration, "404")
		requestTelemetry.Properties["UserType"] = c.Request.Header.Get("UserType")
		requestTelemetry.Properties["OS"] = c.Request.Header.Get("User-Agent")

		c.JSON(http.StatusNotFound, gin.H{"message": "Student not found"})
		return
	}

	requestTelemetry := appinsights.NewRequestTelemetry("GET", c.FullPath(), duration, "200")
	requestTelemetry.Properties["UserType"] = c.Request.Header.Get("UserType")
	requestTelemetry.Properties["OS"] = c.Request.Header.Get("User-Agent")

	c.JSON(http.StatusOK, student)
	telemetryClient.TrackEvent("GetStudentByID_Success")
}

func CreateStudent(c *gin.Context) {
	telemetryClient.TrackEvent("CreateStudent_Start")

	startTime := time.Now()
	var newStudent models.Students
	if err := c.BindJSON(&newStudent); err != nil {
		telemetryClient.TrackException(err)
		requestTelemetry := appinsights.NewRequestTelemetry("POST", c.FullPath(), time.Since(startTime), "400")
		requestTelemetry.Properties["UserType"] = c.Request.Header.Get("UserType")
		requestTelemetry.Properties["OS"] = c.Request.Header.Get("User-Agent")

		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	providers.CreateStudent(newStudent)
	duration := time.Since(startTime)
	requestTelemetry := appinsights.NewRequestTelemetry("POST", c.FullPath(), duration, "201")
	requestTelemetry.Properties["UserType"] = c.Request.Header.Get("UserType")
	requestTelemetry.Properties["OS"] = c.Request.Header.Get("User-Agent")

	c.JSON(http.StatusCreated, newStudent)
	telemetryClient.TrackEvent("CreateStudent_Success")
}

func UpdateStudent(c *gin.Context) {
	telemetryClient.TrackEvent("UpdateStudent_Start")

	startTime := time.Now()
	id := c.Param("id")
	var updatedStudent models.Students
	if err := c.BindJSON(&updatedStudent); err != nil {
		telemetryClient.TrackException(err)
		requestTelemetry := appinsights.NewRequestTelemetry("PUT", c.FullPath(), time.Since(startTime), "400")
		requestTelemetry.Properties["UserType"] = c.Request.Header.Get("UserType")
		requestTelemetry.Properties["OS"] = c.Request.Header.Get("User-Agent")

		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := providers.UpdateStudent(id, updatedStudent); err != nil {
		telemetryClient.TrackException(err)
		requestTelemetry := appinsights.NewRequestTelemetry("PUT", c.FullPath(), time.Since(startTime), "404")
		requestTelemetry.Properties["UserType"] = c.Request.Header.Get("UserType")
		requestTelemetry.Properties["OS"] = c.Request.Header.Get("User-Agent")

		c.JSON(http.StatusNotFound, gin.H{"message": "Student not found"})
		return
	}

	duration := time.Since(startTime)
	requestTelemetry := appinsights.NewRequestTelemetry("PUT", c.FullPath(), duration, "200")
	requestTelemetry.Properties["UserType"] = c.Request.Header.Get("UserType")
	requestTelemetry.Properties["OS"] = c.Request.Header.Get("User-Agent")

	c.JSON(http.StatusOK, updatedStudent)
	telemetryClient.TrackEvent("UpdateStudent_Success")
}

func DeleteStudent(c *gin.Context) {
	telemetryClient.TrackEvent("DeleteStudent_Start")

	startTime := time.Now()
	id := c.Param("id")
	if err := providers.DeleteStudent(id); err != nil {
		telemetryClient.TrackException(err)
		requestTelemetry := appinsights.NewRequestTelemetry("DELETE", c.FullPath(), time.Since(startTime), "404")
		requestTelemetry.Properties["UserType"] = c.Request.Header.Get("UserType")
		requestTelemetry.Properties["OS"] = c.Request.Header.Get("User-Agent")

		c.JSON(http.StatusNotFound, gin.H{"message": "Student not found"})
		return
	}

	duration := time.Since(startTime)
	requestTelemetry := appinsights.NewRequestTelemetry("DELETE", c.FullPath(), duration, "200")
	requestTelemetry.Properties["UserType"] = c.Request.Header.Get("UserType")
	requestTelemetry.Properties["OS"] = c.Request.Header.Get("User-Agent")

	c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
	telemetryClient.TrackEvent("DeleteStudent_Success")
}
