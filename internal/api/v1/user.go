package v1

import (
	"gin_api_frame/pkg/utils/tokenutil"

	"gin_api_frame/internal/serializer"
	"gin_api_frame/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	logger      *logrus.Logger
	userService *service.UserService
}

func NewUserContrller(l *logrus.Logger, us *service.UserService) *UserController {
	return &UserController{
		logger: l,
		userService: us,
	}
}

// UserLogin godoc
// @Summary 用户登录
// @Description  用户登录接口，如果用户不存在则创建用户
// @Tags user
// @Accept  json
// @Produce  json
// @Param LoginUserInfo body serializer.LoginUserInfo true "login user info"
// @Success 200 {object} serializer.Response{data=serializer.TokenData{user=serializer.User}}
// @Router /api/v1/user/login [post]
func (uc *UserController) UserLogin(c *gin.Context) {
	var loginUserInfo serializer.LoginUserInfo
	if err := c.ShouldBind(&loginUserInfo); err == nil {
		res := uc.userService.Login(loginUserInfo)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		uc.logger.Infoln(err)
	}
}

// UserUpdate godoc
// @Summary 用户更新信息
// @Description  用户更新信息接口
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization header parameter"
// @Param UpdateUserInfo body serializer.UpdateUserInfo true "user update info"
// @Success 200 {object} serializer.Response{data=serializer.User}
// @Router /api/v1/user [put]
func (uc *UserController) UserUpdate(c *gin.Context) {
	var updateUserInfo serializer.UpdateUserInfo
	claims := tokenutil.GetTokenClaimsFromContext(c)
	if err := c.ShouldBind(&updateUserInfo); err == nil {
		res := uc.userService.UpdateUserById(claims.UserID, updateUserInfo)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		uc.logger.Infoln(err)
	}
}

// ViewUser godoc
// @Summary 查看用户信息
// @Description  查看用户信息接口
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization header parameter"
// @Param uid path int true "user ID"
// @Success 200 {object} serializer.Response{data=serializer.User}
// @Router /api/v1/user/{uid} [get]
func (uc *UserController) ViewUser(c *gin.Context) {
	uIdStr := c.Param("uid")

	if uId, err := strconv.ParseUint(uIdStr, 10, 32); err == nil {
		res := uc.userService.GetUserById(uint(uId))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		uc.logger.Infoln(err)
	}
}

// UploadUserAvatar godoc
// @Summary 上传用户头像
// @Description  上传用户头像接口
// @Tags user
// @Accept  multipart/form-data
// @Produce  json
// @Param file formData file true "图片文件"
// @Param Authorization header string true "Authorization header parameter"
// @Success 200 {object} serializer.Response{}
// @Router /api/v1/user/avatar [put]
func (uc *UserController) UploadUserAvatar(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		uc.logger.Infoln(err)
	} else {
		claims := tokenutil.GetTokenClaimsFromContext(c)
		res := uc.userService.UploadUserAvatar(claims.UserID, file, fileHeader)
		c.JSON(http.StatusOK, res)
	}
}

// ChangePasswd godoc
// @Summary  修改密码
// @Description  用户修改密码接口
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization header parameter"
// @Param ChangePasswdInfo body serializer.ChangePasswdInfo true "user changeinfo info"
// @Success 200 {object} serializer.Response{}
// @Router /api/v1/user/changepasswd [put]
func (uc *UserController) ChangePasswd(c *gin.Context) {
	var changePasswdInfo serializer.ChangePasswdInfo
	claims := tokenutil.GetTokenClaimsFromContext(c)
	if err := c.ShouldBind(&changePasswdInfo); err == nil {
		res := uc.userService.ChangePasswd(claims.UserID, changePasswdInfo)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		uc.logger.Infoln(err)
	}
}
