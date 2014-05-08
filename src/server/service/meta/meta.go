package meta

import (
	"labix.org/v2/mgo"
)

var (
	DB *mgo.Database
)

type MetaObject struct {
	DBName string
}

func (m *MetaObject) GetMetaObject() {

}
