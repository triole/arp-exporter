package conf

func (conf Conf) GetHostName(mac string) (name string) {
	if val, ok := conf.Hosts[mac]; ok {
		name = val
	}
	return
}
