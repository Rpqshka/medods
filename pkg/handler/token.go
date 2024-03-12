package handler

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"medods"
	"net/http"
)

type userInput struct {
	GUID    string
	Refresh string
	isUsed  bool
}

func (h *Handler) getTokensByGUID(c *gin.Context) {
	guid := c.Param("guid")

	_, err := uuid.FromString(guid)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	access, err := generateAccessToken(guid)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	refresh, err := generateRefreshToken()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err = h.services.CheckUser(guid); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err = h.services.SetRefreshToken(guid, refresh); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	base64Refresh := base64.StdEncoding.EncodeToString(refresh)

	c.JSON(http.StatusOK, map[string]interface{}{
		"access":  access,
		"refresh": base64Refresh,
	})
}

func (h *Handler) refreshTokens(c *gin.Context) {
	var user medods.User
	base64Refresh := c.Param("refresh")

	decodedRefresh, err := decodeBase64(base64Refresh)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if user, err = h.services.GetUser(decodedRefresh); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	access, err := generateAccessToken(user.GUID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	refresh, err := generateRefreshToken()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user.Refresh = string(refresh)

	if err = h.services.UpdateTokens(user); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	base64Refresh = base64.StdEncoding.EncodeToString(refresh)

	c.JSON(http.StatusOK, map[string]interface{}{
		"access":  access,
		"refresh": base64Refresh,
	})
}
