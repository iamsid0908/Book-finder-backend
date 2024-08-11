package domain

type HealthDomain interface {
	GetHealth() (string, error)
}

type HealthDomainCtx struct{}

func (c *HealthDomainCtx) GetHealth() (string, error) {
	return "The service is runningggg", nil
}
