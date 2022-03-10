package models

type PatchResponse struct {
	ReferenceCallUuidv4 string      `json:"reference_call_uuidv4"`
	Signed              bool        `json:"signed"`
	UserToken           interface{} `json:"user_token"`
	ReturnURL           ReturnUrl   `json:"return_url"`
}

type ReturnUrl struct {
	App string `json:"app"`
	Web string `json:"web"`
}
