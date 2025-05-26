package model

import (
	"fmt"
	"github.com/Mateus003/user-api/src/configuration/logger"
	"github.com/Mateus003/user-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.Rest_Err{
	logger.Info("Init createUser model", zap.String("journal", "createUser"))

	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}

