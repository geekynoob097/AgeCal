# Reasoning & Design Decisions {#reasoning-design-decisions}

This document explains the approach, architecture, and key technical
decisions taken while building the **AgeCal** backend project as part
of an internship assignment.

## 🎯 Problem Understanding {#problem-understanding}

The goal was to build a RESTful API that: - Manages users with name
and date of birth (DOB) - Calculates age dynamically when
fetching user details - Follows clean backend practices suitable for a
production-grade service

A key constraint was that **age should not be stored in the database**,
as it changes over time and would lead to data inconsistency.

## 🧱 Architecture Choice {#architecture-choice}

I followed a **layered (clean) architecture**:

    Handler → Service → Repository → Database

### Why this architecture?

- **Separation of concerns**: Each layer has a single responsibility
- **Testability**: Business logic can be tested independently of HTTP or
  DB
- **Maintainability**: Easy to extend or modify without breaking other
  layers
- **Industry-standard**: Commonly used in real-world Go backend systems

## 📦 Layer Responsibilities {#layer-responsibilities}

### Handler Layer

- Handles HTTP requests and responses
- Performs request parsing and input validation
- Maps domain errors to proper HTTP status codes

### Service Layer

- Contains business logic
- Calculates user age dynamically from DOB
- Returns domain-specific errors (no HTTP concepts)

### Repository Layer

- Handles database interactions
- Executes SQL queries
- Abstracts persistence logic from the service layer

## 🗄️ Database Design {#database-design}

The database schema intentionally stores only: - `id` - `name` - `dob`

**Age is not stored** because: - Age changes every year - Storing it
would introduce redundancy - Dynamic calculation ensures correctness at
all times

## ⏱️ Age Calculation Logic {#age-calculation-logic}

Age is calculated using Go's `time` package by: - Subtracting birth year
from the current year - Adjusting if the birthday has not occurred yet
this year

This logic is placed in the **service layer**, as it is part of the
business rules.

## ✅ Input Validation Strategy {#input-validation-strategy}

- Used `go-playground/validator` for declarative validation
- Validation rules are defined using struct tags
- Validation is handled in the **handler layer**, keeping services clean

This approach ensures: - Consistent validation rules - Clear error
messages - Proper `400 Bad Request` responses for invalid input

## 🌐 HTTP Status Code Decisions {#http-status-code-decisions}

Proper HTTP status codes are used to clearly communicate API outcomes:

- `201 Created` → successful resource creation
- `200 OK` → successful data retrieval
- `400 Bad Request` → invalid JSON or validation failure
- `404 Not Found` → resource does not exist
- `500 Internal Server Error` → unexpected server-side errors

This improves API clarity and client-side integration.

## 🧾 Logging & Observability {#logging-observability}

- Implemented structured logging using **Uber Zap**
- Added middleware-based request logging
- Injected a **Request ID** for every request to enable traceability

This design helps in: - Debugging issues - Tracing individual requests -
Observing system behavior in production-like environments

## 🧠 Key Design Principles Followed {#key-design-principles-followed}

- **Single Responsibility Principle**
- **Fail fast on invalid input**
- **Do not mix HTTP concerns with business logic**
- **Avoid redundant data storage**
- **Explicit initialization of shared resources (logger, validator)**

## 📌 Conclusion {#conclusion}

This project was designed to reflect **real-world backend engineering
practices**, not just a minimal working solution. Emphasis was placed on
clean code, correctness, and maintainability, making the project
suitable for both **internship evaluation** and **future extension**.
