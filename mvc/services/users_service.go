package services

import (
	"github.com/quadrosh/microservices/mvc/domailn"
	"github.com/quadrosh/microservices/mvc/utils"
)

// GetUser ...
func GetUser(userID int64) (*domailn.User, *utils.ApplicationError) {
	return domailn.GetUser(userID)
}
