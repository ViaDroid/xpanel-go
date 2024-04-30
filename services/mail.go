package services

type MailService struct{}

func NewMailService() *MailService {
	return &MailService{}
}

func (s *MailService) Send(email, subject, template string, bodyMap map[string]any) error {
	// TODO
	return nil
}
