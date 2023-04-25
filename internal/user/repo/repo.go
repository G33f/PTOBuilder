package repo

import (
	//"PTOBuilder/internal/user"
	packageUser "PTOBuilder/internal/user"
	"PTOBuilder/internal/user/model"
	"PTOBuilder/pkg/logging"
	"PTOBuilder/pkg/storage"
	"PTOBuilder/pkg/utils"
	"context"
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

func (r *repo) UserRegistration(ctx context.Context, user *model.User) error {
	q := `insert into users (email, name, password)
    		  values ($1, $2, $3)
    		  returning users.id;`
	q = utils.FormatQuery(q)
	var id int64
	err := r.client.QueryRow(ctx, q, user.Email, user.Name, user.Password).Scan(&id)
	if err != nil {
		r.log.Info(err)
		return err
	}
	return nil
}

func (r *repo) GetUser(ctx context.Context, user *model.User) error {
	q := `select users.name, users.password from users
    		  where users.email = $1;`
	q = utils.FormatQuery(q)
	err := r.client.QueryRow(ctx, q, user.Email).Scan(&user.Name, &user.Password)
	return err
}
