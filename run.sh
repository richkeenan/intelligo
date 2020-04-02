export PROJECT_DIR=/Users/richardkeenan/countingup/svcs/accountants
export SERVICE_PRIVATE_KEY_FILE=".mock-secrets/accountants_private_key"
export AUTH_PUBLIC_KEYS_FILE=".mock-secrets/auth_public_keys"
export AUTH_HOST="localhost:8070"
export DB_CONNECT_STRING="root:password@tcp(localhost:3306)/accountants"
export USER_ADDR="http://localhost:8013"
export SQS_ADDR="http://localhost:9494"
export COGNITO_ADDR="http://localhost:8077"
export KINESIS_ADDR="http://localhost:4567"
export SEGMENT_ACCOUNTANTS_WRITE_KEY="accountants"
export SEGMENT_WRITE_KEY="clients"
export SEGMENT_MOCK_ADDR="http://localhost:17001"
export COMMS_ADDR="http://localhost:8010"
export EMAIL_TRACK_ADDR="http://localhost:8080"

go run $PROJECT_DIR/main.go