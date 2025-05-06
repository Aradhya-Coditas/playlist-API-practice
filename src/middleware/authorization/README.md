# authorization middleware

This middleware functions written in bff for handling JWT authorization in a backend application.

---

### AuthorizeJWTToken

**Purpose:** This middleware function is responsible for authorizing JWT access tokens.

**Functionality:**
1. Extracts the JWT token from the request header.
2. Validates the presence of the JWT token and its format.
3. Retrieves the secret key for JWT verification from application configuration.
4. Parses the JWT token and verifies its authenticity.
5. Checks the token type and expiration time.
6. Decrypts the token payload and extracts user-related information.
7. Sets relevant user information in the Gin context for downstream handlers.

---

### AuthorizeRefreshJWTToken

**Purpose:** This middleware function is responsible for authorizing JWT refresh tokens.

**Functionality:**
1. Extracts the JWT token from the request header.
2. Validates the presence of the JWT token.
3. Retrieves the secret key for JWT verification from application configuration.
4. Parses the JWT token and verifies its authenticity.
5. Checks the token type and expiration time.
6. Extracts token payload and unmarshals it into a structured format.
7. Sets relevant user information in the Gin context for downstream handlers.

---

## Usage

To use these middleware functions in your Gin application, follow these steps:

1. Import the necessary packages:

```
import (
    "github.com/gin-gonic/gin"
    "omnenest-backend/src/utils/authorization"
)
```

2. Initialize the JWT utilities:

```
jwtUtils := authorization.NewJwtTokenUtils()
```

3. Add middleware to your Gin router:

```
router := gin.Default()

// Example route with JWT token authorization
router.use(authorization.AuthorizeJWTToken(jwtUtils))

// Example route with JWT refresh token authorization
router.GET("/refresh", authorization.AuthorizeRefreshJWTToken(jwtUtils), func(ctx *gin.Context) {
    // Your refresh token route logic here
})

router.Run(":8080")

```