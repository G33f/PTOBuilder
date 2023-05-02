package usecase

import (
	"PTOBuilder/internal/auth"
	"PTOBuilder/internal/auth/model"
	"PTOBuilder/pkg/logging"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type useCase struct {
	log  *logging.Logger
	repo auth.Repo

	signingKey []byte
	tokens     map[string]*jwt.Token
}

func NewUseCase(log *logging.Logger, repo auth.Repo, signingKey []byte) auth.UseCase {
	return &useCase{
		log:        log,
		repo:       repo,
		signingKey: signingKey,
		tokens:     map[string]*jwt.Token{},
	}
}

func (u *useCase) SignUp(ctx context.Context, user *model.User) error {
	if check := u.isUserExist(ctx, user); check == true {
		return fmt.Errorf("auth already exist")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		u.log.Info(err)
		return err
	}
	user.Password = string(hash)
	err = u.repo.AddUser(ctx, user)
	if err != nil {
		u.log.Info(err)
		return err
	}
	return nil
}

func (u *useCase) isUserExist(ctx context.Context, user *model.User) bool {
	err := u.repo.GetUser(ctx, user)
	if err == nil {
		return false
	}
	return true
}

func (u *useCase) SignIn(ctx context.Context, user *model.User) (string, error) {
	user1 := *user
	err := u.repo.GetUser(ctx, &user1)
	if err != nil {
		return "", fmt.Errorf("user not found")
	}
	if bcrypt.CompareHashAndPassword([]byte(user1.Password), []byte(user.Password)) != nil {
		return "", fmt.Errorf("wrong password")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(72 * time.Hour)),
			IssuedAt:  jwt.At(time.Now()),
		},
		Username: user.Name,
	})
	tokenString, err := token.SignedString(u.signingKey)
	if err != nil {
		return "", err
	}
	u.tokens[tokenString] = token
	return tokenString, nil
}

func (u *useCase) ValidToken(tokenString string) error {
	err := u.tokens[tokenString].Claims.Valid(jwt.DefaultValidationHelper)
	if err != nil {
		delete(u.tokens, tokenString)
		u.log.Info(fmt.Sprintf("token: %s time out", tokenString))
		return err
	}
	return nil
}
