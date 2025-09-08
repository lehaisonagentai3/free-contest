#!/bin/bash

# Simple API Test Script
# Tests all Free Contest API endpoints

BASE_URL="http://localhost:8080"

echo "🧪 Testing Free Contest API"
echo "=============================="

echo -e "\n1️⃣  Health Check"
curl -s "$BASE_URL/health" | jq '{status: .status}'

echo -e "\n2️⃣  Get All Units"
curl -s "$BASE_URL/api/v1/units" | jq '{count: .count, sample_unit: .data[0]}'

echo -e "\n3️⃣  Get All Officers"
curl -s "$BASE_URL/api/v1/officers" | jq '{count: .count, sample_officer: .data[0]}'

echo -e "\n4️⃣  Create Test for Officer"
echo "Creating test for officerID=1, subjectID=1..."
TEST_RESPONSE=$(curl -s "$BASE_URL/api/v1/tests/officer-subject?officerID=1&subjectID=1")
TEST_ID=$(echo "$TEST_RESPONSE" | jq -r '.id')
echo "$TEST_RESPONSE" | jq '{test_id: .id, officer: .officer.name, subject: .subject.name, question_count: (.questions | length), start_time: .start_time}'

echo -e "\n5️⃣  Start Test"
if [ "$TEST_ID" != "null" ]; then
    echo "Starting test ID: $TEST_ID for officer 1..."
    START_RESPONSE=$(curl -s -X POST "$BASE_URL/api/v1/tests/start?officerID=1&testID=$TEST_ID")
    echo "$START_RESPONSE" | jq '{test_id: .id, start_time: .start_time, status: (if .start_time > 0 then "started" else "not started" end)}'
else
    echo "❌ Cannot start test - invalid test ID"
fi

echo -e "\n6️⃣  Submit Test"
if [ "$TEST_ID" != "null" ]; then
    echo "Submitting answers for test ID: $TEST_ID..."
    # Get first few question IDs from the test
    QUESTIONS=$(echo "$TEST_RESPONSE" | jq -r '.questions[:3] | map(.id) | @csv' | tr -d '"')
    echo "Sample questions: $QUESTIONS"
    # Submit some sample answers
    SUBMIT_RESPONSE=$(curl -s -X POST "$BASE_URL/api/v1/tests/submit?officerID=1&testID=$TEST_ID" \
        -H "Content-Type: application/json" \
        -d '{"1":"A","2":"B","3":"C","4":"A","5":"B"}')
    echo "$SUBMIT_RESPONSE" | jq '{submission_id: .id, score: .score, submitted_at: .submitted_at, answers_count: (.answers | length)}'
else
    echo "❌ Cannot submit test - invalid test ID"
fi

echo -e "\n7️⃣  Error Handling Examples"
echo "Testing invalid officer ID:"
curl -s "$BASE_URL/api/v1/tests/officer-subject?officerID=999&subjectID=1" | jq '.error'

echo "Testing missing parameters:"
curl -s "$BASE_URL/api/v1/tests/officer-subject" | jq '.error'

echo -e "\n✅ All tests completed!"
echo -e "\n📚 Available Endpoints:"
echo "  GET  /health"
echo "  GET  /api/v1/units"
echo "  GET  /api/v1/officers"
echo "  GET  /api/v1/tests/officer-subject?officerID=X&subjectID=Y"
echo "  POST /api/v1/tests/start?officerID=X&testID=Y"
echo "  POST /api/v1/tests/submit?officerID=X&testID=Y (with JSON body)"

echo -e "\n📖 Swagger Documentation:"
echo "  http://localhost:8080/v1/api/free-contest/swagger/index.html"
