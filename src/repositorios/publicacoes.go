package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO publicacoes (titulo, conteudo, autor_id) values (?,?,?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}
	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil

}

func (repositorio Publicacoes) BuscaPorID(publicacaoID uint64) (modelos.Publicacao, error) {
	linha, erro := repositorio.db.Query(`
		SELECT p.*, u.nick from PUBLICACOES p INNER JOIN usuarios u 
		on u.id = p.autor_id WHERE p.id=?
`, publicacaoID) // o p* quer dizer que quero todas as colunas da tabela de publicacoes
	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	defer linha.Close()
	var publicacao modelos.Publicacao
	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}
	return publicacao, nil

}

func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
	SELECT DISTINCT p.*, u.nick FROM PUBLICACOES p
    INNER JOIN usuarios u on u.id = p.autor_id
    INNER JOIN seguidores s ON p.autor_id = s.usuario_id
    where u.id = ? or s.seguidor_id = ?
    order by 1 desc`,
		usuarioID, usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()
	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil

}

func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao modelos.Publicacao) error {

	statement, erro := repositorio.db.Prepare("UPDATE PUBLICACOES SET titulo = ?, conteudo = ?  WHERE id=?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); erro != nil {
		return erro
	}
	return nil
}

func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from publicacoes where id = ? ")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro

	}
	return nil
}

func (repositorio Publicacoes) BuscaPorUsuario(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
	select p.*, u.nick from publicacoes p 
	join usuarios u on u.id = p.autor_id
	where p.autor_id = ?
`, usuarioID)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()
	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}

func (repositorio Publicacoes) Curtir(PublicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare("update publicacoes set curtidas = curtidas + 1 where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(PublicacaoID); erro != nil {
		return erro
	}
	return nil

}

func (repositorio Publicacoes) Descurtir(PublicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(
		`update publicacoes set curtidas = 
				CASE
				 	WHEN curtidas >0 THEN curtidas - 1 
					ELSE 0
				    END
					WHERE id = ?`)
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(PublicacaoID); erro != nil {
		return erro
	}
	return nil

}
