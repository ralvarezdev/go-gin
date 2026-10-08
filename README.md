# go-gin

Gin helpers for Go projects: a JWT bearer-token authentication middleware, a Redis-backed rate-limiting middleware adapter, and a standard response/error handler. Requires Go 1.25 (per `go.mod`).

## Installation

```bash
go get github.com/ralvarezdev/go-gin
```

Direct dependencies: `gin`, and `go-flags`, `go-jwt`, `go-rate-limiter` from `github.com/ralvarezdev`.

## Packages

- **`gogin`** (root) — `AuthorizationHeaderKey` (`"Authorization"`) and shared errors such as `ErrInvalidAuthorizationHeader`.
- **`middleware/auth`** — `Authenticator` interface and `NewMiddleware(validator, responseHandler, jwtValidatorErrorHandler)`. Reads `Authorization: Bearer <token>`, validates the claims through a `go-jwt` validator, stores claims and raw token in the Gin context, and responds 401 on a malformed header.
- **`middleware/ratelimiter`** — `RateLimiter` interface (`Limit() gin.HandlerFunc`).
- **`middleware/ratelimiter/redis`** — `NewMiddleware(rateLimiter)` wrapping a `go-rate-limiter` Redis limiter.
- **`response`** — `Handler` interface (`HandleSuccess`, `HandleErrorProne`, `HandleError`), `NewDefaultHandler(mode)`, and `NewResponse*` / `NewErrorResponse*` / `NewJSONErrorResponse*` / `SendInternalServerError` helpers.
- **`jwt/validator`** — `DefaultErrorHandler(ctx, err)` translating JWT validation errors into HTTP responses.

## Usage

```go
modeFlag := goflagsmode.NewFlag(goflagsmode.DefaultMode, goflagsmode.AllowedModes)
respHandler, err := goginresponse.NewDefaultHandler(modeFlag)

// jwtValidator is a go-jwt validator
authMW, err := goginauth.NewMiddleware(
    jwtValidator, respHandler, goginjwtvalidator.DefaultErrorHandler,
)

r := gin.New()
r.GET("/me", authMW.Authenticate(accessTokenKind), func(c *gin.Context) {
    respHandler.HandleSuccess(c, goginresponse.NewResponseWithCode(gin.H{"ok": true}, 200))
})
```

`accessTokenKind` is a `go-jwt` `token.Token` value; see that repository for the available kinds.

## Development

```bash
go build ./...
go vet ./...
```

There are no tests.

## License

GNU General Public License v3.0 (see [LICENSE](LICENSE)).
