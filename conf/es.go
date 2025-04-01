package conf

import (
	"log"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"

)

var EsClient *elastic.Client

// InitEs 初始化es
func InitEs() {
	esConn := "http://" + EsHost + ":" + EsPort
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(esConn))
	if err != nil {
		log.Panic(err)
	}
	EsClient = client
}

// EsHookLog 初始化log日志
func EsHookLog() *elogrus.ElasticHook {
	hook, err := elogrus.NewElasticHook(EsClient, EsHost, logrus.DebugLevel, EsIndex)
	if err != nil {
		log.Panic(err)
	}
	return hook
}
