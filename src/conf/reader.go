package conf

func (conf Conf) GetHostName(mac string) (host tHost) {
	if val, ok := conf.Hosts[mac]; ok {
		host = val
	}
	return
}
