package entities

type Port struct {
	Number   int    `json:"number"`
	Protocol string `json:"protocol"`
}

type Device struct {
	Hostname        string `json:"hostname"`
	MacAddress      string `json:"macAddress"`
	Ipv4            string `json:"ipv4"`
	Ipv6            string `json:"ipv6"`
	Oid             string `json:"oid"`
	OperatingSystem string `json:"operatingSystem"`
	Ports           []Port `json:"openPorts"`
}
