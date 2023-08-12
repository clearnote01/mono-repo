package renderer

import (
	"github.com/gc-cli/api"
	"github.com/gc-cli/models"
	"github.com/gc-cli/parser"
	"github.com/gdamore/tcell/v2"
	"github.com/repeale/fp-go"
	"github.com/rivo/tview"
)

func getTableGroups(data *models.GroupStats) *tview.Table {
	table := tview.NewTable().
		SetFixed(40, 40)

	cols, rows := 4, len(data.ErrorGroupStats)

	headers := [3]string{"GroupId", "ResolutionStatus", "Count"}

	headerColor := tcell.ColorDarkBlue
	for c := 0; c < len(headers); c++ {
		table.SetCell(0, c,
			tview.NewTableCell(headers[c]).
				SetBackgroundColor(tcell.ColorMintCream).
				SetTextColor(headerColor).
				SetAlign(tview.AlignCenter))
	}
	bodyColor := tcell.ColorWhite
	for r := 1; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if c == 0 {
				table.SetCell(r, c,
					tview.NewTableCell(data.ErrorGroupStats[r].Group.GroupId).
						SetTextColor(bodyColor).
						SetAlign(tview.AlignCenter))

			} else if c == 1 {
				table.SetCell(r, c,
					tview.NewTableCell(data.ErrorGroupStats[r].Group.ResolutionStatus).
						SetTextColor(bodyColor).
						SetAlign(tview.AlignCenter))

			} else if c == 2 {
				table.SetCell(r, c,
					tview.NewTableCell(data.ErrorGroupStats[r].Count).
						SetTextColor(bodyColor).
						SetAlign(tview.AlignCenter))

			} else if c == 3 {
				// table.SetCell(r, c,
				// 	tview.NewTableCell(data.ErrorGroupStats[r].Group.ResolutionStatus).
				// 		SetTextColor(bodyColor).
				// 		SetAlign(tview.AlignCenter))

			}
		}
	}
	return table
}

func RenderGroup() {
	// table := getTable()
	// RenderWrapper(table)
}

func RenderTree() {
	app := tview.NewApplication()
	rootDir := "."
	root := tview.NewTreeNode(rootDir).
		SetColor(tcell.ColorRed)
	tree := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)

	// A helper function which adds the files and directories of the given path
	// to the given target node.
	add := func(target *tview.TreeNode, path string, ref string) *tview.TreeNode {

		node := tview.NewTreeNode(path).
			SetReference(ref).
			SetSelectable(true)
		// if file.IsDir() {
		// 	node.SetColor(tcell.ColorGreen)
		// }
		target.AddChild(node)
		return node
	}
	groups := parser.GetErrorGroupsOpenStatus()
	for _, group := range groups.ErrorGroupStats {
		text := group.Group.ResolutionStatus + " - last_seen: " + group.LastSeenTime + " - " + group.Representative.ServiceContext.Service + " - " + group.Representative.Message
		node := add(root, text, group.Group.GroupId)
		summary := add(node, "Summary", group.Group.GroupId)
		affectedServices := fp.Reduce[models.AffectedService, string](func(x string, y models.AffectedService) string { return y.Service }, "")(group.AffectedServices)
		add(summary, "Affected Services: "+affectedServices, group.Group.GroupId)
		add(node, "List", group.Group.GroupId)

		// foo(group.AffectedServices)

		// add(node, "Affected Services: "+affectedServices, group.Group.GroupId+"_"+string(i))
	}

	// Add the current directory to the root node.
	add(root, rootDir, "")

	// If a directory was selected, open it.
	tree.
		SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEscape {
				app.Stop()
			}
		}).SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil || reference == "" {
			return // Selecting the root node does nothing.
		}
		if node.GetText() == "List" {
			children := node.GetChildren()
			if len(children) == 0 {
				// Load and show files in this directory.
				errors := api.GetErrorEventsAll(reference.(string), 0, 5)
				for _, err := range errors.ErrorEvent {
					add(node, err.ServiceContext.Service+" - "+err.EventTime+" - "+err.Message, "")
				}
			}
		}
		node.SetExpanded(!node.IsExpanded())
	})

	// ree.SetBackgroundColor(tcell.ColorAliceBlue)

	if err := app.SetRoot(tree, true).Run(); err != nil {
		panic(err)
	}
}
