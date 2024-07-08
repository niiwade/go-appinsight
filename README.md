
# Student API with Application Insights Instrumentation

## Introduction

Welcome to the Student API project! This project demonstrates how to build a simple RESTful API using Go, Gin framework, and Application Insights for monitoring and telemetry. The API provides endpoints to manage student records, including creating, retrieving, updating, and deleting student information. Additionally, it integrates Application Insights to track and monitor the performance and usage of the endpoints.

## Table of Contents

-   [Introduction](#introduction)
-   [Features](#features)
-   [Prerequisites](#prerequisites)
-   [Installation](#installation)
-   [Usage](#usage)
-   [API Endpoints](#api-endpoints)
-   [Application Insights](#application-insights)
-   [Contributing](#contributing)
-   [License](#license)

## Features

-   CRUD operations for student records
-   Integration with Application Insights for telemetry
-   Middleware for tracking request performance
-   Error tracking and logging

## Prerequisites

-   Go 1.16 or later
-   A valid Application Insights Instrumentation Key

## Installation

1.  Clone the repository:
    
    bash
    
    Copy code
    
    `git clone https://github.com/yourusername/student-api.git
    cd student-api` 
    
2.  Install dependencies:
    
    bash
    
    Copy code
    
    `go mod tidy` 
    
3.  Set your Application Insights Instrumentation Key:
    
    bash
    
    Copy code
    
    `export INSTRUMENTATION_KEY=your_instrumentation_key` 
    

## Usage

1.  Run the application:
    
    bash
    
    Copy code
    
    `go run main.go` 
    
2.  The API will be available at `http://localhost:8080`.
    

## API Endpoints

-   `GET /students`: Retrieve all students
-   `GET /students/:id`: Retrieve a student by ID
-   `POST /students`: Create a new student
-   `PUT /students/:id`: Update a student by ID
-   `DELETE /students/:id`: Delete a student by ID
