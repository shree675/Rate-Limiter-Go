# Rate Limiter

A simple rate limiter written in Go following the Leaky Bucket algorithm loosely.  
Processes requests at a constant rate while also allowing a surge in requests to some extent.

## Usage

```bash
go run .
```

This will start the HTTP server at 8080.

```bash
./test.sh
```

This is a sample script to simulate concurrent requests to see the rate limiting in action.
