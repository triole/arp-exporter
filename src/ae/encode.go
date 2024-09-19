package ae

func (ae *tAE) makePrometheusMetrics() (metrics string) {
	for _, el := range ae.ArpTable {
		metrics += "#HELP\n"
		metrics += "#TYPE\n"
		metrics += "arp_exporter{"
		metrics += "ip=\"" + el.IP + "\", mac=\"" + el.MAC + "\""
		if el.Itf != "" {
			metrics += ", itf=\"" + el.Itf + "\""
		}
		if el.Name != "" {
			metrics += ", name=\"" + el.Name + "\""
		}
		if el.Vendor != "" {
			metrics += ", vendor=\"" + el.Vendor + "\""
		}
		metrics += "} 0\n"
	}
	return
}
