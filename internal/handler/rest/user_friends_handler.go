package restHandler

import (
	"errors"
	"friendly/internal/model"
	"friendly/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *RestHandler) getFriends(ctx *gin.Context) {
	var user *model.User

	data, exist := ctx.Get(UserData)
	if !exist {
		utils.Log("Not found user data", errors.New("not found user data"))
		NewErrorResponse(ctx, http.StatusUnauthorized, "not found user data")
		return
	}
	user = data.(*model.User)

	userId := ctx.Param("id")
	if userId != "" {
		userId = user.Uid
	}

	limitParam := ctx.Query("limit")
	if limitParam == "" {
		utils.Log("Not found limit param", errors.New("not found limit param"))
		NewErrorResponse(ctx, http.StatusBadRequest, "Not found limit param")
		return
	}

	pageParam := ctx.Query("page")
	if pageParam == "" {
		utils.Log("Not found page param", errors.New("not found limit page"))
		NewErrorResponse(ctx, http.StatusBadRequest, "Not found page param")
		return
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		utils.Log("Incorrect limit format", errors.New("incorrect limit format"))
		NewErrorResponse(ctx, http.StatusBadRequest, "Incorrect limit format")
		return
	}
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		utils.Log("Incorrect page format", errors.New("incorrect page format"))
		NewErrorResponse(ctx, http.StatusBadRequest, "Incorrect page format")
		return
	}

	res, err := h.userFriendsService.GetFriends(ctx, userId, limit, page)
	if err != nil {
		utils.Log("Not found friends", errors.New("not found friends"))
		NewErrorResponse(ctx, http.StatusNoContent, "Not found friends")
		return
	}

	ctx.JSON(http.StatusOK, DefaultResponse{
		Code:    http.StatusOK,
		Message: res,
	})
}
func (h *RestHandler) generateFriendToken(ctx *gin.Context) {
	var user *model.User

	data, exist := ctx.Get(UserData)
	if !exist {
		utils.Log("Not found user data", errors.New("not found user data"))
		NewErrorResponse(ctx, http.StatusUnauthorized, "not found user data")
		return
	}
	user = data.(*model.User)

	token, err := h.userFriendsService.GetFriendsToken(user.Uid)
	if err != nil {
		utils.Log("Failed create friend token", errors.New("failed create friend token"))
		NewErrorResponse(ctx, http.StatusBadRequest, "Failed create friend token")
		return
	}
	ctx.JSON(http.StatusOK, DefaultResponse{
		Code:    http.StatusOK,
		Message: token,
	})
}
func (h *RestHandler) addFriendsByParam(ctx *gin.Context) {
	var user *model.User

	data, exist := ctx.Get(UserData)
	if !exist {
		utils.Log("Not found user data", errors.New("not found user data"))
		NewErrorResponse(ctx, http.StatusUnauthorized, "not found user data")
		return
	}
	user = data.(*model.User)

	friendId := ctx.Param("id")

	err := h.userFriendsService.AddFriends(ctx, user.Uid, friendId)
	if err != nil {
		utils.Log("Failed add friend", errors.New("failed add friend"))
		NewErrorResponse(ctx, http.StatusBadRequest, "Failed add friend")
		return
	}

	ctx.JSON(http.StatusOK, DefaultResponse{
		Code:    http.StatusOK,
		Message: "successful",
	})
}

func (h *RestHandler) addFriendsByToken(ctx *gin.Context) {
	var token = new(string)
	var user *model.User

	data, exist := ctx.Get(UserData)
	if !exist {
		utils.Log("Not found user data", errors.New("not found user data"))
		NewErrorResponse(ctx, http.StatusUnauthorized, "not found user data")
		return
	}
	user = data.(*model.User)

	err := ctx.Bind(token)

	if err != nil {
		utils.Log("Failed parse token", errors.New("failed parse token"))
		NewErrorResponse(ctx, http.StatusBadRequest, "Failed parse token")
		return
	}

	err = h.userFriendsService.AddFriendsByFriendsToken(ctx, user.Uid, *token)
	if err != nil {
		utils.Log("Failed add friend", errors.New("failed add friend"))
		NewErrorResponse(ctx, http.StatusBadRequest, "Failed add friend")
		return
	}

	ctx.JSON(http.StatusOK, DefaultResponse{
		Code:    http.StatusOK,
		Message: "successful",
	})
}

func (h *RestHandler) deleteFriendsByParam(ctx *gin.Context) {
	var user *model.User

	data, exist := ctx.Get(UserData)
	if !exist {
		utils.Log("Not found user data", errors.New("not found user data"))
		NewErrorResponse(ctx, http.StatusUnauthorized, "not found user data")
		return
	}
	user = data.(*model.User)

	friendId := ctx.Param("id")

	err := h.userFriendsService.DeleteFriends(ctx, user.Uid, friendId)
	if err != nil {
		utils.Log("Failed add friend", errors.New("failed add friend"))
		NewErrorResponse(ctx, http.StatusBadRequest, "Failed add friend")
		return
	}

	ctx.JSON(http.StatusOK, DefaultResponse{
		Code:    http.StatusOK,
		Message: "successful",
	})
}
