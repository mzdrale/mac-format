// Package macformat contains functions to format MAC address
// Supported formats:
// xx:xx:xx:xx:xx:xx
// xx-xx-xx-xx-xx-xx
// xxxx.xxxx.xxxx

package macformat

import (
	"fmt"
	"regexp"
	"strings"
)

// Validate MAC address format
func ValidateMAC(mac string) bool {
	mac = strings.TrimSpace(mac)
	re, _ := regexp.Compile(`^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2}$`)
	if re.MatchString(mac) {
		return true
	}
	return false
}

// Convert MAC address into different format
func FormatMAC(mac string, format string) string {
	mac = strings.TrimSpace(mac)
	format = strings.TrimSpace(format)
}
