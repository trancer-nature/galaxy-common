package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

//es client
type Client struct {
	name    string //名称
	addr    string
	timeout int64
	Client  *elastic.Client
}

//new es client
func NewEsClient(name string, addr string, timeout int64, userName string, password string) *Client {
	var addrList []string
	addrList = append(addrList, addr)
	var configFunc []elastic.ClientOptionFunc
	configFunc = append(configFunc, elastic.SetSniff(false))
	configFunc = append(configFunc, elastic.SetURL(addrList...))

	if userName != "" {
		configFunc = append(configFunc, elastic.SetBasicAuth(userName, password))
	}

	client, err := elastic.NewClient(configFunc...)
	if err != nil {
		panic(err)
		return nil
	}

	info, code, errPing := client.Ping(addrList[0]).Do(context.Background())
	if errPing != nil {
		panic(errPing.Error() + addrList[0])
		return nil
	}

	fmt.Println(fmt.Sprintf("es ping code. code: %v, info: %v", code, info))
	esClt := &Client{
		name,
		addr,
		timeout,
		client,
	}

	return esClt
}
