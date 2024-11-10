package main

func main() {
	flamengoFans := Publisher{}
	fluminenseFans := Publisher{}
	subscriberFla1 := NewSubscription("Carol")
	subscriberFlu1 := NewSubscription("Jeff")
	subscriberFlu2 := NewSubscription("José")
	subscriberFla2 := NewSubscription("César")
	flamengoFans.AddListener(subscriberFla1)
	flamengoFans.AddListener(subscriberFla2)
	fluminenseFans.AddListener(subscriberFlu1)
	fluminenseFans.AddListener(subscriberFlu2)
	fluminenseFans.AddListener(subscriberFla1)
	fluminenseFans.NotifyAll(Event{Description: "Golllll do Fluminense: German Cano!"})
	flamengoFans.NotifyAll(Event{Description: "Golllll do Flamengo: Gabigol!"})
	flamengoFans.RemoveListener(subscriberFla1)
	flamengoFans.NotifyAll(Event{Description: "Gabigol expulso!"})
}
