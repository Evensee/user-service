package dependency

import (
	"github.com/Evensee/user-service/internal"
	"github.com/Evensee/user-service/internal/domain/user"
	"github.com/Evensee/user-service/internal/infrastructure/database/repository"
	"github.com/Evensee/user-service/internal/interface/transaction"
	"gorm.io/gorm"
)

type AppService struct {
	dbConnection *gorm.DB
	userService  user.DomainUserService
}

func (s AppService) GetUserService() user.DomainUserService {
	return s.userService
}

func CreateAppService(appTransaction transaction.AppTransaction, config *internal.AppConfig) AppService {
	db := appTransaction.GetOrmTx()

	return AppService{
		dbConnection: db,
		userService: *user.NewUserDomainService(
			repository.NewUserRepository(db),
			config,
		),
	}
}
