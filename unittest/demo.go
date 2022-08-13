package unittest

import (
	"encoding/json"
	"io/ioutil"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func (this *Monster) Store() bool {
	data, err := json.Marshal(this)
	if err != nil {
		return false
	}
	path := "/Users/ichpan/blogs/unittest/monster.json"
	err = ioutil.WriteFile(path, data, 0666)
	if err != nil {
		return false
	}
	return true
}

func (this *Monster) ReStore() bool {
	err := this.Store()
	if err != true {
		return false
	}
	path := "/Users/ichpan/blogs/unittest/monster.json"
	file, _ := ioutil.ReadFile(path)
	e := json.Unmarshal(file, this)
	if e != nil {
		return false
	}

	return true
}
