package util

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"net"
	"net/http"
	"time"
)

/**
 * @Author: zhaojia
 * @Description:
 * @Version: 1.0.0
 * @Date: 2021/4/7
 */
var Client *elastic.Client

//初始化
func InitEs() {
	HttpClient := &http.Client{}
	HttpClient.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   100,
		IdleConnTimeout:       60 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	fmt.Println("ES连接的IP:" + Cfg.Es.EsIps)
	var err error
	for i := 1; i <= 10; i++ {
		Client, err = elastic.NewClient(elastic.SetHttpClient(HttpClient), elastic.SetSniff(false), elastic.SetURL(Cfg.Es.EsIps), elastic.SetHealthcheck(false))
		//defer client.Close()
		if err != nil {
			log.Println("Elasticsearch,尝试重连……")
		} else {
			break
		}
		if i == 10 {
			log.Println("Elasticsearch连接失败！")
			return
		}
		time.Sleep(6 * time.Second)
	}

	//Client, err = elastic.NewClient(elastic.SetHttpClient(HttpClient),elastic.SetSniff(false), elastic.SetURL(Cfg.Es.EsIps),elastic.SetHealthcheck(false))
	////Client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(Cfg.Es.EsIps))
	//if err != nil {
	//	LoggerInit().Debug("连接失败"+err.Error())
	//	return
	//}
	//info, code, err := Client.Ping(Cfg.Es.EsIps).Do(context.Background())
	//if err != nil {
	//	LoggerInit().Debug("连接失败")
	//	return
	//}
	//LoggerInit().Debug("连接成功：Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	//
	//esversion, err := Client.ElasticsearchVersion(Cfg.Es.EsIps)
	//if err != nil {
	//	LoggerInit().Debug("连接失败")
	//	return
	//}
	//defer Client.Stop()
	//LoggerInit().Debug("连接成功：Elasticsearch version %s\n", esversion)
}
