package ae

func (arp *tAE) ReturnPrometheusMetrics() (metrics string) {
	for _, el := range arp.ArpTable {
		metrics += "#HELP\n"
		metrics += "#TYPE\n"
		metrics += "arp_exporter{" +
			"ip=\"" + el.IP + "\", mac=\"" + el.MAC +
			"\"} 0\n"
	}
	return
}
