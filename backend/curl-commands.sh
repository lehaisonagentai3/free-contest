#!/bin/bash

# Free Contest API - All cURL Commands
# Copy and paste these commands to test the API

echo "🚀 Free Contest API - cURL Commands"
echo "====================================="
echo ""

echo "📋 Copy these commands to test the API:"
echo ""

echo "# 1. Health Check"
echo "curl \"http://localhost:8080/health\""
echo ""

echo "# 2. Get All Units"
echo "curl \"http://localhost:8080/api/v1/units\""
echo ""

echo "# 3. Get All Officers"
echo "curl \"http://localhost:8080/api/v1/officers\""
echo ""

echo "# 4. Get Officer by ID"
echo "curl \"http://localhost:8080/api/v1/officers/1\""
echo "curl \"http://localhost:8080/api/v1/officers/2\""
echo "curl \"http://localhost:8080/api/v1/officers/3\""
echo ""

echo "# 5. Get All Subjects"
echo "curl \"http://localhost:8080/api/v1/subjects\""
echo ""

echo "# 6. Get Test for Officer (creates a test with random questions)"
echo "curl \"http://localhost:8080/api/v1/tests/officer-subject?officerID=1&subjectID=1\""
echo "curl \"http://localhost:8080/api/v1/tests/officer-subject?officerID=2&subjectID=2\""
echo "curl \"http://localhost:8080/api/v1/tests/officer-subject?officerID=3&subjectID=3\""
echo ""

echo "# 5. Start Test (use actual test ID from step 4)"
echo "curl -X POST \"http://localhost:8080/api/v1/tests/start?officerID=1&testID=YOUR_TEST_ID\""
echo ""

echo "# 6. Submit Test (use actual test ID from step 4, requires JSON body with answers)"
echo "curl -X POST \"http://localhost:8080/api/v1/tests/submit?officerID=1&testID=YOUR_TEST_ID\" \\"
echo "  -H \"Content-Type: application/json\" \\"
echo "  -d '{\"1\":\"A\",\"2\":\"B\",\"3\":\"C\",\"4\":\"A\",\"5\":\"B\"}'"
echo ""

echo "# 7. Error Examples"
echo "curl \"http://localhost:8080/api/v1/tests/officer-subject?officerID=999&subjectID=1\"  # Invalid officer"
echo "curl \"http://localhost:8080/api/v1/tests/officer-subject?officerID=1&subjectID=999\"  # Invalid subject"
echo "curl \"http://localhost:8080/api/v1/tests/officer-subject\"  # Missing parameters"
echo ""

echo "# 7. Pretty JSON Output (with jq)"
echo "curl \"http://localhost:8080/api/v1/units\" | jq ."
echo "curl \"http://localhost:8080/api/v1/officers\" | jq ."
echo ""

echo "# 8. Complete Flow Example"
echo "# Step 1: Get officers list"
echo "curl \"http://localhost:8080/api/v1/officers\" | jq '.data[0]'"
echo ""
echo "# Step 2: Create test for first officer"
echo "TEST_ID=\$(curl -s \"http://localhost:8080/api/v1/tests/officer-subject?officerID=1&subjectID=1\" | jq -r '.id')"
echo ""
echo "# Step 3: Start the test"
echo "curl -X POST \"http://localhost:8080/api/v1/tests/start?officerID=1&testID=\$TEST_ID\" | jq ."
echo ""
echo "# Step 4: Submit answers"
echo "curl -X POST \"http://localhost:8080/api/v1/tests/submit?officerID=1&testID=\$TEST_ID\" \\"
echo "  -H \"Content-Type: application/json\" \\"
echo "  -d '{\"1\":\"A\",\"2\":\"B\",\"3\":\"C\"}' | jq ."
echo ""

echo "📖 Documentation:"
echo "Swagger UI: http://localhost:8080/v1/api/free-contest/swagger/index.html"
echo ""

echo "📊 Available Data:"
echo "- Officers: 3 officers (IDs: 1, 2, 3)"
echo "- Units: 9 units (IDs: 1-9)"
echo "- Subjects: 3 subjects (IDs: 1, 2, 3)"
echo "  - Subject 1: 'Đề thi HCKT' (60 minutes, 13 questions)"
echo "  - Subject 2: 'Đề thi chính trị' (60 minutes, 15 questions)"
echo "  - Subject 3: 'Đề thi tham mưu' (120 minutes, 15 questions)"
