package common

import (
	"flag"
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v2"
)

var (
	configOnce sync.Once
	configer   *cfg
)

type cfg struct {
	Runmode string `yaml:"runmode"`
	Name    string `yaml:"name"`
	Http    struct {
		Port int `yaml:"port"`
	} `yaml:"http"`
	Https struct {
		Port    int    `yaml:"port"`
		CrtPath string `yaml:"crt_path"`
		KeyPath string `yaml:"key_path"`
	} `yaml:"https"`
	Log struct {
		FilePath string `yaml:"filePath"`
	} `yaml:"log"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
		Db       int    `yaml:"db"`
	} `yaml:"redis"`
	Rabbitmq struct {
		Addr string `yaml:"addr"`
	} `yaml:"rabbitmq"`
	Mysql struct {
		Master struct {
			Addr         string `yaml:"addr"`
			MaxOpenConns int    `yaml:"maxOpenConns"`
			MaxIdleConns int    `yaml:"maxIdleConns"`
		} `yaml:"master"`
		AppRead struct {
			Addr         string `yaml:"addr"`
			MaxOpenConns int    `yaml:"maxOpenConns"`
			MaxIdleConns int    `yaml:"maxIdleConns"`
		} `yaml:"appRead"`
		AdminRead struct {
			Addr         string `yaml:"addr"`
			MaxOpenConns int    `yaml:"maxOpenConns"`
			MaxIdleConns int    `yaml:"maxIdleConns"`
		} `yaml:"adminRead"`
	} `yaml:"mysql"`
	ScriptPath string `yaml:"scriptPath"`
	Pgsql      struct {
		Master struct {
			Addr         string `yaml:"addr"`
			Port         int    `yaml:"port"`
			MaxOpenConns int    `yaml:"maxOpenConns"`
			MaxIdleConns int    `yaml:"maxIdleConns"`
		} `yaml:"master"`
		AppRead struct {
			Addr         string `yaml:"addr"`
			MaxOpenConns int    `yaml:"maxOpenConns"`
			MaxIdleConns int    `yaml:"maxIdleConns"`
		} `yaml:"appRead"`
		AdminRead struct {
			Addr         string `yaml:"addr"`
			MaxOpenConns int    `yaml:"maxOpenConns"`
			MaxIdleConns int    `yaml:"maxIdleConns"`
		} `yaml:"adminRead"`
	} `yaml:"pgsql"`
	Grpc struct {
		Port             int      `yaml:"port"`
		AdminAuthServers []string `yaml:"admin_auth_servers"`
		BaseServers      []string `yaml:"base_servers"`
		WxServers        []string `yaml:"wx_servers"`
		AuthServers      []string `yaml:"auth_servers"`
		UserServers      []string `yaml:"user_servers"`
		CalculateServers []string `yaml:"calculate_servers"`
		OrderServers     []string `yaml:"order_servers"`
	} `yaml:"grpc"`
	Jwt struct {
		Secret string `yaml:"secret"`
		Ttl    int64  `yaml:"ttl"`
	} `yaml:"jwt"`
	Cos struct {
		Domain string `yaml:"domain"`
	} `yaml:"cos"`
	Oss struct {
		AccessKeyId     string `yaml:"access_key_id"`
		AccessKeySecret string `yaml:"access_key_secret"`
		Endpoint        string `yaml:"endpoint"`
		Domain          string `yaml:"domain"`
	} `yaml:"oss"`
	ShortHost string `yaml:"short_host"`
	MiniApp   struct {
		Appid  string `yaml:"appid"`
		Secret string `yaml:"secret"`
	} `yaml:"miniApp"`
	TtMiniApp struct {
		Appid  string `yaml:"appid"`
		Secret string `yaml:"secret"`
	} `yaml:"ttMiniApp"`
	WxMc struct {
		Id     string `yaml:"id"`
		Apikey string `yaml:"apikey"`
	} `yaml:"wx_mc"`
	Pay struct {
		WxCallback string `yaml:"wx_callback"`
	} `yaml:"pay"`
	Algorithm struct {
		Addr        string   `yaml:"addr"`
		FashionCmd  string   `yaml:"fashion_cmd"`
		FashionArgs []string `yaml:"fashion_args"`
	} `yaml:"algorithm"`
	Dispatch struct {
		Host    string `yaml:"host"`
		CutPart struct {
			Run     int      `yaml:"run"`
			Cmd     string   `yaml:"cmd"`
			Args    []string `yaml:"args"`
			Timeout int      `yaml:"timeout"`
		} `yaml:"cut_part"`
		ProductionMaterial struct {
			Run     int      `yaml:"run"`
			Cmd     string   `yaml:"cmd"`
			Args    []string `yaml:"args"`
			Timeout int      `yaml:"timeout"`
		} `yaml:"production_material"`
		ComposingMaterial struct {
			Run     int      `yaml:"run"`
			Cmd     string   `yaml:"cmd"`
			Args    []string `yaml:"args"`
			Timeout int      `yaml:"timeout"`
		} `yaml:"composing_material"`
	} `yaml:"dispatch"`
}

func GetConfig() *cfg {
	configOnce.Do(func() {
		conf := &cfg{}
		path := ""
		flag.StringVar(&path, "conf", "./configs/dev.yml", "help")
		flag.Parse()
		bytes, err := ioutil.ReadFile(path)
		if nil != err {
			return
		}
		err = yaml.Unmarshal(bytes, conf)
		if nil != err {
			return
		}
		configer = conf
	})
	return configer
}
