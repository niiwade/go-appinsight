package providers

import (
	"appinsights/test/models"
	"errors"
	"os"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

var telemetryClient = appinsights.NewTelemetryClient(os.Getenv("InstrumentationKey"))

var students = []models.Students{
	{ID: "1", Name: "John Doe", Level: "Freshman", Programme: 1},
}

func GetAllStudents() []models.Students {
	return students
}

func GetStudentByID(id string) (models.Students, error) {
	for _, student := range students {
		if student.ID == id {
			return student, nil
		}
	}
	return models.Students{}, errors.New("student not found")
}

func CreateStudent(student models.Students) {
	students = append(students, student)
	// Track custom event
	event := appinsights.NewEventTelemetry("StudentCreated")
	event.Properties["StudentID"] = student.ID
	telemetryClient.TrackEvent("Create Student")
}

func UpdateStudent(id string, updatedStudent models.Students) error {
	for i, student := range students {
		if student.ID == id {
			students[i] = updatedStudent
			// Track custom event
			event := appinsights.NewEventTelemetry("StudentUpdated")
			event.Properties["StudentID"] = updatedStudent.ID
			telemetryClient.TrackEvent("Update Student")
			return nil
		}
	}
	return errors.New("student not found")
}

func DeleteStudent(id string) error {
	for i, student := range students {
		if student.ID == id {
			students = append(students[:i], students[i+1:]...)
			// Track custom event
			event := appinsights.NewEventTelemetry("StudentDeleted")
			event.Properties["StudentID"] = id
			telemetryClient.TrackEvent("Update Student")
			return nil
		}
	}
	return errors.New("student not found")
}
