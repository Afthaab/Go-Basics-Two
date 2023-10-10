package db

type Conn struct {
	db string
}

func NewConn(db string) (Conn, bool) {
	if db == "" {
		return Conn{}, false
	}
	return Conn{
		db: db,
	}, true
}
