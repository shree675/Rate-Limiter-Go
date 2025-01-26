# Rate Limiter

A simple rate limiter written in Go following the Leaky Bucket algorithm loosely.  
Processes requests at a constant rate while also allowing a surge in requests to some extent.

## Usage

Start the HTTP server at 8080:
```bash
go run .
```

Run a sample script to simulate concurrent requests to see the rate limiting in action:
```bash
./test.sh
```
