package pmgo

import "github.com/globalsign/mgo"

// CollectionManager is an interface for mgo.Collection struct.
// All implemented methods returns interfaces when needed
type CollectionManager interface {
	Count() (int, error)
	Create(*mgo.CollectionInfo) error
	DropCollection() error
	DropIndex(key ...string) error
	DropIndexName(name string) error
	EnsureIndex(index mgo.Index) error
	EnsureIndexKey(key ...string) error
	Find(interface{}) QueryManager
	FindId(id interface{}) QueryManager
	Indexes() (indexes []mgo.Index, err error)
	Insert(docs ...interface{}) error
	//NewIter(session *mgo.Session, firstBatch []bson.Raw, cursorId int64, err error) *mgo.Iter
	Pipe(interface{}) PipeManager
	Remove(selector interface{}) error
	RemoveAll(selector interface{}) (info *mgo.ChangeInfo, err error)
	RemoveId(id interface{}) error
	//Repair() *mgo.Iter
	Update(selector interface{}, update interface{}) error
	UpdateAll(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error)
	UpdateId(id interface{}, update interface{}) error
	Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error)
	UpsertId(id interface{}, update interface{}) (info *mgo.ChangeInfo, err error)
	//With(s *mgo.Session) *CollectionManager
}

type Collection struct {
	collection *mgo.Collection
}

func NewCollectionManager(c *mgo.Collection) CollectionManager {
	return &Collection{
		collection: c,
	}
}

func (c *Collection) Count() (int, error) {
	return c.collection.Count()
}

func (c *Collection) Create(info *mgo.CollectionInfo) error {
	return c.collection.Create(info)
}

func (c *Collection) DropCollection() error {
	return c.collection.DropCollection()
}

func (c *Collection) DropIndex(key ...string) error {
	return c.collection.DropIndex(key...)
}

func (c *Collection) DropIndexName(name string) error {
	return c.collection.DropIndexName(name)
}

func (c *Collection) EnsureIndex(index mgo.Index) error {
	return c.collection.EnsureIndex(index)
}

func (c *Collection) EnsureIndexKey(key ...string) error {
	return c.collection.EnsureIndexKey(key...)
}

func (c *Collection) Find(qu interface{}) QueryManager {
	return &Query{
		query: c.collection.Find(qu),
	}
}

func (c *Collection) Indexes() (indexes []mgo.Index, err error) {
	return c.collection.Indexes()
}

func (c *Collection) FindId(id interface{}) QueryManager {
	return &Query{
		query: c.collection.FindId(id),
	}
}

func (c *Collection) Insert(docs ...interface{}) error {
	return c.collection.Insert(docs...)
}

func (c *Collection) Pipe(query interface{}) PipeManager {
	return &Pipe{
		pipe: c.collection.Pipe(query),
	}
}

func (c *Collection) Remove(selector interface{}) error {
	return c.collection.Remove(selector)
}

func (c *Collection) RemoveAll(selector interface{}) (info *mgo.ChangeInfo, err error) {
	return c.collection.RemoveAll(selector)
}

func (c *Collection) RemoveId(id interface{}) error {
	return c.collection.RemoveId(id)
}

func (c *Collection) Update(selector interface{}, update interface{}) error {
	return c.collection.Update(selector, update)
}

func (c *Collection) UpdateAll(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	return c.collection.UpdateAll(selector, update)
}

func (c *Collection) UpdateId(id interface{}, update interface{}) error {
	return c.collection.UpdateId(id, update)
}

func (c *Collection) Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	return c.collection.Upsert(selector, update)
}

func (c *Collection) UpsertId(id interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	return c.collection.UpsertId(id, update)
}
