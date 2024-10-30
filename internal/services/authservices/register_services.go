package authservices

import (
	"app-fullstack-gotth/internal/database/db"
	"app-fullstack-gotth/structs"
	"context"
	"time"

	"github.com/google/uuid"
)

type RegisterServices struct {
	repo *db.Queries
}

func NewRegisterServices(repo *db.Queries) *RegisterServices {
	return &RegisterServices{repo: repo}
}

func (ls *RegisterServices) Register(args structs.AuthRegisterFormRequest) error {
	user := db.CreateOneUserParams{
		ID:        uuid.NewString(),
		FirstName: args.FirstName,
		LastName:  args.LastName,
		Email:     args.Email,
		Password:  args.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := ls.repo.CreateOneUser(context.TODO(), user); err != nil {
		return err
	}

	return nil
}
