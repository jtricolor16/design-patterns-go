package main

type History struct {
	memento MementoInterface
}

func (h *History) Backup(route *Route, oldName string) {
	h.memento = route.Save(oldName)
}

func (h *History) Undo() {
	if h.memento != nil {
		h.memento.Restore()
	}
}
