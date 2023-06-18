package restHandler

import (
	"errors"
	"fmt"
	"friendly/internal/model"
	"friendly/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *RestHandler) addUserDescription(ctx *gin.Context) {
	var desc = new(string)

	user, err := GetUserData(ctx)
	if err != nil {
		return
	}

	err = ctx.Bind(&desc)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "incorrect body")
		return
	}

	err = h.userDescriptionService.SaveDescription(ctx, user.Uid, *desc)
	if err != nil {
		utils.Log(fmt.Sprintf("failed save user description. user { %v }", user), err)
		NewErrorResponse(ctx, http.StatusInternalServerError, "failed save user description")
		return
	}

	ctx.JSON(http.StatusOK, DefaultResponse{
		Code:    http.StatusOK,
		Message: "Saved",
	})
}

func (h *RestHandler) getUserDescription(ctx *gin.Context) {
	var user *model.User

	data, exist := ctx.Get(UserData)
	if !exist {
		utils.Log("Not found user data", errors.New("not found user data"))
		NewErrorResponse(ctx, http.StatusInternalServerError, "not found user data")
		return
	}
	user = data.(*model.User)

	userId := ctx.Param("id")
	if userId != "" {
		user.Uid = userId
	}

	desc, err := h.userDescriptionService.GetUserDescription(ctx, userId)

	if err != nil {
		utils.Log("Not found user description", errors.New("not found user description"))
		NewErrorResponse(ctx, http.StatusBadRequest, "not found description")
		return

	}
	ctx.JSON(http.StatusOK, desc)
}
