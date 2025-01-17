package model

import (
	"GOMS-BACKEND-GO/model/data/constant"
	"GOMS-BACKEND-GO/model/data/input"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Email      string
	Password   string
	Grade      int
	Name       string
	Gender     constant.Gender
	Major      constant.Major
	ProfileURL *string
	Authority  constant.Authority
	CreatedAt  time.Time
}

type AccountUseCase interface {
	WithDrawAccount(c *gin.Context, ctx context.Context) error
}

type AccountRepository interface {
	SaveAccount(ctx context.Context, account *Account) error
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	FindByEmail(ctx context.Context, email string) (*Account, error)
	FindByAccountID(ctx context.Context, accountID primitive.ObjectID) (*Account, error)
	FindAllAccount(ctx context.Context) ([]Account, error)
	FindByAccountByStudentInfo(ctx context.Context, searchAccountInput *input.SearchAccountInput) ([]Account, error)
	UpdateAccountAuthority(ctx context.Context, authorityInput *input.UpdateAccountAuthorityInput) error
	DeleteAccount(ctx context.Context, account *Account) error
}
