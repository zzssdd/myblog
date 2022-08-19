package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	filePath := "log/log"
	linkName := "latest_log.log"
	src, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		fmt.Println(err)
	}
	logger := logrus.New()

	logger.Out = src

	logger.SetLevel(logrus.DebugLevel)

	logWriter, _ := retalog.New(
		filePath+"%Y%m%d.log",
		retalog.WithMaxAge(7*24*time.Hour),
		retalog.WithRotationTime(24*time.Hour),
		retalog.WithLinkName(linkName),
	)

	writerMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writerMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(Hook)

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0)))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknow"
		}
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		UserAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		path := c.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"Status":    statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIP,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"UserAgent": UserAgent,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
