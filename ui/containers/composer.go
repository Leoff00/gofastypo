package containers

import "fyne.io/fyne/v2"

type Components struct {
	Header   *fyne.Container
	Textarea *fyne.Container
	Meter    *fyne.Container
}

func Composer(header, meter, textarea *fyne.Container) *Components {
	return &Components{
		Header:   header,
		Meter:    meter,
		Textarea: textarea,
	}
}
