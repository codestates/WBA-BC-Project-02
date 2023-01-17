package contract

import "encoding/json"

type Nonce struct {
	NonceValue uint64 `json:"nonce_exists"`
}

func (l *Nonce) MarshalBinary() (data []byte, err error) {
	return json.Marshal(l)
}
