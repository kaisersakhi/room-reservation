# echo "Cleaning cache....# go clean -cache
echo "Building binary..."
go build -o ./tmp/main ./cmd/room-reservation
echo "Attaching debugger..."

# dlv debug ./cmd/room-reservation/*.go --headless --listen=:2345 --api-version=2 --log

go build -o ./tmp/main ./cmd/room-reservation && dlv exec ./tmp/main --headless --listen=:2345 --api-version=2 --log

# air