package file

import (
	"fmt"
	"github.com/victorlss/flaky-api/pkg/shared/types"
	"path/filepath"
	"regexp"
)

// GenerateFilename returns a filename with format "id-[NNN]-[address].[ext]"
func GenerateFilename(house types.House) string {
	id := fmt.Sprintf("%03d", house.Id)
	address := formatAddress(house.Address)
	ext := filepath.Ext(house.PhotoURL)

	return fmt.Sprintf("id-%s-%s%s", id, address, ext)
}

// formatAddress formats an address removing period, comma and space.
func formatAddress(address string) string {
	reg := regexp.MustCompile("[,. ]+")
	return reg.ReplaceAllString(address, "_")
}
