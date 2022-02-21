package file

import "strings"

func getTextExtensions() []string {
	return []string{".txt", ".php", ".go", ".js", ".css", ".log", ".ini", ".xml", ".gitignore"}
}

func getPictureExtensions() []string {
	return []string{".png", ".jpg", ".jpeg", ".gif", ".webp"}
}

func containExt(ext string, allExt []string) bool {
	ext = strings.ToLower(ext)

	for _, e := range allExt {
		if e == ext {
			return true
		}
	}

	return false
}
