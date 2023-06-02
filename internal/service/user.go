package service

import (
	"gin_api_frame/internal/dao"
	"gin_api_frame/internal/model"
	"gin_api_frame/internal/serializer"
	"gin_api_frame/pkg/consts"
	"gin_api_frame/pkg/e"
	"gin_api_frame/pkg/storages/qiniu"
	"gin_api_frame/pkg/utils/tokenutil"
	"strconv"
	"strings"

	"mime/multipart"
	"time"

	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	logger          *logrus.Logger
	userDao      *dao.UserDao
	qiniuStroage *qiniu.QiNiuStroage
}

func NewUserService(l *logrus.Logger, ud *dao.UserDao, qs *qiniu.QiNiuStroage) *UserService {
	return &UserService{
		logger:          l,
		userDao:      ud,
		qiniuStroage: qs,
	}
}

var UserServiceProviderSet = wire.NewSet(NewUserService)

// Login 用户登陆函数
func (us *UserService) Login(loginUserInfo serializer.LoginUserInfo) serializer.Response {
	var user *model.User
	code := e.SUCCESS

	user, exist, _ := us.userDao.ExistOrNotByUserName(loginUserInfo.UserName)
	if exist { // 如果存在，则校验密码
		if !user.CheckPassword(loginUserInfo.Password) {
			code = e.ErrorPasswordNotCompare
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	} else { // 如果不存在则新建
		user = &model.User{
			UserName: loginUserInfo.UserName,
			Status:   model.ACTIVE,
		}
		if loginUserInfo.Type == consts.LOGIN_TYPE_EMAIL {
			user.Phone = loginUserInfo.UserName
		} else {
			user.Email = loginUserInfo.UserName
		}
		// 加密密码
		if err := user.SetPassword(loginUserInfo.Password); err != nil {
			us.logger.Info(err)
			code = e.ErrorUserCreate
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		// 创建用户
		if err := us.userDao.CreateUser(user); err != nil {
			us.logger.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		// 新注册用户，返回的最后一次登录时间为当前时间
		user.LastLogin = time.Now().UnixMilli()
	}

	token, err := tokenutil.GenerateToken(user.ID, loginUserInfo.UserName, 0)
	if err != nil {
		us.logger.Info(err)
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	go us.userDao.UpdateLastLoginById(user.ID, time.Now().UnixMilli())

	return serializer.Response{
		Status: code,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    e.GetMsg(code),
	}
}

// Update 用户修改信息
func (us *UserService) UpdateUserById(uId uint, updateUserInfo serializer.UpdateUserInfo) serializer.Response {
	var user *model.User
	var err error
	code := e.SUCCESS
	// 找到用户
	user, err = us.userDao.GetUserById(uId)
	if err != nil {
		us.logger.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 更新字段
	user.Biography = updateUserInfo.Biography
	user.Address = updateUserInfo.Address
	user.Email = updateUserInfo.Address
	user.Phone = updateUserInfo.Address
	user.Location = model.Point{
		Lat: updateUserInfo.Location.Lat,
		Lng: updateUserInfo.Location.Lng,
	}

	err = us.userDao.UpdateUserById(uId, user)
	if err != nil {
		us.logger.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
	}
}

func (us *UserService) GetUserById(uId uint) serializer.Response {
	var err error
	var user *model.User
	code := e.SUCCESS
	// 找到用户
	user, err = us.userDao.GetUserById(uId)
	if err != nil {
		us.logger.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
	}
}

func (us *UserService) UploadUserAvatar(uId uint, file multipart.File, fileHeader *multipart.FileHeader) serializer.Response {
	code := e.SUCCESS
	var err error

	//重命名文件的名称
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	ti := tm.Format("2006010203040501")
	//提取文件后缀类型
	var ext string
	if pos := strings.LastIndexByte(fileHeader.Filename, '.'); pos != -1 {
		ext = fileHeader.Filename[pos:]
		if ext == "." {
			ext = ""
		}
	}
	filename := "user_avatar/" + strconv.Itoa(int(uId)) + "_" + ti + "_" + strconv.FormatInt(time.Now().Unix(), 10) + ext

	path, err := us.qiniuStroage.UploadToQiNiu(filename, file, fileHeader.Size)

	if err != nil {
		us.logger.Info(err)
		code = e.ErrorUploadFile
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   path,
		Msg:    e.GetMsg(code),
	}
}

func (us *UserService) ChangePasswd(uId uint, changePasswdInfo serializer.ChangePasswdInfo) serializer.Response {

	var code = e.SUCCESS
	user, err := us.userDao.GetUserById(uId)

	if err != nil {
		us.logger.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	pass := user.CheckPassword(changePasswdInfo.OldPasswd)
	if !pass {
		code = e.ErrorPasswordNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	user.SetPassword(changePasswdInfo.NewPasswd)

	err = us.userDao.UpdateUserById(uId, user)
	if err != nil {
		us.logger.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
