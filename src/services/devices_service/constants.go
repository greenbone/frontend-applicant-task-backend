package devices_service

import "applicant-backend/src/entities"

var DUMMY_DEVICES = []entities.Device{
	entities.Device{Oid: "1", Hostname: "Laboratory1", Ipv4: "192.168.178.13", MacAddress: "F7-82-9F-FE-95-17", Ipv6: "7c14:03d8:9539:aa61:fbbc:b55d:1fe0:942e", OperatingSystem: "Windows", Ports: []entities.Port{{Number: 3306, Protocol: "TCP"}, {Number: 22, Protocol: "TCP"}}},
	entities.Device{Oid: "2", Hostname: "Laboratory2", Ipv4: "192.168.178.14", MacAddress: "F5-10-5F-A8-F6-BA", Ipv6: "64f1:dede:0c01:46ba:8847:9f95:b4a6:d1ce", OperatingSystem: "Linux Debian", Ports: []entities.Port{{Number: 25565, Protocol: "TCP"}, {Number: 22, Protocol: "TCP"}}},
	entities.Device{Oid: "3", Hostname: "Laboratory3", Ipv4: "192.168.178.15", MacAddress: "BD-00-4D-78-52-17", Ipv6: "e5d7:fff3:1035:1b70:5481:8362:9e05:e16c", OperatingSystem: "Linux NixOS", Ports: []entities.Port{{Number: 27017, Protocol: "TCP"}, {Number: 22, Protocol: "TCP"}}},
	entities.Device{Oid: "4", Hostname: "Entrance1", Ipv4: "192.168.178.43", MacAddress: "CA-06-3E-08-76-36", Ipv6: "3183:7bf8:1273:574f:f3a5:5faf:b777:6857", OperatingSystem: "Windows Server 2017", Ports: []entities.Port{{Number: 445, Protocol: "TCP"}, {Number: 22, Protocol: "TCP"}}},
	entities.Device{Oid: "5", Hostname: "Entrance2", Ipv4: "192.168.178.44", MacAddress: "E5-C1-29-4B-CB-71", Ipv6: "d6cc:e97e:dc3b:08f8:26b8:331d:7261:7d1b", OperatingSystem: "MacOS", Ports: []entities.Port{{Number: 443, Protocol: "TCP"}, {Number: 22, Protocol: "TCP"}}},
	entities.Device{Oid: "6", Hostname: "Entrance3", Ipv4: "192.168.178.45", MacAddress: "6C-0E-8B-84-94-56", Ipv6: "c5e5:0a7f:a69f:b40e:f029:48fd:d2f3:4bd0", OperatingSystem: "Windows", Ports: []entities.Port{{Number: 80, Protocol: "TCP"}, {Number: 22, Protocol: "TCP"}}},
}
