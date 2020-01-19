package repositories

import (
	"context"
	"github.com/udayangaac/mobile-api/internal/entities"
)

type MobileAppUserRepo interface {
	//AddMobileUser(ctx context.Context, mobileUser entities.MobileAppUser) (isUpdate bool)
	AddMobileUser(ctx context.Context, mobileUser entities.MobileAppUser)
	GetMobileUserByEmail(ctx context.Context, email string) (mobileUser entities.MobileAppUser, err error)
}
