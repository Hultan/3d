package three_d

import (
	_ "embed"
	"math"
	"os"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

const applicationTitle = "3d"
const applicationVersion = "v 0.01"
const applicationCopyRight = "©SoftTeam AB, 2020"

type MainForm struct {
	window      *gtk.ApplicationWindow
	builder     *gtkBuilder
	aboutDialog *gtk.AboutDialog
	extraForm   *gtk.Window
	dialog      *gtk.Dialog
}

//go:embed assets/main.glade
var gladeFile string

//go:embed assets/application.png
var applicationIcon []byte

// NewMainForm : Creates a new MainForm object
func NewMainForm() *MainForm {
	mainForm := new(MainForm)
	return mainForm
}

// OpenMainForm : Opens the MainForm window
func (m *MainForm) OpenMainForm(app *gtk.Application) {
	// Initialize gtk
	gtk.Init(&os.Args)

	// Create a new softBuilder
	builder, err := newBuilder(gladeFile)
	if err != nil {
		panic(err)
	}
	m.builder = builder

	// Get the main window from the glade file
	m.window = m.builder.getObject("main_window").(*gtk.ApplicationWindow)

	// Set up main window
	m.window.SetApplication(app)
	m.window.SetTitle("3d main window")

	// Hook up the destroy event
	m.window.Connect("destroy", m.window.Destroy)

	// Quit button
	button := m.builder.getObject("main_window_quit_button").(*gtk.ToolButton)
	button.Connect("clicked", m.window.Destroy)

	// Status bar
	statusBar := m.builder.getObject("main_window_status_bar").(*gtk.Statusbar)
	statusBar.Push(statusBar.GetContextId("3d"), "3d : version 0.1.0")

	// Menu
	m.setupMenu()

	da := m.builder.getObject("drawingArea").(*gtk.DrawingArea)
	da.Connect("draw", m.onDraw)

	// Show the main window
	m.window.ShowAll()

	go func() {
		for {
			da.QueueDraw()
		}
	}()
}

func (m *MainForm) setupMenu() {
	menuQuit := m.builder.getObject("menu_file_quit").(*gtk.MenuItem)
	menuQuit.Connect("activate", m.window.Destroy)

	menuHelpAbout := m.builder.getObject("menu_help_about").(*gtk.MenuItem)
	menuHelpAbout.Connect(
		"activate", func() {
			m.openAboutDialog()
		},
	)
}

var width, height float64
var fps = 0
var points3d []Vector3

func (m *MainForm) onDraw(da *gtk.DrawingArea, ctx *cairo.Context) {
	fps++
	height, width = float64(da.GetAllocatedHeight()), float64(da.GetAllocatedWidth())
	m.drawBackground(ctx)

	// The screen is at z=1.0, so we need to place it somewhere behind that
	dz := math.Sin(float64(fps)/60.0) * 0.5

	if len(points3d) == 0 {
		m.createCube()
	}

	var points2d []Vector2
	for _, vector3 := range points3d {
		// For rotating around y-axis:
		// 	Add cos(fps) to x
		//	Add sin(fps) to z
		vector3.Z += dz
		points2d = append(points2d, projectToScreen(projectTo2d(vector3)))
	}

	m.drawCube(ctx, points2d)
}

func (m *MainForm) drawBackground(ctx *cairo.Context) {
	ctx.SetSourceRGB(0.8, 0.8, 0.8)
	ctx.Paint()
}

func (m *MainForm) drawLine(ctx *cairo.Context, p1, p2 Vector2) {
	ctx.SetSourceRGB(1, 0, 0)

	ctx.MoveTo(p1.X, p1.Y)
	ctx.LineTo(p2.X, p2.Y)
	ctx.Stroke()
}

func (m *MainForm) createCube() {
	z := 1.5

	points3d = append(points3d, Vector3{-0.5, -0.5, z})
	points3d = append(points3d, Vector3{0.5, -0.5, z})
	points3d = append(points3d, Vector3{0.5, 0.5, z})
	points3d = append(points3d, Vector3{-0.5, 0.5, z})
	points3d = append(points3d, Vector3{-0.5, -0.5, z + 1})
	points3d = append(points3d, Vector3{0.5, -0.5, z + 1})
	points3d = append(points3d, Vector3{0.5, 0.5, z + 1})
	points3d = append(points3d, Vector3{-0.5, 0.5, z + 1})
}

func (m *MainForm) drawCube(ctx *cairo.Context, points2d []Vector2) {
	m.drawLine(ctx, points2d[0], points2d[1])
	m.drawLine(ctx, points2d[1], points2d[2])
	m.drawLine(ctx, points2d[2], points2d[3])
	m.drawLine(ctx, points2d[3], points2d[0])

	m.drawLine(ctx, points2d[4], points2d[5])
	m.drawLine(ctx, points2d[5], points2d[6])
	m.drawLine(ctx, points2d[6], points2d[7])
	m.drawLine(ctx, points2d[7], points2d[4])

	m.drawLine(ctx, points2d[0], points2d[4])
	m.drawLine(ctx, points2d[1], points2d[5])
	m.drawLine(ctx, points2d[2], points2d[6])
	m.drawLine(ctx, points2d[3], points2d[7])
}
