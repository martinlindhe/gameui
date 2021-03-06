package ui

import "log"

// CheckUI performs some sanity checks on the UI instance in order to detect programmatic errors
// returns true if checks passed
func CheckUI(ui *UI) bool {
	for _, c := range ui.children {
		if btn, ok := c.(*Button); ok {
			if btn.Image == nil {
				continue
			}
			allB := btn.Image.Bounds()
			if allB.Max.X > btn.Dimension.Width || allB.Max.Y > btn.Dimension.Height {
				log.Println("warning: button.drawImage image is bigger than container button")
			}
		}
		if txt, ok := c.(*Text); ok {
			if txt.font.size < 3 {
				log.Println("warning: text size too small:", txt.font.size)
			}
		}
	}
	return true
}
