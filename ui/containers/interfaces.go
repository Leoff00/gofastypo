package containers

import "fyne.io/fyne/v2"

type FyneContainers interface {
	HeaderContainer() *fyne.Container
	TextAreaContainer() *fyne.Container
	WpmContainer() *fyne.Container
	ExitButtonContainer() *fyne.Container
}
