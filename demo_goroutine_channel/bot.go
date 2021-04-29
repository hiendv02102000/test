package demo_goroutine_channel

type Bot struct {
	Health int
	Damage int
}

func (b *Bot) Attack(b1 *Bot) {
	b1.Health -= b.Damage
}
