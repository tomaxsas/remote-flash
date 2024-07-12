package main

import (
	"log"
	"net"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

type (
	// C quick alias for Context.
	C = layout.Context
	// D quick alias for Dimensions.
	D = layout.Dimensions
)

func main() {
	go func() {
		window := new(app.Window)
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	th := material.NewTheme()
	window.Option(
		app.Title("Remote flash helper"),
		app.Size(unit.Dp(280), unit.Dp(270)))
	var ops op.Ops
	var ipInputField component.TextField
	var getCarInfoBtn widget.Clickable
	var startProxyBtn widget.Clickable
	var statusBar material.LabelStyle
	var connectedTocar = false
	var startedProxy = false
	var labelText = ""
	statusBar = material.Caption(th, labelText)
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			if getCarInfoBtn.Clicked(gtx) {
				ip := ipInputField.Text()
				if net.ParseIP(ip) == nil {
					labelText = "Invalid IP address"
					continue
				}
				err := getCarInfo(ip)
				if err != nil {
					labelText = err.Error()
					continue
				}
				connectedTocar = true
				labelText = "Connected. Click start"
			}
			if startProxyBtn.Clicked(gtx) {
				// start proxy
				if !startedProxy {
					startProxy(ipInputField.Text())
					startedProxy = true
					log.Println("Starting proxy clicked")
				}
				labelText = "Server started. Proceed with flashing."
			}

			layout.Center.Layout(gtx, func(gtx C) D {
				gtx.Constraints.Max.X = gtx.Dp(unit.Dp(300))
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx C) D {
						return ipInputField.Layout(gtx, th, "Enter ZGW IP address")
					}),
					layout.Rigid(func(gtx C) D {
						return layout.Spacer{Height: unit.Dp(10)}.Layout(gtx)
					}),
					layout.Rigid(func(gtx C) D {
						return material.Button(th, &getCarInfoBtn, "Connect to car").Layout(gtx)
					}),
					layout.Rigid(func(gtx C) D {
						return layout.Spacer{Height: unit.Dp(10)}.Layout(gtx)
					}),
					layout.Rigid(func(gtx C) D {
						if connectedTocar {
							return material.Button(th, &startProxyBtn, "START").Layout(gtx)
						} else {

							return material.Button(th, &startProxyBtn, "START").Layout(gtx.Disabled())
						}
					}),
					layout.Rigid(func(gtx C) D {
						return layout.Spacer{Height: unit.Dp(10)}.Layout(gtx)
					}),
					layout.Rigid(func(gtx C) D {
						statusBar.Text = labelText
						return statusBar.Layout(gtx)
					}),
				)
			})
			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
		}
	}
}
