package restHandler

import (
	"friendly/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *RestHandler) addUserDescription(ctx *gin.Context) {
	var user *model.User
	var desc = new(string)

	data, exist := ctx.Get(UserData)
	if !exist {
		logrus.Error("Not found user data! user_description_handler.addUserDescription")
		NewErrorResponse(ctx, http.StatusInternalServerError, "not found user data")
		return
	}

	user = data.(*model.User)

	err := ctx.Bind(&desc)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "incorrect body")
		return
	}

	ok, err := h.fb.AddDescription(ctx, user.Uid, *desc)
	if !ok || err != err {
		logrus.Error("failed save user description. err: ", err, " ok: ", ok, " user: ", user, " desc: ", desc)
		NewErrorResponse(ctx, http.StatusInternalServerError, "failed save user description")
		return
	}

	_, err = h.rd.PutUserDescription(user.Uid, *desc)

	if err != nil {
		logrus.Error("Failed save data on redis: err: ", err)
	}

	ctx.JSON(http.StatusOK, DefaultResponse{Message: "Saved"})
}

func (h *RestHandler) getUserDescription(ctx *gin.Context) {
	var user *model.User

	data, exist := ctx.Get(UserData)
	if !exist {
		logrus.Error("Not found user data! user_description_handler.getUserDescription")
		NewErrorResponse(ctx, http.StatusInternalServerError, "not found user data")
		return
	}
	user = data.(*model.User)

	userId := ctx.Param("id")
	if userId != "" {
		user.Uid = userId
	}

	desc, err := h.rd.GetUserDescription(user.Uid)

	if err != nil {
		desc, err = h.fb.GetUserDescription(ctx, user.Uid)
		if err != nil {
			NewErrorResponse(ctx, http.StatusBadRequest, "not found description")
			return
		}
	}
	ctx.JSON(http.StatusOK, desc)
}
