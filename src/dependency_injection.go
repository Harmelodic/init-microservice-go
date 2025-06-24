package main

import (
	"github.com/Harmelodic/init-microservice-go/src/account"
	"github.com/Harmelodic/init-microservice-go/src/commons"
	"github.com/gin-gonic/gin"
)

func dependencyInjection(engine *gin.Engine) {
	account.Controller(engine, account.Service{})

	commons.HealthController(engine)
}
