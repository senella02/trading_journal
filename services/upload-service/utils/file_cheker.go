package utils

import (
	"strings"
)

var allowedExtensions = []string{".jpg", ".jpeg", ".png"}

func IsValidImageExtension(currentExtension string) bool {
	for _, ext := range allowedExtensions {
		if  strings.EqualFold(ext, currentExtension) { //case insensitive comparison
			return true
		}
	}
	return false
}