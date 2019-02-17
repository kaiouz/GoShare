package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 全局配置
var Config Conf

// 配置结构
type Conf struct {
	Dir      []string `yaml:"dir"`      // 扫描的目录
	CacheDir string   `yaml:"cacheDir"` // 缓存目录
	Port     int      `yaml:"port"`     // 服务端口
}

const (
	defaultConfFile = "config.yaml" // 默认的配置文件名称
)

// 获取执行文件的目录
func appDir() (string, error) {
	if exec, err := os.Executable(); err == nil {
		return filepath.Dir(exec), nil
	} else {
		return "", err
	}
}

// 工作目录
func wdDir() (string, error) {
	if wd, err := os.Getwd(); err == nil {
		return wd, nil
	} else {
		return "", err
	}
}

// 返回配置文件
func configFileData() ([]byte, error) {
	var err error
	if confDir, err := wdDir(); err == nil {
		if data, err := ioutil.ReadFile(filepath.Join(confDir, defaultConfFile)); err != nil {
			if os.IsNotExist(err) {
				if confDir, err := appDir(); err == nil {
					return ioutil.ReadFile(filepath.Join(confDir, defaultConfFile))
				}
			}
		} else {
			return data, nil
		}
	}
	return nil, err
}

// 初始化配置
func InitConfig() error {
	if data, err := configFileData(); err == nil {
		Config = Conf{}
		err = yaml.Unmarshal(data, &Config)
		return nil
	} else {
		return fmt.Errorf("读取配置文件失败: %s", err)
	}
}
