package dependency

import (
	"github.com/Evensee/user-service/internal"
	"github.com/Evensee/user-service/internal/domain/auth"
	"github.com/Evensee/user-service/internal/domain/user"
	databaseRepo "github.com/Evensee/user-service/internal/infrastructure/database/repository"
	memoryRepo "github.com/Evensee/user-service/internal/infrastructure/memory/repository"
	"github.com/Evensee/user-service/internal/interface/transaction"
	"gorm.io/gorm"
)

type AppService struct {
	dbConnection *gorm.DB
	userService  user.DomainUserService
	authService  auth.AuthService
}

func (s AppService) GetUserService() user.DomainUserService {
	return s.userService
}

func (s AppService) GetAuthService() auth.AuthService {
	return s.authService
}

func CreateAppService(appTransaction transaction.AppTransaction, config *internal.AppConfig) AppService {
	db := appTransaction.GetOrmTx()
	rdb := appTransaction.GetMemoryTx()

	println("CreateAppService")

	return AppService{
		dbConnection: db,
		userService: *user.NewUserDomainService(
			databaseRepo.NewUserRepository(db),
		),
		authService: *auth.NewAuthService(
			memoryRepo.NewAuthTokenRepository(rdb),
			databaseRepo.NewUserRepository(db),
			config,
		),
	}
}
