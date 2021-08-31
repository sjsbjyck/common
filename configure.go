// Package common
// @Description: 在当前程序目录读取conf/app.conf配置参数
package common

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

/*
 *  AppConf
 *  @Description:	返回一个Config对象
 *  @Author: ChenKun
 *  @return *Config
 */
func AppConf() *Config {
	return NewConfig()
}

type Config struct {
	config map[string]string
}

/*
 *  getConfig
 *  @Description:	装载当前应用所在目录conf/app.conf配置文件
 *  @Author: ChenKun
 *  @return *Config
 */
func NewConfig() *Config {
	//获取当前目录
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalln(err)
	}
	_, err = os.Stat(dir + "/conf/app.conf")
	if err != nil {
		if os.IsNotExist(err) {
			panic((dir + "/conf/app.conf文件不存在."))
		} else {
			panic("配置文件错误, err:" + err.Error())
		}
	}

	config := make(map[string]string)
	f, err := os.Open(dir + "/conf/app.conf")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		config[key] = value
	}
	return &Config{config: config}
}

/*
 *  Get
 *  @Description:	获取指定配置key对应的值（返回string）
 *  @Author: ChenKun
 *  @receiver this
 *  @param key
 *  @return string
 */
func (this *Config) GetString(key string) string {
	if val, ok := this.config[key]; ok {
		return val
	} else {
		return ""
	}
}

/*
 *  GetStrings
 *  @Description:	获取指定配置key对应的值（返回字符串数组，使用逗号间隔）
 *  @Author: ChenKun
 *  @receiver this
 *  @param key
 *  @return []string
 */
func (this *Config) GetStrings(key string) []string {
	res := make([]string, 0)
	if val, ok := this.config[key]; ok {
		slic := strings.Split(val, ",")
		for _, item := range slic {
			if item != "" {
				res = append(res, item)
			}
		}
	}
	return res
}
