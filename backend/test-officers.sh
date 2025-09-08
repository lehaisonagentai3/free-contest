#!/bin/bash

# Simple Officer API Test Script

echo "🔍 Testing Officer API Endpoints"
echo "================================="
echo ""

BASE_URL="http://localhost:8080"

echo "1. Get All Officers:"
curl -s "$BASE_URL/api/v1/officers" | jq '.data[] | {id, name, unit: .unit.name}'
echo ""

echo "2. Get Individual Officers:"
for id in 1 2 3; do
    echo "Officer ID $id:"
    curl -s "$BASE_URL/api/v1/officers/$id" | jq '.data | {id, name, unit: .unit.name}'
    echo ""
done

echo "3. Error Cases:"
echo "Non-existent officer (ID 999):"
curl -s "$BASE_URL/api/v1/officers/999" | jq .
echo ""

echo "Invalid officer ID format (abc):"
curl -s "$BASE_URL/api/v1/officers/abc" | jq .
echo ""

echo "✅ Officer API testing completed!"
