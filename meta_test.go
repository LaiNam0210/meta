package meta

import (
	_ "encoding/json"
	"io/ioutil"
	"libs/meta"
	"log"
	"testing"
)

func TestMap(t *testing.T) {
	db := meta.ConnectMongoDB("127.0.0.1", "financial")

	c := db.Map("meta")

	metaPackage := c.Memory("package")

	result := metaPackage[0].Find(map[string]interface{}{"Short": true})

	log.Println(result.ToJSON())
	ioutil.WriteFile("output/output.txt", result.ToJSON(), 0777)
	t.Error("checking")
}

func TestCompare(t *testing.T) {
	attr := meta.Attr{Short: true, Search: false, Name: "Period"}

	opts := map[string]interface{}{"Short": true, "Search": false, "Name": "Period"}

	ok := meta.Compare(attr, opts)
	if !ok {
		t.Error("Should be equal")
		return
	}
}
