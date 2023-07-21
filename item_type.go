package main

import (
	"encoding/json"
	"strconv"
)

// https://www.zabbix.com/documentation/current/en/manual/api/reference/item/object

type ItemType int

const (
	ItemTypeAgent ItemType = iota	// 0
	_
	ItemTypeTrapper			// 2
	ItemTypeSimpleCheck		// 3
	_
	ItemTypeInternal		// 5
	_
	ItemTypeAgentActive		// 7
	_
	ItemTypeWeb			// 9
	ItemTypeExternalCheck		// 10
	ItemTypeDatabaseMonitor		// 11
	ItemTypeIPMI			// 12
	ItemTypeSSH			// 13
	ItemTypeTELNET			// 14
	ItemTypeCalculated		// 15
	ItemTypeJMX			// 16
	ItemTypeSNMPTrap		// 17
	ItemTypeDependent		// 18
	ItemTypeHTTPAgent		// 19
	ItemTypeSNMPAgent		// 20
	Script				// 21
)

func (it *ItemType) UnmarshalJSON(data []byte) error {
	var stringValue string

	err := json.Unmarshal(data, &stringValue)
	if err != nil {
		return err
	}

	intValue, err := strconv.ParseInt(stringValue, 10, 64)
	if err != nil {
		return err
	}

	*it = ItemType(intValue)

	return nil
}

func (it ItemType) String() string {
	switch it {
	case ItemTypeAgent:
		return "agent"
	case ItemTypeTrapper:
		return "trapper"
	case ItemTypeSimpleCheck:
		return "check"
	case ItemTypeInternal:
		return "internal"
	case ItemTypeAgentActive:
		return "active"
	case ItemTypeWeb:
		return "web"
	case ItemTypeExternalCheck:
		return "external"
	case ItemTypeDatabaseMonitor:
		return "dbmon"
	case ItemTypeIPMI:
		return "ipmi"
	case ItemTypeSSH:
		return "ssh"
	case ItemTypeTELNET:
		return "telnet"
	case ItemTypeCalculated:
		return "calc"
	case ItemTypeJMX:
		return "jmx"
	case ItemTypeSNMPTrap:
		return "snmptrap"
	case ItemTypeDependent:
		return "dependent"
	case ItemTypeHTTPAgent:
		return "httpagent"
	case ItemTypeSNMPAgent:
		return "snmpagent"
	default:
		return "unknown"
	}
}
