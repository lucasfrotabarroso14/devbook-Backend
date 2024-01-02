package modelos

import (
	"errors"
	"strings"
	"time"
)

type Publicacao struct {
	ID        uint64    `json:"id, omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autorId,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadaEm  time.Time `json:"criadaEm,omitempty"`
}

func (publicao *Publicacao) Preparar() error {
	if erro := publicao.validar(); erro != nil {
		return erro
	}
	publicao.formatar()
	return nil

}

func (publicao *Publicacao) validar() error {
	if publicao.Titulo == "" {
		return errors.New("O título é obrigatório e não pdoe estar em branco")
	}

	if publicao.Conteudo == "" {
		return errors.New("O conteúdo é obrigatorio e não pode estar em branco")
	}
	return nil
}

func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
