package repo

import (
	//"PTOBuilder/internal/auth"
	packageUser "PTOBuilder/internal/auth"
	"PTOBuilder/internal/auth/model"
	"PTOBuilder/pkg/logging"
	"PTOBuilder/pkg/storage"
	"PTOBuilder/pkg/utils"
	"context"
	"fmt"
)

type repo struct {
	log    *logging.Logger
	client storage.Client
}

func NewRepo(log *logging.Logger, client storage.Client) packageUser.Repo {
	return &repo{
		log:    log,
		client: client,
	}
}

func (r *repo) AddUser(ctx context.Context, user *model.User) error {
	q := `insert into users (email, name, password)
    		  values ($1, $2, $3)
    		  returning users.id, users.role;`
	q = utils.FormatQuery(q)
	var id int64
	err := r.client.QueryRow(ctx, q, user.Email, user.Name, user.Password).Scan(&id, &user.Role)
	if err != nil {
		r.log.Info(err)
		return err
	}
	return nil
}

func (r *repo) GetUser(ctx context.Context, user *model.User) error {
	var id int64
	q := `select users.id, users.name, users.password, users.role from users
    		  where users.email = $1;`
	q = utils.FormatQuery(q)
	err := r.client.QueryRow(ctx, q, user.Email).Scan(&id, &user.Name, &user.Password, &user.Role)
	fmt.Println(id)
	if id == 0 {
		return fmt.Errorf("user does not exist")
	}
	return err
}
