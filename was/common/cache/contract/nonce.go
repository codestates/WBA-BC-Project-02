package contract

import "encoding/json"

type Nonce struct {
	NonceValues []uint64 `json:"nonce_values"`
}

func (l *Nonce) MarshalBinary() (data []byte, err error) {
	return json.Marshal(l)
}
