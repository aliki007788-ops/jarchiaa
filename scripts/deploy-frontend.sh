#!/bin/bash

echo "?? Deploying frontend..."
cd frontend
npm ci
npm run build
aws s3 sync dist/ s3://jarchia-cdn/
echo "? Frontend deployed"


