package postgres

import (
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/qsliu2017/devdb"
)

var _ devdb.EmbeddedDatabase = (*Postgres)(nil)

type Postgres struct {
	pg       *embeddedpostgres.EmbeddedPostgres
	username string
	password string
	port     uint32
}

func (p *Postgres) Password() string {
	return p.password
}

func (p *Postgres) Port() uint32 {
	return p.port
}

func (p *Postgres) Username() string {
	return p.username
}

func New() *Postgres {
	return &Postgres{}
}

func (p *Postgres) Start() error {
	cfg := embeddedpostgres.DefaultConfig().Version(embeddedpostgres.V13).Username(p.username).Password(p.password).Database("postgres").Port(p.port)
	p.pg = embeddedpostgres.NewDatabase(cfg)
	return p.pg.Start()
}

func (p *Postgres) Stop() error {
	return p.pg.Stop()
}
