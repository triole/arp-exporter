package conf

import (
	"os"
	"reflect"
	"strings"

	"github.com/triole/logseal"
	yaml "gopkg.in/yaml.v3"
)

type Conf struct {
	ArpTable       string
	HostnameConfig string
	Server         bool
	Bind           string
	Info           string
	List           bool
	EnableVendors  bool
	Hosts          map[string]tHost
	Lg             logseal.Logseal
}

type tHost struct {
	MAC  string `yaml:"mac"`
	Name string `yaml:"name"`
	Itf  string `yaml:"itf"`
}

func Init(cli interface{}, lg logseal.Logseal) (conf Conf) {
	conf.Lg = lg
	conf.ArpTable = getcli(cli, "ArpTableFile").(string)
	if conf.ArpTable != "" {
		conf.ArpTable = absPathFatal(conf.ArpTable, lg)
		existsFatal(conf.ArpTable, lg)
	}
	conf.HostnameConfig = getcli(cli, "HostnameConfig").(string)
	if conf.HostnameConfig != "" {
		conf.HostnameConfig = absPathFatal(conf.HostnameConfig, lg)
		existsFatal(conf.HostnameConfig, lg)
		conf.loadHostnameConfig()
	}
	conf.Bind = getcli(cli, "Bind").(string)
	conf.Info = getcli(cli, "MacInfo").(string)
	conf.List = getcli(cli, "ListVendors").(bool)
	conf.EnableVendors = getcli(cli, "EnableVendors").(bool)
	return
}

func getcli(cli interface{}, keypath string) (r interface{}) {
	key := strings.Split(keypath, ".")
	val := reflect.ValueOf(cli)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		if fieldType.Name == key[0] {
			r = field.Interface()
			if len(key) > 1 {
				return getcli(r, key[1])
			}
			if fieldType.Type.Name() == "" {
				r = true
			} else {
				r = field.Interface()
			}
		}
	}
	return
}

func (conf *Conf) loadHostnameConfig() {
	var hostnamesList []tHost
	var raw []byte
	var err error
	raw, err = os.ReadFile(conf.HostnameConfig)
	if err != nil {
		conf.Lg.Error(
			"can not read config",
			logseal.F{"path": conf.HostnameConfig, "error": err},
		)
	}
	err = yaml.Unmarshal(raw, &hostnamesList)
	if err != nil {
		conf.Lg.Error(
			"can not unmarshal conf",
			logseal.F{"path": conf.HostnameConfig, "error": err},
		)
	}
	conf.Hosts = make(map[string]tHost)
	for _, host := range hostnamesList {
		conf.Hosts[host.MAC] = host
	}
	return
}
