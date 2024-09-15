package conf

import (
	"os"
	"path/filepath"

	"github.com/triole/logseal"
	yaml "gopkg.in/yaml.v3"
)

type Conf struct {
	ConfigFile    interface{}
	Bind          string
	Info          string
	List          bool
	EnableVendors bool
	Hosts         map[string]string
	Lg            logseal.Logseal
}

func Init(configFile, bind string, info string, list, enableVendors bool, lg logseal.Logseal) (conf Conf) {
	conf.ConfigFile = nil
	if configFile != "" {
		cf, err := filepath.Abs(configFile)
		conf.ConfigFile = cf
		lg.IfErrFatal(
			"unable to construct full path",
			logseal.F{"path": configFile, "error": err},
		)

		_, err = os.Stat(configFile)
		lg.IfErrFatal(
			"file does not exist",
			logseal.F{"path": configFile, "error": err},
		)

		var raw []byte
		raw, err = os.ReadFile(configFile)
		if err != nil {
			lg.Error(
				"can not read config",
				logseal.F{"path": configFile, "error": err},
			)
		}
		err = yaml.Unmarshal(raw, &conf.Hosts)
		if err != nil {
			lg.Error(
				"can not unmarshal conf",
				logseal.F{"path": configFile, "error": err},
			)
		}
	}
	conf.Bind = bind
	conf.Info = info
	conf.List = list
	conf.EnableVendors = enableVendors
	conf.Lg = lg
	return
}
