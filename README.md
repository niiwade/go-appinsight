Student API with Application Insights Instrumentation
Introduction
Welcome to the Student API project! This project demonstrates how to build a simple RESTful API using Go, Gin framework, and Application Insights for monitoring and telemetry. The API provides endpoints to manage student records, including creating, retrieving, updating, and deleting student information. Additionally, it integrates Application Insights to track and monitor the performance and usage of the endpoints.

Table of Contents
Introduction
Features
Prerequisites
Installation
Usage
API Endpoints
Application Insights
Contributing
License
Features
CRUD operations for student records
Integration with Application Insights for telemetry
Middleware for tracking request performance
Error tracking and logging
Prerequisites
Go 1.16 or later
A valid Application Insights Instrumentation Key
Installation
Clone the repository:

bash
Copy code
git clone https://github.com/yourusername/student-api.git
cd student-api
Install dependencies:

bash
Copy code
go mod tidy
Set your Application Insights Instrumentation Key:

bash
Copy code
export INSTRUMENTATION_KEY=your_instrumentation_key
Usage
Run the application:

bash
Copy code
go run main.go
The API will be available at http://localhost:8080.

API Endpoints
GET /students: Retrieve all students
GET /students/:id: Retrieve a student by ID
POST /students: Create a new student
PUT /students/:id: Update a student by ID
DELETE /students/:id: Delete a student by ID
