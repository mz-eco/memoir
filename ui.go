package memoir

type UI interface {
	Component() Component
}

func UIToComponent(uiList ...UI) []Component {

	components := make([]Component, 0)

	for _, ui := range uiList {
		components = append(components, ui.Component())
	}

	return components

}
