package ninerouter

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"strings"
)

const cliTokenSalt = "9r-cli-auth"

func deriveCLIToken() string {
	machineID := readMachineID()
	if machineID == "" {
		return ""
	}
	machineHash := sha256Hex(machineID)
	return sha256Hex(machineHash + cliTokenSalt)[:16]
}

func readMachineID() string {
	for _, path := range []string{"/etc/machine-id", "/var/lib/dbus/machine-id"} {
		data, err := os.ReadFile(path)
		if err == nil {
			if id := strings.TrimSpace(string(data)); id != "" {
				return id
			}
		}
	}
	return ""
}

func sha256Hex(value string) string {
	sum := sha256.Sum256([]byte(value))
	return hex.EncodeToString(sum[:])
}
