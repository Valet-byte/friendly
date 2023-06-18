package restHandler

import (
	"errors"
	"friendly/internal/model"
	"friendly/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthStatusResponse struct {
	IfExists bool `json:"if_exists"`
}

type AuthTokenResponse struct {
	Token string `json:"token"`
}

type RegistrationRequest struct {
	RegistrationUserData model.RegistrationUserData `json:"registration_user_data"`
	UserData             model.User                 `json:"user_data"`
}

func (h *RestHandler) deleteUser(ctx *gin.Context) {
	var user *model.User

	data, exist := ctx.Get(UserData)
	if !exist {
		utils.Log("Not found user data", errors.New("not found user data"))
		NewErrorResponse(ctx, http.StatusInternalServerError, "not found user data")
		return
	}

	user = data.(*model.User)

	err := h.userService.DeleteUser(ctx, user.Uid)

	if err != nil {
		NewErrorResponse(ctx, http.StatusBadGateway, "failed delete user!")
		return
	}

	ctx.JSON(http.StatusOK, DefaultResponse{
		Code:    http.StatusOK,
		Message: "user deleted",
	})
}

func (h *RestHandler) updateUserData(ctx *gin.Context) {
	var user = new(model.User)

	uid, exist := ctx.Get(UserUid)
	if !exist {
		utils.Log("Not found user uid", errors.New("not found user uid"))
		NewErrorResponse(ctx, http.StatusInternalServerError, "not found user uid")
		return
	}

	err := ctx.BindJSON(user)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "Incorrect body!")
		return
	}

	user.Uid = uid.(string)

	_, err = h.userService.UpdateUser(ctx, *user)

	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "Failed update user data!")
		return
	}

	ctx.JSON(http.StatusOK, DefaultResponse{
		Code:    http.StatusOK,
		Message: "Updated",
	})
}

func (h *RestHandler) getUserData(ctx *gin.Context) {
	var user *model.User

	data, exist := ctx.Get(UserData)
	if !exist {
		utils.Log("Not found user data", errors.New("not found user data"))
		NewErrorResponse(ctx, http.StatusInternalServerError, "not found user data")
		return
	}

	user = data.(*model.User)

	userId := ctx.Param("id")

	if userId == "" {
		ctx.JSON(http.StatusOK, user)
	} else {
		res, err := h.userService.GetUserByUid(ctx, userId)
		if err != nil {
			utils.Log("Not found user data", errors.New("not found user data"))
			NewErrorResponse(ctx, http.StatusBadRequest, "Not found user!")
			return
		}
		ctx.JSON(http.StatusOK, res)
	}
}
