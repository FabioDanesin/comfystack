package auth_routes

import (
	"comfystack/data/models"
	utils "comfystack/services/database"
	"comfystack/services/logger"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginPostRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func POST_login(ctx *gin.Context) {
	dbConn := utils.GetConnectionString()

	var loginPostReq LoginPostRequest
	if err := ctx.ShouldBindJSON(loginPostReq); err != nil {
		panic(err)
	}

	// Check dell'esistenza utente
	result, err := dbConn.
		NewSelect().
		Model(new(models.Utente)).
		Limit(1).
		Where("nickname = ? AND password = ?", loginPostReq.Name, loginPostReq.Password).
		Count(context.Background())

	if err != nil {
		panic(err)
	} else if result != 1 {
		logger.Instance.LogWrite("Error: found more than 1 matching user")
	}

	// Inserimento del token
	var token models.Token
	insertErr := dbConn.
		NewInsert().
		Model(new(models.Token)).
		Scan(context.Background(), token)

	if insertErr != nil {
		panic(insertErr)
	} else {
		logger.Instance.LogWrite(`Token created: ` + token.Token)
		ctx.SetCookie("cmf-token",
			token.Token,
			int(token.ValidityTime.Unix()), // Questa esiste grazie all'hook impostato su model.Token
			"",
			"localhost", // TODO: capire come mettere il nome del sito corretto.
			true,
			true,
		)
		ctx.String(http.StatusAccepted, "Login ok")
	}
}

func InitializeAuthEndpoints(engine *gin.Engine) *gin.RouterGroup {
	authRoutes := engine.Group("/auth")
	authRoutes.POST("/login", POST_login)
	return authRoutes
}
