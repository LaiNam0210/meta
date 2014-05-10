package meta

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"encoding/json"
	_ "errors"
	_ "log"
	"reflect"
)

type Meta struct {
	Name   string `bson:"name"`
	Fields []Attr `bson:"fields"`
}

type MetaHash struct {
	index map[string]Attr
}

type Attr struct {
	Label    string        `bson:"label"`
	Name     string        `bson:"name"`
	Required bool          `bson:"required"`
	Short    bool          `bson:"short"`
	Options  []interface{} `bson:"options"`
	Search   bool          `bson:search`
	Template string        `bson:"template"`
	Type     []interface{} `bson:"type"`
}

type Collection struct {
	c *mgo.Collection
}

type Database struct {
	db *mgo.Database
}

func ConnectMongoDB(url, name string) *Database {
	session, err := mgo.Dial(url)
	if err != nil {
		panic("error: Connecting, try restart mongod server")
	}

	return &Database{session.DB(name)}
}

func (d *Database) Map(name string) *Collection {
	return &Collection{d.db.C(name)}
}

// convert record from database mongodb to memory for executing
func (c *Collection) Memory(name string) []*Meta {

	// define some variable
	var result []*Meta

	query := c.c.Find(bson.M{"name": name})

	// fetch data from Meta Table
	err := query.All(&result)
	panicIfError(err)
	return result
}

func (m *Meta) ToJSON() []byte {
	metaByte, err := json.Marshal(*m)
	if err != nil {
		panic("error: encoding json " + err.Error())
	}

	return metaByte
}

// hash Meta ~~> MetaHash
func (m *Meta) Hash() *MetaHash {
	hashMeta := make(map[string]Attr)

	for i := 0; i < len(m.Fields); i++ {
		hashMeta[m.Fields[i].Name] = m.Fields[i]
	}
	return &MetaHash{hashMeta}
}

func (m *Meta) Find(opts map[string]interface{}) *MetaHash {

	newMeta := MetaHash{make(map[string]Attr)}

	for i := 0; i < len(m.Fields); i++ {

		// hanlde field by field
		field := m.Fields[i]
		ok := Compare(field, opts)

		if ok {
			newMeta.index[field.Name] = field
		}
	}

	return &newMeta
}

func (m *MetaHash) ToJSON() []byte {
	metaByte, err := json.Marshal(m.index)
	if err != nil {
		panic("error: encoding json " + err.Error())
	}

	return metaByte
}

func Compare(attr Attr, opts map[string]interface{}) bool {

	r := reflect.ValueOf(attr)

	for k, v := range opts {
		switch r.FieldByName(k).Kind() {

		case reflect.String:
			if r.FieldByName(k).String() != v.(string) {
				return false
			}
			break
		case reflect.Int, reflect.Int16, reflect.Int8, reflect.Int64, reflect.Int32:
			if r.FieldByName(k).Int() != int64(v.(int)) {
				return false
			}
			break
		case reflect.Bool:
			if r.FieldByName(k).Bool() != v.(bool) {
				return false
			}
		default:
			return false
		}

	}
	return true
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
