package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/MicahParks/keyfunc"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func main() {
	r := gin.Default()

	r.GET("/api/userinfo", func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
			return
		}

		// Ensure it's a Bearer token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
			return
		}

		// Parse and validate JWT
		claims, err := parseAndValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("invalid token: %v", err)})
			return
		}

		// Return the full claims JSON
		c.JSON(http.StatusOK, claims)
	})

	// Start the server on port 8080
	r.Run(":8080")
}

// parseAndValidateJWT extracts the issuer from the JWT and validates it using the correct JWKS URL
func parseAndValidateJWT(tokenString string) (map[string]interface{}, error) {
	// Extract issuer (iss) from token without verification
	issuer, err := extractIssuer(tokenString)
	if err != nil {
		return nil, fmt.Errorf("failed to extract issuer: %v", err)
	}

	// Construct JWKS URL
	jwksURL := fmt.Sprintf("%s/protocol/openid-connect/certs", issuer)

	// Fetch JWKS (NO caching, always fresh request)
	jwks, err := fetchJWKS(jwksURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch JWKS: %v", err)
	}

	// Validate the JWT using the correct JWKS
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, jwks.Keyfunc)
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	// Extract claims and return as a map
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid claims format")
}

// extractIssuer extracts the "iss" (issuer) field from the JWT without verifying it
func extractIssuer(tokenString string) (string, error) {
	parts := strings.Split(tokenString, ".")
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid JWT format")
	}

	// Decode the payload (second part of JWT)
	payloadJSON, err := decodeBase64URL(parts[1])
	if err != nil {
		return "", fmt.Errorf("failed to decode JWT payload: %v", err)
	}

	// Parse JSON to extract "iss" field
	var payload map[string]interface{}
	if err := json.Unmarshal(payloadJSON, &payload); err != nil {
		return "", fmt.Errorf("failed to parse JWT payload: %v", err)
	}
	issuer, ok := payload["iss"].(string)
	if !ok {
		return "", fmt.Errorf("issuer (iss) claim not found in token")
	}

	return issuer, nil
}

// fetchJWKS fetches JWKS from the issuer without caching
func fetchJWKS(url string) (*keyfunc.JWKS, error) {
	return keyfunc.Get(url, keyfunc.Options{})
}

// decodeBase64URL decodes a base64 URL-encoded string
func decodeBase64URL(input string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(input)
}
