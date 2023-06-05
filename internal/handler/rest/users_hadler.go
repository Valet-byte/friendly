package restHandler

import (
	"friendly/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type AuthStatusResponse struct {
	IfExists bool `json:"if_exists"`
}

type AuthTokenResponse struct {
	Token string `json:"token"`
}

func (h *RestHandler) auth(ctx *gin.Context) {
	token, err := GetUserToken(ctx)

	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	status, err := h.fb.CheckToken(ctx, token)

	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "incorrect token format")
		return
	}

	response := AuthStatusResponse{IfExists: status}

	ctx.JSON(http.StatusOK, response)
}

type RegistrationRequest struct {
	RegistrationUserData model.RegistrationUserData `json:"registration_user_data"`
	UserData             model.User                 `json:"user_data"`
}

type RegistrationResponse struct {
	FbToken string `json:"token"`
}

func (h *RestHandler) registrationUser(ctx *gin.Context) {
	var request RegistrationRequest

	err := ctx.BindJSON(&request)

	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "incorrect body format")
		return
	}

	uid, err := h.fb.RegistrationUser(ctx, request.RegistrationUserData)

	if err != nil {
		NewErrorResponse(ctx, http.StatusBadGateway, "registration failed")
		return
	}

	ok, err := h.fb.AddUserData(ctx, uid, request.UserData)

	if err != nil || !ok {
		NewErrorResponse(ctx, http.StatusBadGateway, "add user data failed")
		return
	}

	token, err := h.fb.GetToken(ctx, uid)

	if err != nil {
		NewErrorResponse(ctx, http.StatusBadGateway, "failed generate token")
		return
	}

	response := RegistrationResponse{token}

	ctx.JSON(http.StatusOK, response)
}

func (h *RestHandler) deleteUser(ctx *gin.Context) {
	var user *model.User

	data, exist := ctx.Get(UserData)
	if !exist {
		logrus.Error("Not found user data! user_handler.deleteUser")
		NewErrorResponse(ctx, http.StatusInternalServerError, "not found user data")
		return
	}

	user = data.(*model.User)

	err := h.fb.DeleteUser(ctx, user.Uid)

	if err != nil {
		NewErrorResponse(ctx, http.StatusBadGateway, "failed delete user!")
		return
	}

	err = h.rd.DeleteUser(user.Uid)

	if err != nil {
		NewErrorResponse(ctx, http.StatusBadGateway, "failed delete user!")
		return
	}

	ctx.JSON(http.StatusOK, DefaultResponse{"user deleted"})
}

func (h *RestHandler) updateUserData(ctx *gin.Context) {
	var user = new(model.User)
	var updateUserData = new(model.User)

	data, exist := ctx.Get(UserData)
	if !exist {
		logrus.Error("Not found user data! user_handler.deleteUser")
		NewErrorResponse(ctx, http.StatusInternalServerError, "not found user data")
		return
	}

	user = data.(*model.User)
	err := ctx.BindJSON(updateUserData)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "Incorrect body!")
		return
	}

	user.Combine(updateUserData)

	update, err := h.fb.AddUserData(ctx, user.Uid, *updateUserData)

	if err != nil || !update {
		NewErrorResponse(ctx, http.StatusInternalServerError, "Failed update user data!")
		return
	}

	uid, err := h.rd.PutUser(user.Uid, user)

	if err != nil {
		logrus.Error("Failed update cash for user data! err: ", err, " uid: ", uid, "UserData: ", user)
		err = h.rd.DeleteUser(user.Uid)
		return
	}

	ctx.JSON(http.StatusOK, DefaultResponse{"Updated"})
}

func (h *RestHandler) getUserData(ctx *gin.Context) {
	var user *model.User

	data, exist := ctx.Get(UserData)
	if !exist {
		logrus.Error("Not found user data! user_handler.deleteUser")
		NewErrorResponse(ctx, http.StatusInternalServerError, "not found user data")
		return
	}

	user = data.(*model.User)

	userId := ctx.Param("id")

	if userId == "" {
		ctx.JSON(http.StatusOK, user)
	} else {
		res, err := h.fb.GetUserData(ctx, userId)
		if err != nil {
			NewErrorResponse(ctx, http.StatusBadRequest, "Not found user!")
			logrus.Error("users_handler.getUserData, 160. err : ", err)
			return
		}
		ctx.JSON(http.StatusOK, res)
	}
}
