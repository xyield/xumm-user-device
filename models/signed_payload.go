package models

type SignPayload struct {
	SignedBlob  string     `json:"signed_blob"`
	TxID        string     `json:"tx_id"`
	Multisigned string     `json:"multisigned,omitempty"`
	Dispatched  Dispatched `json:"dispatched"`
	Permission  Permission `json:"permission"`
}
type Dispatched struct {
	To     string `json:"to"`
	Result string `json:"result"`
}
type Permission struct {
	Push bool `json:"push"`
	Days int  `json:"days"`
}
