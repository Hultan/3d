package three_d

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func (m *MainForm) openAboutDialog() {
	if m.aboutDialog == nil {
		m.aboutDialog = m.builder.getObject("about_dialog").(*gtk.AboutDialog)

		m.aboutDialog.SetModal(true)
		m.aboutDialog.SetTransientFor(m.window)
		m.aboutDialog.SetPosition(gtk.WIN_POS_CENTER_ON_PARENT)

		// Hook up the destroy event
		m.aboutDialog.Connect(
			"destroy", func() {
				m.aboutDialog.Destroy()
				m.aboutDialog = nil
			},
		)

		// About application
		m.aboutDialog.SetProgramName(applicationTitle)
		m.aboutDialog.SetComments("An application...")
		m.aboutDialog.SetVersion(applicationVersion)
		m.aboutDialog.SetCopyright(applicationCopyRight)
		image, err := gdk.PixbufNewFromBytesOnly(applicationIcon)
		if err == nil {
			m.aboutDialog.SetLogo(image)
		}
	}

	m.aboutDialog.Run()
	m.aboutDialog.Hide()
}
