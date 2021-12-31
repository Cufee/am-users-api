export PATH=$PATH:$(go env GOPATH)/bin
# Env
export PORT=3001
export MONGO_URI=mongodb://localhost
export MONGO_PORT=27017
export MONGO_USER=root
export MONGO_PASS=root

# Free up port
kill -9 $(lsof -ti:$PORT)
# Run with hot-reload
air