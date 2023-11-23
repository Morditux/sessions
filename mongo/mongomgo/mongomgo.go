package mongomgo

import (
	"github.com/Morditux/mongostore"
	"github.com/Morditux/sessions"
	"github.com/globalsign/mgo"
)

var _ sessions.Store = (*store)(nil)

func NewStore(c *mgo.Collection, maxAge int, ensureTTL bool, keyPairs ...[]byte) sessions.Store {
	return &store{mongostore.NewMongoStore(c, maxAge, ensureTTL, keyPairs...)}
}

type store struct {
	*mongostore.MongoStore
}

func (c *store) Options(options sessions.Options) {
	c.MongoStore.Options = options.ToGorillaOptions()
}
