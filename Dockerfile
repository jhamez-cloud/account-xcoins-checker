FROM cleanstart/go:latest
COPY . /app
WORKDIR /app
CMD ["go", "run", "cmd/api/main.go"]