package devices_service

import "applicant-backend/src/entities"

var DUMMY_DEVICES = []entities.Device{
	{Oid: "1", Hostname: "Laboratory1", Ipv4: "192.168.178.13", MacAddress: "F7-82-9F-FE-95-17", Ipv6: "7c14:03d8:9539:aa61:fbbc:b55d:1fe0:942e", OperatingSystem: "Windows", Ports: []entities.Port{{Number: 3306, Protocol: "TCP"}, {Number: 22, Protocol: "TCP"}}},
	{Oid: "2", Hostname: "Laboratory2", Ipv4: "192.168.178.14", MacAddress: "F5-10-5F-A8-F6-BA", Ipv6: "64f1:dede:0c01:46ba:8847:9f95:b4a6:d1ce", OperatingSystem: "Linux Debian", Ports: []entities.Port{{Number: 25565, Protocol: "TCP"}, {Number: 22, Protocol: "TCP"}}},
	{Oid: "3", Hostname: "Laboratory3", Ipv4: "192.168.178.15", MacAddress: "BD-00-4D-78-52-17", Ipv6: "e5d7:fff3:1035:1b70:5481:8362:9e05:e16c", OperatingSystem: "Linux NixOS", Ports: []entities.Port{{Number: 27017, Protocol: "TCP"}, {Number: 22, Protocol: "TCP"}}},
	{Oid: "4", Hostname: "Entrance1", Ipv4: "192.168.178.43", MacAddress: "CA-06-3E-08-76-36", Ipv6: "3183:7bf8:1273:574f:f3a5:5faf:b777:6857", OperatingSystem: "Windows Server 2017", Ports: []entities.Port{{Number: 445, Protocol: "TCP"}, {Number: 22, Protocol: "TCP"}}},
	{Oid: "5", Hostname: "Entrance2", Ipv4: "192.168.178.44", MacAddress: "E5-C1-29-4B-CB-71", Ipv6: "d6cc:e97e:dc3b:08f8:26b8:331d:7261:7d1b", OperatingSystem: "MacOS", Ports: []entities.Port{{Number: 443, Protocol: "TCP"}, {Number: 22, Protocol: "TCP"}}},
	{Oid: "6", Hostname: "Entrance3", Ipv4: "192.168.178.45", MacAddress: "6C-0E-8B-84-94-56", Ipv6: "c5e5:0a7f:a69f:b40e:f029:48fd:d2f3:4bd0", OperatingSystem: "Windows", Ports: []entities.Port{{Number: 80, Protocol: "TCP"}, {Number: 22, Protocol: "TCP"}}},
	{Oid: "7", Hostname: "Conference1", Ipv4: "192.168.178.16", MacAddress: "D9-AB-6C-5E-F2-8A", Ipv6: "a1d0:37f6:95a8:96ed:e7c8:ee91:01a3:ff47", OperatingSystem: "Linux Ubuntu", Ports: []entities.Port{{Number: 8080, Protocol: "TCP"}, {Number: 22, Protocol: "TCP"}}},
	{Oid: "8", Hostname: "Conference2", Ipv4: "192.168.178.17", MacAddress: "8F-3D-42-71-19-63", Ipv6: "4b49:a9c2:bf24:6a7e:d0d7:97ff:95a6:4be1", OperatingSystem: "Windows Server 2019", Ports: []entities.Port{{Number: 3389, Protocol: "TCP"}, {Number: 80, Protocol: "TCP"}}},
	{Oid: "9", Hostname: "Conference3", Ipv4: "192.168.178.18", MacAddress: "2C-7E-16-AB-ED-58", Ipv6: "f8d2:b150:3f3b:ec0a:6bf7:baae:15d8:94c7", OperatingSystem: "MacOS", Ports: []entities.Port{{Number: 443, Protocol: "TCP"}, {Number: 22, Protocol: "TCP"}}},
	{Oid: "10", Hostname: "ServerRoom1", Ipv4: "192.168.178.46", MacAddress: "76-F1-93-5D-28-72", Ipv6: "97a5:47e1:641a:06d4:64c9:3652:8459:1e5d", OperatingSystem: "Linux CentOS", Ports: []entities.Port{{Number: 3306, Protocol: "TCP"}, {Number: 5432, Protocol: "TCP"}}},
	{Oid: "11", Hostname: "ServerRoom2", Ipv4: "192.168.178.47", MacAddress: "A3-6D-8A-BC-7F-2E", Ipv6: "86f8:de72:ac8b:80d2:1e5a:9d9a:1e49:6f3d", OperatingSystem: "Windows Server 2016", Ports: []entities.Port{{Number: 445, Protocol: "TCP"}, {Number: 22, Protocol: "TCP"}}},
	{Oid: "12", Hostname: "ServerRoom3", Ipv4: "192.168.178.48", MacAddress: "E8-4B-0F-2A-CE-91", Ipv6: "b1a7:c2a8:cc05:a979:9fb3:01a0:7b3a:fc6d", OperatingSystem: "Linux Red Hat", Ports: []entities.Port{{Number: 27017, Protocol: "TCP"}, {Number: 80, Protocol: "TCP"}}},
}
