//go:build !nogui && !nosplash

package splash

import (
	_ "image/png"

	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Style int

const (
	Compact Style = iota
	Familiar
)

func (s Style) Size() (w, h unit.Dp) {
	switch s {
	case Compact:
		w = unit.Dp(448)
		h = unit.Dp(150)
	case Familiar:
		w = unit.Dp(480)
		h = unit.Dp(240)
	}
	return
}

func (ui *Splash) drawCompact(gtx C) D {
	return layout.UniformInset(16).Layout(gtx, func(gtx C) D {
		return layout.Flex{
			Axis:      layout.Vertical,
		}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				return layout.Flex{
					Axis:      layout.Horizontal,
					Alignment: layout.Start,
				}.Layout(gtx,
					layout.Rigid(widget.Image{Src: paint.NewImageOp(ui.logo)}.Layout),
					layout.Rigid(func(gtx C) D {
						return layout.Flex{
							Axis:      layout.Vertical,
							Alignment: layout.Start,
						}.Layout(gtx,
							layout.Rigid(material.Label(ui.Theme, unit.Sp(16), ui.message).Layout),
							layout.Rigid(layout.Spacer{Height: unit.Dp(2)}.Layout),
							layout.Rigid(func(gtx C) D {
								return ui.drawDesc(gtx)
							}),
							layout.Rigid(layout.Spacer{Height: unit.Dp(20)}.Layout),
							layout.Rigid(func(gtx C) D {
								pb := ProgressBar(ui.Theme, ui.progress)
								pb.TrackColor = rgb(ui.Config.Gray1)
								return pb.Layout(gtx)
							}),
						)
					}),
				)
			}),
			layout.Rigid(layout.Spacer{Height: unit.Dp(16)}.Layout),
			layout.Rigid(func(gtx C) D {
				return ui.buttons(gtx, layout.SpaceStart)
			}),
		)
	})
}

func (ui *Splash) drawFamiliar(gtx C) D {
	return layout.Center.Layout(gtx, func(gtx C) D {
		return layout.Flex{
			Axis:      layout.Vertical,
			Alignment: layout.Middle,
		}.Layout(gtx,
			layout.Rigid(widget.Image{Src: paint.NewImageOp(ui.logo)}.Layout),
			layout.Rigid(func(gtx C) D {
				return layout.Flex{
					Axis:      layout.Vertical,
					Alignment: layout.Middle,
				}.Layout(gtx,
					layout.Rigid(layout.Spacer{Height: unit.Dp(8)}.Layout),
					layout.Rigid(material.Label(ui.Theme, unit.Sp(16), ui.message).Layout),
					layout.Rigid(func(gtx C) D {
						return layout.Inset{
							Top:    unit.Dp(14),
							Bottom: unit.Dp(14),
							Left:   unit.Dp(40),
							Right:  unit.Dp(40),
						}.Layout(gtx, func(gtx C) D {
							pb := ProgressBar(ui.Theme, ui.progress)
							pb.TrackColor = rgb(ui.Config.Gray1)
							return pb.Layout(gtx)
						})
					}),
					layout.Rigid(func(gtx C) D {
						return ui.drawDesc(gtx)
					}),
				)
			}),
			layout.Rigid(func(gtx C) D {
				return ui.buttons(gtx, layout.SpaceAround)
			}),
		)
	})
}