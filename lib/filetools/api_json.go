package filetools

type ApiJSON struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Database   string `json:"database"`
	PublicKey  string `json:"pubkey"`
	PrivateKey string `json:"privkey"`
}
