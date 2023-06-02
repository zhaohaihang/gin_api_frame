package cron

import (
	"gin_api_frame/internal/dao"

	"github.com/jordan-wright/email"
	"github.com/sirupsen/logrus"
)

type Tasks struct {
	logger   *logrus.Logger
	userDao  *dao.UserDao
	mailPool *email.Pool
}

func NewTasks(l *logrus.Logger, ud *dao.UserDao, mp *email.Pool) *Tasks {
	return &Tasks{
		logger:   l,
		userDao:  ud,
		mailPool: mp,
	}
}

func (t *Tasks) Task1() {

}

func (t *Tasks) Task2() {

}
