package ui

import "log"

// CheckUI performs some sanity checks on the UI instance in order to detect programmatic errors
// returns true if checks passed
func CheckUI(ui *UI) bool {
	for _, c := range ui.components {
		if btn, ok := c.(*Button); ok {
			if btn.Image == nil {
				continue
			}
			allB := btn.Image.Bounds()
			if allB.Max.X > btn.Dimension.Width || allB.Max.Y > btn.Dimension.Height {
				log.Println("CheckUI warning: button.drawImage image is bigger than container button")
			}
		}
		if txt, ok := c.(*Text); ok {
			if txt.size < 3 {
				log.Println("CheckUI warning: text size too small:", txt.size)
			}
		}
	}
	return true
}
