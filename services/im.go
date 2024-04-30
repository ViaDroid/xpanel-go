package services

type IM struct{}

type imInterface interface {
	Send(to, msg string)
}

type Discord struct{}

type Slack struct{}

type Telegram struct{}

func (*Discord) Send(to, msg string) {
	//TODO
}

func (*Slack) Send(to, msg string) {
	//TODO
}

func (*Telegram) Send(to, msg string) {
	//TODO
}

func (IM) Send(to, msg string, tp int) {
	var im imInterface
	switch tp {
	case 1:
		im = &Discord{}
	case 2:
		im = &Slack{}
	default:
		im = &Telegram{}
	}
	im.Send(to, msg)
}
