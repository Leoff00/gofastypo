package ui

import (
	"fyne.io/fyne/v2"
)

type ContainerComposer struct {
	objs *[]fyne.Container
}

func NewContainerComposer(comp ContainerComposer) (*ContainerComposer, error) {
	c := &ContainerComposer{
		objs: comp.objs,
	}

	return c, nil
}
