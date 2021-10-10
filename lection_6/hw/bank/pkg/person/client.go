package person

type client struct {
	ID        int
	MobNumber string
	FIO       string
}

type Client interface {
	//TODO описать что может делать клиента банка
}
