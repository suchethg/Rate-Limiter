# Rate-Limiter

# Per-Client Rate Limiting

This project implements per-client rate limiting to control how many requests each client can make within a given time window. The goal is to protect backend services, prevent abuse, and ensure fair usage across clients.

# Token Bucket Algorithm

How it works:

Each client has a bucket of tokens

Tokens are added to the bucket at a fixed rate (e.g., 10 tokens per second)

Each request consumes 1 token

If no tokens are available → request is rejected (429 Too Many Requests)

Key Properties:

Allows short bursts of traffic

Enforces an average request rate

More flexible than fixed window limits

Example:
Bucket size: 100 tokens
Refill rate: 10 tokens/sec

Client sends:
- 50 requests instantly → allowed
- 60th request immediately → rejected
- After 1 second → 10 tokens refilled

# Tollbooth Analogy

Think of rate limiting like a highway tollbooth 🚗

Each car = a request

Each client has a wallet of tokens

Passing the toll consumes a token

Tokens refill over time

No tokens → car must wait

This ensures:

No traffic jams

Fair access for all drivers

Predictable flow

