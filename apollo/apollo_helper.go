package apollo

import (
	"bytes"
	"context"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/constant"
	apolloConfig "github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/extension"
	"github.com/apolloconfig/agollo/v4/storage"
	"sync/atomic"
)

type listener struct {
	lock   int64
	cancel context.CancelFunc
}

func (l *listener) OnChange(_ *storage.ChangeEvent) {
	if ok := atomic.CompareAndSwapInt64(&l.lock, 0, 1); !ok {
		return
	}

	l.cancel()
}

func (l *listener) OnNewestChange(_ *storage.FullChangeEvent) {}

//GetConfig 获取配置
func GetConfig(appID, ip, secret, cluster string, namespaceName []string) (context.Context, []byte, error) {
	ctx, cancel := context.WithCancel(context.Background())

	l := listener{
		lock:   0,
		cancel: cancel,
	}

	extension.AddFormatParser(constant.YAML, &Parser{})

	buff := bytes.Buffer{}

	for _, name := range namespaceName {
		apollo := &apolloConfig.AppConfig{
			AppID:         appID,
			Cluster:       cluster,
			IP:            ip,
			NamespaceName: name,
			Secret:        secret,
		}

		client, err := agollo.StartWithConfig(func() (*apolloConfig.AppConfig, error) {
			return apollo, nil
		})

		if err != nil {
			return nil, nil, err
		}

		client.AddChangeListener(&l)

		configMap := RangeKey(apollo.NamespaceName, client)
		content := configMap["content"]

		buffer := bytes.NewBufferString(content.(string))
		if err := buffer.WriteByte('\n'); err != nil {
			return nil, nil, err
		}

		buff.Write(buffer.Bytes())
	}

	return ctx, buff.Bytes(), nil
}

func RangeKey(namespace string, client agollo.Client) map[string]interface{} {
	configMap := make(map[string]interface{}, 0)
	cache := client.GetConfigCache(namespace)
	cache.Range(func(key, value interface{}) bool {
		//fmt.Printf("k:%v,v:%v\n", key, value)
		configMap[key.(string)] = value
		return true
	})
	if len(configMap) == 0 {
		panic("config key can not be null")
	}

	return configMap
}

// Parser properties转换器
type Parser struct {
}

// Parse 内存内容=>yml文件转换器
func (d *Parser) Parse(configContent interface{}) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	m["content"] = configContent
	return m, nil
}
