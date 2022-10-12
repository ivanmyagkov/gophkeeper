package service

import (
	"context"
	"time"

	"gophkeeper/internal/domain"
	"gophkeeper/internal/storage"
	"gophkeeper/pkg/auth"
	"gophkeeper/pkg/hash"
)

type UserSignUpInput struct {
	Login    string
	Password string
}

type UserSignInInput struct {
	Login    string
	Password string
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Users interface {
	SignUp(ctx context.Context, input UserSignUpInput) error
	SignIn(ctx context.Context, input UserSignInInput) (Tokens, error)
	RefreshTokens(ctx context.Context, token string) (Tokens, error)
}

//**********************************************************************************************************************
type Materials interface {
	GetAllTextData(ctx context.Context, userID int) ([]domain.TextData, error)
	UpdateTextDataByID(ctx context.Context, userID int, data domain.TextData) error
	CreateNewTextData(ctx context.Context, userID int, data domain.TextData) error
	CreateNewCredData(ctx context.Context, id int, data domain.CredData) error
	GetAllCredData(ctx context.Context, userID int) ([]domain.CredData, error)
	UpdateCredDataByID(ctx context.Context, userID int, data domain.CredData) error
	CreateNewCardData(ctx context.Context, id int, data domain.CardData) error
	GetAllCardData(ctx context.Context, userID int) ([]domain.CardData, error)
	UpdateCardDataByID(ctx context.Context, userID int, data domain.CardData) error
	CreateNewBlobData(ctx context.Context, id int, data domain.BlobData) error
	GetAllBlobData(ctx context.Context, userID int) ([]domain.BlobData, error)
	UpdateBlobDataByID(ctx context.Context, userID int, data domain.BlobData) error
}

//**********************************************************************************************************************
type Updater interface {
}

//**********************************************************************************************************************
type Services struct {
	Users     Users
	Updater   Updater
	Materials Materials
}

type Deps struct {
	Storages        *storage.Storages
	Hasher          hash.PasswordHasher
	TokenManager    auth.TokenManager
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

func NewServices(deps Deps) *Services {
	users := NewUserService(deps.Hasher, deps.Storages.Users, deps.TokenManager, deps.AccessTokenTTL, deps.AccessTokenTTL)
	updaterService := NewUpdaterService(deps.Storages.Users)
	materials := NewMaterialsService(deps.Storages.Materials)

	return &Services{
		Users:     users,
		Updater:   updaterService,
		Materials: materials,
	}
}
