#!/bin/bash

# Build the frontend and output to the backend UI directory
echo "Building React frontend..."
npm run build

echo "Moving build output to backend UI directory..."
# Remove existing UI directory if it exists
rm -rf ../backend/cmd/api/ui

# Create the UI directory structure
mkdir -p ../backend/cmd/api/ui

# Copy the built files to the backend UI directory
cp -r build/* ../backend/cmd/api/ui/

echo "Build completed successfully!"
echo "Frontend files have been deployed to ../backend/cmd/api/ui"