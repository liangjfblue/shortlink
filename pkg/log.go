/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package pkg

import (
	"shortlink/config"

	"github.com/sirupsen/logrus"
)

func InitLog() {
	logrus.SetReportCaller(config.Config().Log.ReportCaller)
	logrus.SetLevel(logrus.Level(config.Config().Log.Level))
	logrus.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp: false,
		PrettyPrint:      true,
		TimestampFormat:  "2006-01-02 115:04:05",
	})
}
