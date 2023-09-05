package RGB

import "fmt"

// This function enables us to print text in the color that we want
// The color must be written in RGB code
func RGB_Text(r, g, b int, s string) string {
	color := fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", r, g, b, s)
	return color
}
