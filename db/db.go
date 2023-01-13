package db

import (
	"encoding/json"
	"fmt"
	"github.com/mchirico/switch-context/constants"
	"github.com/mchirico/switch-context/fixtures"
	"os"
	"time"
)

var C *DBControler

func init() {
	C = NewDBController()
}

type DB struct {
	Login    time.Time     `json:"login"`
	Duration time.Duration `json:"duration"`
}

type MDB map[string]*DB

func (m *MDB) Add(key string) {
	(*m)[key] = &DB{
		Login:    time.Now(),
		Duration: 0,
	}
}

func (m *MDB) Get(key string) time.Duration {
	if _, ok := (*m)[key]; ok {
		return time.Since((*m)[key].Login)
	}
	return 0
}

func (m *MDB) Write(file string) error {
	jsonStr, err := json.Marshal(*m)
	if err != nil {
		return err
	}
	return os.WriteFile(file, []byte(jsonStr), 0644)
}

func (m *MDB) Read(file string) error {
	n, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(n, m)
	if err != nil {
		return err
	}
	return nil
}

type DBControler struct {
	m    *MDB
	File string
}

func NewDBController(file ...string) *DBControler {
	if len(file) == 0 {
		if dir, err := fixtures.HomeDirectory(); err != nil {
			file = append(file, "login.db")
		} else {
			file = append(file, dir+"/.switchcontext/login.db")
		}
	}

	return setup(file[0])
}

func (d *DBControler) ChangeDir(file string) {
	d = setup(file)
}

func setup(file string) *DBControler {
	m := &MDB{}
	err := m.Read(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}

	return &DBControler{
		m:    m,
		File: file,
	}
}

func (d *DBControler) Add(key string) {
	d.m.Add(key)
	d.m.Write(d.File)
}

func (d *DBControler) Read() {
	d.m.Read(d.File)
}

func (d *DBControler) Get(key string) time.Duration {
	d.m.Read(d.File)
	return d.m.Get(key)
}

func (d *DBControler) GetS(key string) string {
	err := d.m.Read(d.File)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return ""
	}
	result := d.m.Get(key)
	if result == 0 {
		return ""
	}
	if result.Hours() > constants.HOURS_BEFORE_EXPIRE {
		return ""
	}

	if result.Hours() > constants.HOURS_BEFORE_WARN {
		return fmt.Sprintf("Warning: %s about to expire", result.String())
	}

	return result.String()
}
