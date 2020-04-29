package encode

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
)

func CalcSign(timestamp int64, signToken string) string {
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, signToken)
	h := hmac.New(sha256.New, []byte(signToken))
	h.Write([]byte(stringToSign))
	sumValue := h.Sum(nil)
	return url.QueryEscape(base64.StdEncoding.EncodeToString(sumValue))
}
