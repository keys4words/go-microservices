package controllers

import (
	"mvc/services"
	"mvc/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.AppError{
			Message:    "user_id must be a integer\n",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	user, apiErr := services.UsersService.GetUser(userId)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, user)
	c.JSON(http.StatusOK, user)
}
