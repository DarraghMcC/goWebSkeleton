## Web app skeleton readme

## Description
Basic golang skeleton for a web app. Includes:
- Basic ping endpoint
- Metrics 
- Unit tests


**End points**

GET `/ping` should return a 200 and a string of "pong" when server up

GET `/metrics` returns prometheus metrics

## Run Set up

**Optional env vars:**

* `PORT` (defaults to 8089)
* `DISABLE_LOG_TIMESTAMP` (defaults to false)


## How to run
**Command**

`go run .`

**Port**

`8090` or value of `PORT`

## Testing
**Unit Tests:**
`go test -v`
