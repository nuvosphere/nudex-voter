package http

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestFireblocksCosignerTxSignJwt(t *testing.T) {
	// Create a valid JWT token
	bodyBytes := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJ0eElkIjoiZDEwYzE0NTgtZjM1Mi00NTg4LWJmMDgtMTMzYzIzYzZkZDA4Iiwib3BlcmF0aW9uIjoiVFJBTlNGRVIiLCJzb3VyY2VUeXBlIjoiVkFVTFQiLCJzb3VyY2VJZCI6IjEiLCJkZXN0VHlwZSI6Ik9ORV9USU1FX0FERFJFU1MiLCJkZXN0SWQiOiIiLCJhc3NldCI6IkVUSF9URVNUNSIsImFtb3VudCI6MC4wMDEwMDAwMCwiYW1vdW50U3RyIjoiMC4wMDEiLCJyZXF1ZXN0ZWRBbW91bnQiOjAuMDAxMDAwMDAsInJlcXVlc3RlZEFtb3VudFN0ciI6IjAuMDAxIiwiZGVzdEFkZHJlc3MiOiIweDI4QTMxMzQzZDdkNzAwYzRDQTUwYkJFMjIzNTU1ZjU3YzlhOTc1MDYiLCJleHRyYVBhcmFtZXRlcnMiOnsic2VjdXJpdHlFbnJpY2htZW50Ijp7fX0sImRlc3RpbmF0aW9ucyI6W3siYW1vdW50TmF0aXZlIjowLjAwMTAwMDAwLCJhbW91bnROYXRpdmVTdHIiOiIwLjAwMSIsImFtb3VudFVTRCI6Mi4zNzIwMzQ4NiwiZHN0QWRkcmVzcyI6IjB4MjhBMzEzNDNkN2Q3MDBjNENBNTBiQkUyMjM1NTVmNTdjOWE5NzUwNiIsImRzdEFkZHJlc3NUeXBlIjoiT05FX1RJTUUiLCJkc3RJZCI6IiIsImRzdFRhZyI6ImxldCdzIGdvb28iLCJkc3RUeXBlIjoiT05FX1RJTUVfQUREUkVTUyIsImRpc3BsYXlEc3RBZGRyZXNzIjoiMHgyOEEzMTM0M2Q3ZDcwMGM0Q0E1MGJCRTIyMzU1NWY1N2M5YTk3NTA2IiwiZGlzcGxheURzdFRhZyI6ImxldCdzIGdvb28iLCJhY3Rpb24iOiIyLVRJRVIiLCJhY3Rpb25JbmZvIjp7ImNhcHR1cmVkUnVsZU51bSI6MCwicnVsZXNTbmFwc2hvdElkIjo0NzE1NiwiYnlHbG9iYWxQb2xpY3kiOmZhbHNlLCJieVJ1bGUiOnRydWUsInJ1bGVUeXBlIjoiVEVOQU5UIiwiY2FwdHVyZWRSdWxlIjoie1wiZHN0XCI6e1wiaWRzXCI6W1tcIipcIl1dfSxcInNyY1wiOntcImlkc1wiOltbXCIxXCIsXCJWQVVMVFwiLFwiKlwiXV19LFwidHlwZVwiOlwiVFJBTlNGRVJcIixcImFzc2V0XCI6XCIqXCIsXCJhY3Rpb25cIjpcIjItVElFUlwiLFwiYW1vdW50XCI6MCxcIm9wZXJhdG9yc1wiOntcIndpbGRjYXJkXCI6XCIqXCJ9LFwicGVyaW9kU2VjXCI6MCxcImFtb3VudFNjb3BlXCI6XCJTSU5HTEVfVFhcIixcImFtb3VudEN1cnJlbmN5XCI6XCJVU0RcIixcImRzdEFkZHJlc3NUeXBlXCI6XCJPTkVfVElNRVwiLFwiYXBwbHlGb3JBcHByb3ZlXCI6ZmFsc2UsXCJ0cmFuc2FjdGlvblR5cGVcIjpcIlRSQU5TRkVSXCIsXCJhbGxvd2VkQXNzZXRUeXBlc1wiOlwiRlVOR0lCTEVcIixcImV4dGVybmFsRGVzY3JpcHRvclwiOlwie1xcXCJpZFxcXCI6XFxcImQ0MDMwZmQ5LWZiY2UtNGE3Ny05NjNjLWQwZTIwNTljZmZkMVxcXCJ9XCIsXCJhdXRob3JpemF0aW9uR3JvdXBzXCI6e1wibG9naWNcIjpcIk9SXCIsXCJncm91cHNcIjpbe1widGhcIjoyLFwidXNlcnNcIjpbXCJhMjQzMjlhYy1mYjVhLTQ3ZGQtYmJjMC1mMTI1M2I4MzUxNDhcIl0sXCJ1c2Vyc0dyb3Vwc1wiOltcIjE3YzhmYTY5LTQ4MDUtNDlhMS04YTE1LTFlOTdmYzA1YTUwMlwiXX1dLFwiYWxsb3dPcGVyYXRvckFzQXV0aG9yaXplclwiOnRydWV9fSJ9LCJhdXRob3JpemF0aW9uR3JvdXBzIjp7ImxvZ2ljIjoiT1IiLCJncm91cHMiOlt7InRoIjoyLCJ1c2VycyI6WyJhMjQzMjlhYy1mYjVhLTQ3ZGQtYmJjMC1mMTI1M2I4MzUxNDgiXSwidXNlcnNHcm91cHMiOlsiMTdjOGZhNjktNDgwNS00OWExLThhMTUtMWU5N2ZjMDVhNTAyIl19XSwiYWxsb3dPcGVyYXRvckFzQXV0aG9yaXplciI6dHJ1ZX19XSwibm90ZSI6IkdPQVQgTDIgU0QhIiwicmVxdWVzdElkIjoiZDEwYzE0NTgtZjM1Mi00NTg4LWJmMDgtMTMzYzIzYzZkZDA4Iiwic2lnbmVySWQiOiJiMzJkY2VmNC02OWU4LTQyZTctYWFmYi0yZjU4ZjkyZDViZmUifQ.aNQw-Kzj4wh3-7aO01li09QxEyuT0-EitPJOJFx2sJQSYaPq29FdoRN1eFQKhZlJTVDVyovJvghuUH87thyM8MUEgZL6v9g9XUefCVANyE8ZaY1u1ADyM0QUzcuevQonE_-986zry7RAzQqzgiGM_4vv9-9HEAdHNrYV4nL88olBjghqtP5aaaL9Qajtalm2m_7YLOVFvuGh02xN4NYAPpmLREwZGDMb5anBXY93rCEmSGsfQPMf44-qmSjBMTlHh0-vuPFXBiHD69I6EScx-R4c6VFUk3Aqx8vML2rIMl8smwQ4v4byTfXfJ9hMwFG0kUYLJGZ2oZkVznKS44sFmQ"
	cosignerPubKey := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsp4b/tDJfke0LtWoQmBn
HEPevpL7hR5QwbwmqRSuQN8eU/nCWZRi5K7gy09fwgPQxaGy8868+VqkJBXRuENS
5GizTWArX/NeO0ERdPzxGU9YUqWjC5wmz4il9AjEkhZ32kEEEt8FiY5SwCvaa2a0
8h7VVD2AKi+IlVwHhzOm9VXwTNbckO+JFilWET+/6SoQnSnWTld458jYKMD3tTTs
obwpt4QjfL+lNlfjkbrqhLHWA1pUKLfFaDG4LdmtOe9uC7lcwnO4kiyLI2ppC0mZ
mzDC2REwx8Kq4c5UnhwnSaUUQ/3JEnbGe0EPCImmb9TyFXPszJzHolC5Yf0nz2yh
BQIDAQAB
-----END PUBLIC KEY-----`
	rawBody := string(bodyBytes)

	block, _ := pem.Decode([]byte(cosignerPubKey))
	if block == nil || block.Type != "PUBLIC KEY" {
		t.Errorf("failed to parse PEM block containing the public key")
		return
	}

	parsedKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		t.Errorf("failed to parse public key: %v", err)
		return
	}

	pubKey, ok := parsedKey.(*rsa.PublicKey)
	if !ok {
		t.Error("not an RSA public key")
		return
	}

	// Parse and verify JWT
	tx, err := jwt.Parse(rawBody, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, jwt.ErrInvalidKey
		}
		return pubKey, nil
	})
	assert.NoError(t, err)
	assert.True(t, tx.Valid)

	// Extract requestId from the JWT claims
	claims, ok := tx.Claims.(jwt.MapClaims)
	assert.True(t, ok)

	requestId := claims["requestId"].(string)
	txId := claims["txId"].(string)

	t.Logf("Cosigner callback JWT claim received, requestId %s, txId %s", requestId, txId)
}