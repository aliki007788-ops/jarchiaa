#!/bin/bash

echo "?? Generating Swagger documentation..."
cd backend
swag init -g cmd/main.go
echo "? Swagger docs generated"


