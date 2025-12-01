package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"
)

func (m *Middlewares) AuthenticateJWT(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		if header == "" {
			http.Error(w, "Unathorized", http.StatusUnauthorized)
			return
		}

		headerArr := strings.Split(header, " ")

		if len(headerArr) != 2 {
			http.Error(w, "Unathorized", http.StatusUnauthorized)
			return
		}

		accesstoken := headerArr[1]

		tokenParts := strings.Split(accesstoken, ".")

		if len(tokenParts) != 3 {
			http.Error(w, "Unathorized", http.StatusUnauthorized)
			return
		}

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		signature := tokenParts[2]

		message := jwtHeader + "." + jwtPayload

		byteArrSecret := []byte(m.cnf.JwtSecretKey)

		byteArrMessage := []byte(message)

		h := hmac.New(sha256.New, byteArrSecret)
		h.Write(byteArrMessage)

		hash := h.Sum(nil)
		newSignature := base64UrlEencode(hash)

		if newSignature != signature {
			http.Error(w, "Unathorized. Hackerrrr,", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})

}
func base64UrlEencode(data []byte) string {
	return base64.URLEncoding.Strict().WithPadding(base64.NoPadding).EncodeToString(data)
}
