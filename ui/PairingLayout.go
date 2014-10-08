package ui

import (
	"image"
	"image/color"

	"github.com/ninjasphere/go-ninja/api"
	"github.com/ninjasphere/go-ninja/logger"
)

type PairingLayout struct {
	currentPane Pane
	conn        *ninja.Connection
	log         *logger.Logger
}

func NewPairingLayout(conn *ninja.Connection) *PairingLayout {
	layout := &PairingLayout{
		log: logger.GetLogger("PaneLayout"),

		conn: conn,
	}

	return layout
}

func (l *PairingLayout) ShowColor(c color.Color) {
	l.currentPane = NewColorPane(c)
}

func (l *PairingLayout) ShowCode(text string) {
	l.currentPane = NewPairingCodePane(text)
}

func (l *PairingLayout) ShowIcon(image string) {
	l.currentPane = NewImagePane("./images/" + image)
}

func (l *PairingLayout) Render() (*image.RGBA, error) {
	if l.currentPane != nil {
		return l.currentPane.Render()
	}

	return &image.RGBA{}, nil
}
