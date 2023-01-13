package db

import (
	"fmt"
	"github.com/mchirico/switch-context/fixtures"
	"testing"
)

func TestMDB(t *testing.T) {
	m := &MDB{}
	m.Add("test")
	m.Add("test2")

	m.Write(fixtures.Path("./testing/login.db"))

	m2 := MDB{}

	m2.Read(fixtures.Path("./testing/login.db"))

	if m2.Get("test") == 0 {
		t.Errorf("Expected value")
	}

	if m2.Get("junk") != 0 {
		t.Errorf("Expected 0")
	}
	fmt.Println(m2.Get("test").String())

}

func TestNewDBControler(t *testing.T) {
	c := NewDBController(fixtures.Path("./testing/login.db"))
	d := c.Get("test")
	fmt.Println(d.String())

	d = c.Get("junk")
	fmt.Println(d.String())

	c.Add("test33")
	sd := c.GetS("test3asd")
	fmt.Printf(">%s<", sd)

}

func TestNewDBController2(t *testing.T) {

	C.ChangeDir(fixtures.Path("./testing/login.db"))
	d := C.Get("test")
	fmt.Println(d.String())

	d = C.Get("junk")
	fmt.Println(d.String())

	C.Add("test33")
	sd := C.GetS("test3asd")
	fmt.Printf(">%s<", sd)

}
