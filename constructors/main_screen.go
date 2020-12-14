package constructor

import (
	"clx/screen"
	"clx/types"
	"github.com/gdamore/tcell/v2"
	"gitlab.com/tslocum/cview"
)

const (
	maximumStoriesToDisplay = 30
	helpPage                = "help"
	offlinePage             = "offline"
)

func NewScreenController() *types.ScreenController {
	sc := new(types.ScreenController)
	sc.Application = cview.NewApplication()

	sc.Submissions = []*types.Submissions{}
	sc.Submissions = append(sc.Submissions, new(types.Submissions))
	sc.Submissions = append(sc.Submissions, new(types.Submissions))
	sc.Submissions = append(sc.Submissions, new(types.Submissions))
	sc.Submissions = append(sc.Submissions, new(types.Submissions))

	sc.ApplicationState = new(types.ApplicationState)
	sc.ApplicationState.ScreenWidth = screen.GetTerminalWidth()
	sc.ApplicationState.ScreenHeight = screen.GetTerminalHeight()
	sc.ApplicationState.ViewableStoriesOnSinglePage = screen.GetViewableStoriesOnSinglePage(
		sc.ApplicationState.ScreenHeight,
		maximumStoriesToDisplay)

	sc.Submissions[types.FrontPage].MaxPages = 2
	sc.Submissions[types.New].MaxPages = 2
	sc.Submissions[types.Ask].MaxPages = 1
	sc.Submissions[types.Show].MaxPages = 1

	sc.MainView = NewMainView()

	newsList := NewList()
	sc.MainView.Panels.AddPanel(types.FrontPagePanel, newsList, true, false)
	sc.MainView.Panels.AddPanel(types.NewestPanel, NewList(), true, false)
	sc.MainView.Panels.AddPanel(types.ShowPanel, NewList(), true, false)
	sc.MainView.Panels.AddPanel(types.AskPanel, NewList(), true, false)

	sc.MainView.Panels.SetCurrentPanel(types.FrontPagePanel)

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

func NewMainView() *types.MainView {
	main := new(types.MainView)
	main.Panels = cview.NewPanels()
	main.Grid = cview.NewGrid()
	main.LeftMargin = newTextViewPrimitive("")
	main.LeftMargin.SetTextAlign(cview.AlignRight)
	main.RightMargin = newTextViewPrimitive("")
	main.Header = newTextViewPrimitive("")
	main.Footer = newTextViewPrimitive("")

	main.Grid.SetBorder(false)
	main.Grid.SetRows(2, 0, 1)
	main.Grid.SetColumns(7, 0, 3)
	main.Grid.SetBackgroundColor(tcell.ColorDefault)
	main.Grid.AddItem(main.Header, 0, 0, 1, 3, 0, 0, false)
	main.Grid.AddItem(main.Footer, 2, 0, 1, 3, 0, 0, false)
	main.Grid.AddItem(main.LeftMargin, 1, 0, 1, 1, 0, 0, false)
	main.Grid.AddItem(main.Panels, 1, 1, 1, 1, 0, 0, true)
	main.Grid.AddItem(main.RightMargin, 1, 2, 1, 1, 0, 0, false)

	main.Panels.AddPanel(helpPage, GetHelpScreen(), true, false)
	main.Panels.AddPanel(offlinePage, GetOfflineScreen(), true, false)

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
