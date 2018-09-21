package formatter

import (
	"encoding/json"
	"fmt"
	"strings"
)

// WallettitleFormatter ...
func WalletTitleFormatter(actionType string) string {
	switch actionType {
	case "payment":
		return "Pembayaran"
	case "topup":
		return "Top Up Wallet"
	case "transferIn":
	case "receive":
		return "Saldo Masuk"
	case "transferOut":
	case "send":
		return "Saldo Keluar"
	case "withdraw":
		return "Penarikan Dana"
	case "refund":
		return "Refund dari Payfazz"
	}
	return actionType
}

// WalletDetailToDescription ...
func WalletDetailToDescription(detail string) string {
	temp := map[string]string{}
	json.Unmarshal([]byte(detail), &temp)
	resp := ""
	for k, v := range temp {
		resp = fmt.Sprintf("%s%s: %s ", resp, strings.Title(k), v)
	}
	return resp
}
