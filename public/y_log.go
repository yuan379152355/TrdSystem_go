package tspublic

import (
	"github.com/sirupsen/logrus"
	//"os"
)

var log = logrus.New()

type SYLog struct {
	log *logrus.Logger
}

// 初始化
//func (slog *SYLog) Init()  {
//	log.LoadConfiguration("./log.json")
//	slog.Log = log
//}

//func (slog *SYLog)TsWriteLog()  {

//}
