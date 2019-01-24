package pmgo

import "github.com/globalsign/mgo"

type DatabaseManager interface {
	AddUser(username, password string, readOnly bool) error
	C(name string) CollectionManager
	CollectionNames() (names []string, err error)
	CreateView(view string, source string, pipeline interface{}, collation *mgo.Collation) error
	DropDatabase() error
	FindRef(ref *mgo.DBRef) *mgo.Query
	GridFS(prefix string) *mgo.GridFS
	Login(user, pass string) error
	Logout()
	RemoveUser(user string) error
	Run(cmd interface{}, result interface{}) error
	UpsertUser(user *mgo.User) error
	With(s SessionManager) DatabaseManager
}

type Database struct {
	db *mgo.Database
}

func NewDatabaseManager(d *mgo.Database) DatabaseManager {
	return &Database{
		db: d,
	}
}

func (d *Database) AddUser(username, password string, readOnly bool) error {
	return d.db.AddUser(username, password, readOnly)
}

func (d *Database) C(name string) CollectionManager {
	c := &Collection{
		collection: d.db.C(name),
	}
	return c
}

func (d *Database) CollectionNames() ([]string, error) {
	return d.db.CollectionNames()
}

func (d *Database) CreateView(view string, source string, pipeline interface{}, collation *mgo.Collation) error {
	return d.db.CreateView(view, source, pipeline, collation)
}

func (d *Database) DropDatabase() error {
	return d.db.DropDatabase()
}

func (d *Database) FindRef(ref *mgo.DBRef) *mgo.Query {
	return d.db.FindRef(ref)
}

func (d *Database) GridFS(prefix string) *mgo.GridFS {
	return d.db.GridFS(prefix)
}

func (d *Database) Run(cmd interface{}, result interface{}) error {
	return d.db.Run(cmd, result)
}

func (d *Database) Login(user, pass string) error {
	return d.db.Login(user, pass)
}

func (d *Database) Logout() {
	d.db.Logout()
}

func (d *Database) RemoveUser(user string) error {
	return d.db.RemoveUser(user)
}

func (d *Database) UpsertUser(user *mgo.User) error {
	return d.db.UpsertUser(user)
}

func (d *Database) With(s SessionManager) DatabaseManager {
	return &Database{
		db: d.db,
	}
}
