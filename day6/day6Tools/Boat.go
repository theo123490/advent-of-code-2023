package day6Tools

type Boat struct {
	timePressed int
	speed       int
	distance    int
}

func NewBoat() Boat {
	return Boat{0, 0, 0}
}

func (b *Boat) calculateDistance(totalTime int) {
	var remainingTime int = totalTime - b.timePressed
	b.speed = b.timePressed
	b.distance = b.speed * remainingTime
}
func (b *Boat) isWinRace(race Race) bool {
	b.calculateDistance(race.time)
	return b.distance > race.distance
}
