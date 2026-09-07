package fixture

import "fmt"

func RenderHeader(settings map[string]any) string {
	return fmt.Sprintf("%v (%v)", settings["title"], settings["region"])
}

func RenderFooter(settings map[string]any) string {
	return fmt.Sprintf("Generated for %v", settings["owner"])
}
