package assets

import (
	_ "embed"
	"encoding/base64"
)

//go:embed platnova.png
var Logo []byte
var Base64Logo string

func init() {
	Base64Logo = base64.StdEncoding.EncodeToString(Logo)
}
