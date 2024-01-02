package modelos

// para redefinir senha
type Senha struct {
	Nova  string `json:"nova"`
	Atual string `json:"atual"`
}
