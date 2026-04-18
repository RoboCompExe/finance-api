# Finance API

Simple backend system for financial transaction simulation.

## Features
- Account balance check
- Money transfer (ACID transaction)
- PostgreSQL integration

## Run
1. Setup PostgreSQL
2. go run main.go

## Design
- Uses transaction to ensure consistency
- Row locking prevents race condition
