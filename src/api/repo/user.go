package repo

import (
	"context"

	"go-clean-arch/src/domain"
	Error "go-clean-arch/src/pkg/error"
)

func (m *Repo) Login(ctx context.Context, b domain.Login) (data domain.UserLogin, err error) {
	query := `SELECT u.id, u.uuid, u.email, u.password, u.role_id 	
	FROM users u 	
	where u.email = ?`
	err = m.Conn.QueryRow(query, b.Email).Scan(&data.ID, &data.UUID, &data.Email, &data.Password,
		&data.RoleId)

	if err != nil {
		Error.Error(err)
		return data, err
	}
	return
}

func (m *Repo) GetUserByUuid(ctx context.Context, uid string) (res domain.User, err error) {
	query := `SELECT id, uuid, email FROM users WHERE uuid = ?`
	err = m.Conn.QueryRow(query, uid).Scan(&res.ID, &res.UUID, &res.Email)
	if err != nil {
		return
	}
	return
}
