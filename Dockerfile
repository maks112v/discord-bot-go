FROM golang:1.23-alpine

WORKDIR /app

# Copy only dependency files first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./main.go

# Use full path and exec form
CMD ["/app/main"]
