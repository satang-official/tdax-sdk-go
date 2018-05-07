package signature

import (
	"encoding/hex"

	"bitbucket.org/satangcorp/xsig"
)

func Sign(secret string, payload map[string]interface{}) string {
	sig := xsig.Sign([]byte(secret), payload)
	dst := make([]byte, hex.EncodedLen(len(sig)))
	hex.Encode(dst, sig)
	return string(dst)
}

// "Market":      opt.Market,
// "Nonce":       body.Nonce,
// "Price":       body.Price,
// "Qty":         body.Qty,
// "Side":        "SELL",
// "Symbol":      body.Symbol,
// "Type":        "LIMIT",
