package conf

import (
	"os"
	"path/filepath"
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
	Hosts          map[string]string
	Lg             logseal.Logseal
}

func Init(cli interface{}, lg logseal.Logseal) (conf Conf) {
	conf.ArpTable = getcli(cli, "ArpTableFile").(string)
	conf.HostnameConfig = getcli(cli, "HostnameConfig").(string)
	var err error
	if conf.HostnameConfig != "" {
		conf.HostnameConfig, err = filepath.Abs(conf.HostnameConfig)
		lg.IfErrFatal(
			"unable to construct full path",
			logseal.F{"path": conf.HostnameConfig, "error": err},
		)

		_, err = os.Stat(conf.HostnameConfig)
		lg.IfErrFatal(
			"file does not exist",
			logseal.F{"path": conf.HostnameConfig, "error": err},
		)

		var raw []byte
		raw, err = os.ReadFile(conf.HostnameConfig)
		if err != nil {
			lg.Error(
				"can not read config",
				logseal.F{"path": conf.HostnameConfig, "error": err},
			)
		}
		err = yaml.Unmarshal(raw, &conf.Hosts)
		if err != nil {
			lg.Error(
				"can not unmarshal conf",
				logseal.F{"path": conf.HostnameConfig, "error": err},
			)
		}
	}
	conf.Bind = getcli(cli, "Bind").(string)
	conf.Info = getcli(cli, "Info").(string)
	conf.List = getcli(cli, "List").(bool)
	conf.EnableVendors = getcli(cli, "EnableVendors").(bool)
	conf.Lg = lg
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
