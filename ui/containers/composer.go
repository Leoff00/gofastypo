package containers

import "fyne.io/fyne/v2"

type Components struct {
	Header   *fyne.Container
	Textarea *fyne.Container
}

func Composer(header, textarea *fyne.Container) *Components {
	return &Components{
		Header:   header,
		Textarea: textarea,
	}
}
