package constructor

import (
	"clx/constants"
	"clx/screen"
	"clx/structs"
	"github.com/gdamore/tcell/v2"
	"gitlab.com/tslocum/cview"
)

const (
	maximumStoriesToDisplay = 30
)

func NewScreenController() *structs.ScreenController {
	sc := new(structs.ScreenController)
	sc.Application = cview.NewApplication()

	sc.Submissions = []*structs.Submissions{}
	sc.Submissions = append(sc.Submissions, new(structs.Submissions))
	sc.Submissions = append(sc.Submissions, new(structs.Submissions))
	sc.Submissions = append(sc.Submissions, new(structs.Submissions))
	sc.Submissions = append(sc.Submissions, new(structs.Submissions))

	sc.ApplicationState = new(structs.ApplicationState)
	sc.ApplicationState.ScreenWidth = screen.GetTerminalWidth()
	sc.ApplicationState.ScreenHeight = screen.GetTerminalHeight()
	sc.ApplicationState.SubmissionsToShow = screen.GetSubmissionsToShow(
		sc.ApplicationState.ScreenHeight,
		maximumStoriesToDisplay)

	sc.Submissions[constants.FrontPage].MaxPages = 2
	sc.Submissions[constants.New].MaxPages = 2
	sc.Submissions[constants.Ask].MaxPages = 1
	sc.Submissions[constants.Show].MaxPages = 1

	sc.List = NewList()
	sc.MainView = NewMainView()
	sc.MainView.Panels.AddPanel(constants.SubmissionsPanel, sc.List, true, true)

	return sc
}

func NewList() *cview.List {
	list := cview.NewList()
	list.SetBackgroundTransparent(false)
	list.SetBackgroundColor(tcell.ColorDefault)
	list.SetMainTextColor(tcell.ColorDefault)
	list.SetSecondaryTextColor(tcell.ColorDefault)
	list.SetSelectedTextAttributes(tcell.AttrReverse)
	list.SetSelectedTextColor(tcell.ColorDefault)
	list.SetSelectedBackgroundColor(tcell.ColorDefault)
	list.SetScrollBarVisibility(cview.ScrollBarNever)

	return list
}

func NewMainView() *structs.MainView {
	main := new(structs.MainView)
	main.Panels = cview.NewPanels()
	main.Grid = cview.NewGrid()
	main.LeftMargin = newTextViewPrimitive("")
	main.LeftMargin.SetTextAlign(cview.AlignRight)
	main.Header = newTextViewPrimitive("")
	main.PageCounter = newTextViewPrimitive("")
	main.StatusBar = newTextViewPrimitive("")
	main.StatusBar.SetTextAlign(cview.AlignCenter)
	main.StatusBar.SetPadding(0,0,-4,0)

	main.Grid.SetBorder(false)
	main.Grid.SetRows(2, 0, 1)
	main.Grid.SetColumns(7, 0, 4)
	main.Grid.SetBackgroundColor(tcell.ColorDefault)
	main.Grid.AddItem(main.Header, 0, 0, 1, 3, 0, 0, true)
	main.Grid.AddItem(main.LeftMargin, 1, 0, 1, 1, 0, 0, false)
	main.Grid.AddItem(main.Panels, 1, 1, 1, 2, 0, 0, false)
	main.Grid.AddItem(main.StatusBar, 2, 1, 1, 1, 0, 0, false)
	main.Grid.AddItem(main.PageCounter, 2, 2, 1, 1, 0, 0, false)

	main.Panels.AddPanel(constants.InfoPanel, GetInfoScreen(), true, false)
	main.Panels.AddPanel(constants.KeymapsPanel, GetHelpScreen(), true, false)
	main.Panels.AddPanel(constants.EnvironmentPanel, GetEnvironmentScreen(), true, false)


	return main
}

func newTextViewPrimitive(text string) *cview.TextView {
	tv := cview.NewTextView()
	tv.SetTextAlign(cview.AlignLeft)
	tv.SetText(text)
	tv.SetBorder(false)
	tv.SetBackgroundColor(tcell.ColorDefault)
	tv.SetTextColor(tcell.ColorDefault)
	tv.SetDynamicColors(true)
	tv.SetScrollBarVisibility(cview.ScrollBarNever)

	return tv
}
