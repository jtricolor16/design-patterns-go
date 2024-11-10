package main

func main() {
	r1 := NewRoute("left")
	r2 := NewRoute("up")

	control := History{}
	control.Backup(r2, r1.name)
	control.memento.PrintRouteName()
	control.Undo()
	control.memento.PrintRouteName()
	r2.SetName("down")
	r1.SetName("right")
	control.Backup(r2, r1.name)
	control.memento.PrintRouteName()
	control.Undo()
	control.memento.PrintRouteName()
}
