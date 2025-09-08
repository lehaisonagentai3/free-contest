#!/bin/bash

# Free Contest API Test Script
# This script tests all available API endpoints

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# API Base URL
BASE_URL="http://localhost:8080"

# Function to print colored output
print_header() {
    echo -e "${BLUE}================================================${NC}"
    echo -e "${BLUE}$1${NC}"
    echo -e "${BLUE}================================================${NC}"
}

print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

print_info() {
    echo -e "${YELLOW}ℹ️  $1${NC}"
}

# Function to test API endpoint
test_endpoint() {
    local method=$1
    local endpoint=$2
    local description=$3
    local expected_status=$4
    
    echo -e "\n${YELLOW}Testing: $description${NC}"
    echo "Endpoint: $method $endpoint"
    
    if [ "$method" = "GET" ]; then
        response=$(curl -s -w "\n%{http_code}" "$BASE_URL$endpoint")
    elif [ "$method" = "POST" ]; then
        response=$(curl -s -w "\n%{http_code}" -X POST "$BASE_URL$endpoint")
    fi
    
    # Extract status code (last line) and body (everything else) - macOS compatible
    status_code=$(echo "$response" | tail -n1)
    body=$(echo "$response" | sed '$d')
    
    if [ "$status_code" = "$expected_status" ]; then
        print_success "Status: $status_code"
        echo "Response: $body" | jq . 2>/dev/null || echo "Response: $body"
    else
        print_error "Expected: $expected_status, Got: $status_code"
        echo "Response: $body"
    fi
}

# Start testing
print_header "FREE CONTEST API TESTING"

# 1. Health Check
print_header "1. HEALTH CHECK"
test_endpoint "GET" "/health" "Health check endpoint" "200"

# 2. Get All Units
print_header "2. GET ALL UNITS"
test_endpoint "GET" "/api/v1/units" "Get all units in the system" "200"

# 3. Get All Officers
print_header "3. GET ALL OFFICERS"
test_endpoint "GET" "/api/v1/officers" "Get all officers with unit information" "200"

# 4. Get Officer by ID
print_header "4. GET OFFICER BY ID"

print_info "Testing with valid officer ID"
test_endpoint "GET" "/api/v1/officers/1" "Get officer with ID 1" "200"

print_info "Testing with another valid officer ID"
test_endpoint "GET" "/api/v1/officers/2" "Get officer with ID 2" "200"

print_info "Testing with non-existent officer ID"
test_endpoint "GET" "/api/v1/officers/999" "Get officer with non-existent ID" "404"

print_info "Testing with invalid officer ID format"
test_endpoint "GET" "/api/v1/officers/abc" "Get officer with invalid ID format" "400"

# 5. Get All Subjects
print_header "5. GET ALL SUBJECTS"
test_endpoint "GET" "/api/v1/subjects" "Get all subjects without questions and chapters" "200"

# 6. Get Test for Officer and Subject
print_header "6. GET TEST FOR OFFICER AND SUBJECT"

print_info "Testing with valid parameters (officerID=1, subjectID=1)"
test_endpoint "GET" "/api/v1/tests/officer-subject?officerID=1&subjectID=1" "Get test for officer 1, subject 1" "200"

print_info "Testing with invalid officer ID"
test_endpoint "GET" "/api/v1/tests/officer-subject?officerID=999&subjectID=1" "Get test with invalid officer ID" "404"

print_info "Testing with invalid subject ID"
test_endpoint "GET" "/api/v1/tests/officer-subject?officerID=1&subjectID=999" "Get test with invalid subject ID" "404"

print_info "Testing with invalid parameter format"
test_endpoint "GET" "/api/v1/tests/officer-subject?officerID=abc&subjectID=1" "Get test with invalid parameter format" "400"

print_info "Testing with missing parameters"
test_endpoint "GET" "/api/v1/tests/officer-subject" "Get test with missing parameters" "400"

# 7. Start Test
print_header "7. START TEST"

print_info "First, creating a test to get a valid test ID..."
# Get a test and extract the test ID
test_response=$(curl -s "$BASE_URL/api/v1/tests/officer-subject?officerID=2&subjectID=2")
test_id=$(echo "$test_response" | jq -r '.id' 2>/dev/null)

if [ "$test_id" != "null" ] && [ "$test_id" != "" ]; then
    print_success "Created test with ID: $test_id"
    
    print_info "Testing start test with valid parameters"
    test_endpoint "POST" "/api/v1/tests/start?officerID=2&testID=$test_id" "Start test for officer 2" "200"
    
    print_info "Testing start same test again (should fail)"
    test_endpoint "POST" "/api/v1/tests/start?officerID=2&testID=$test_id" "Start already started test" "409"
else
    print_error "Failed to create test, skipping start test scenarios"
fi

print_info "Testing start test with invalid officer ID"
test_endpoint "POST" "/api/v1/tests/start?officerID=999&testID=123" "Start test with invalid officer ID" "404"

print_info "Testing start test with invalid test ID"
test_endpoint "POST" "/api/v1/tests/start?officerID=1&testID=999" "Start test with invalid test ID" "404"

print_info "Testing start test with invalid parameter format"
test_endpoint "POST" "/api/v1/tests/start?officerID=abc&testID=123" "Start test with invalid parameter format" "400"

print_info "Testing start test with missing parameters"
test_endpoint "POST" "/api/v1/tests/start" "Start test with missing parameters" "400"

# 8. Test Complete Flow
print_header "8. COMPLETE FLOW TEST"

print_info "Testing complete flow: Get officers -> Get units -> Get subjects -> Create test -> Start test"

echo -e "\n${YELLOW}Step 1: Get all officers${NC}"
officers_response=$(curl -s "$BASE_URL/api/v1/officers")
officer_count=$(echo "$officers_response" | jq '.count' 2>/dev/null)
print_success "Found $officer_count officers"

echo -e "\n${YELLOW}Step 2: Get all units${NC}"
units_response=$(curl -s "$BASE_URL/api/v1/units")
unit_count=$(echo "$units_response" | jq '.count' 2>/dev/null)
print_success "Found $unit_count units"

echo -e "\n${YELLOW}Step 3: Get all subjects${NC}"
subjects_response=$(curl -s "$BASE_URL/api/v1/subjects")
subject_count=$(echo "$subjects_response" | jq '.count' 2>/dev/null)
print_success "Found $subject_count subjects"

echo -e "\n${YELLOW}Step 4: Create test for officer 3, subject 3${NC}"
new_test_response=$(curl -s "$BASE_URL/api/v1/tests/officer-subject?officerID=3&subjectID=3")
new_test_id=$(echo "$new_test_response" | jq -r '.id' 2>/dev/null)
officer_name=$(echo "$new_test_response" | jq -r '.officer.name' 2>/dev/null)
subject_name=$(echo "$new_test_response" | jq -r '.subject.name' 2>/dev/null)

if [ "$new_test_id" != "null" ] && [ "$new_test_id" != "" ]; then
    print_success "Created test ID: $new_test_id for $officer_name in subject: $subject_name"
    
    echo -e "\n${YELLOW}Step 5: Start the test${NC}"
    start_response=$(curl -s -X POST "$BASE_URL/api/v1/tests/start?officerID=3&testID=$new_test_id")
    start_time=$(echo "$start_response" | jq -r '.start_time' 2>/dev/null)
    
    if [ "$start_time" != "null" ] && [ "$start_time" != "0" ]; then
        print_success "Test started successfully at timestamp: $start_time"
    else
        print_error "Failed to start test"
    fi
else
    print_error "Failed to create test for complete flow"
fi

# 7. Performance Test
print_header "7. QUICK PERFORMANCE TEST"

print_info "Running 5 concurrent requests to test API responsiveness..."

for i in {1..5}; do
    (curl -s "$BASE_URL/health" > /dev/null) &
done
wait

print_success "All concurrent health checks completed"

# Final Summary
print_header "TEST SUMMARY"
echo -e "${GREEN}✅ All API endpoints have been tested${NC}"
echo -e "${BLUE}📚 Available endpoints:${NC}"
echo "  - GET  /health"
echo "  - GET  /api/v1/units"
echo "  - GET  /api/v1/officers"
echo "  - GET  /api/v1/tests/officer-subject"
echo "  - POST /api/v1/tests/start"
echo ""
echo -e "${BLUE}📖 Documentation available at:${NC}"
echo "  - Swagger: http://localhost:8080/v1/api/free-contest/swagger/index.html"
echo ""
echo -e "${YELLOW}💡 Usage examples:${NC}"
echo "  curl \"$BASE_URL/api/v1/units\""
echo "  curl \"$BASE_URL/api/v1/officers\""
echo "  curl \"$BASE_URL/api/v1/tests/officer-subject?officerID=1&subjectID=1\""
echo "  curl -X POST \"$BASE_URL/api/v1/tests/start?officerID=1&testID=123\""

print_header "TESTING COMPLETED"
