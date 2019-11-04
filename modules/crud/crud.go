package crud

import (
	"log"
	"sync"

	"github.com/spaceuptech/space-cloud/config"
	"github.com/spaceuptech/space-cloud/model"
	"github.com/spaceuptech/space-cloud/utils"

	"github.com/spaceuptech/space-cloud/modules/crud/driver"
)

// Module is the root block providing convenient wrappers
type Module struct {
	sync.RWMutex
	blocks    map[string]*stub
	primaryDB string
	project   string

	// Variables to store the hooks
	hooks      *model.CrudHooks
	metricHook model.MetricCrudHook

	// Drivers handler
	h *driver.Handler
}

// Init create a new instance of the Module object
func Init(h *driver.Handler) *Module {
	return &Module{blocks: make(map[string]*stub), h: h}
}

// SetConfig set the rules adn secret key required by the crud block
func (m *Module) SetConfig(project string, crud config.Crud) error {
	m.Lock()
	defer m.Unlock()

	m.project = project

	// Close the previous database connections
	for _, v := range m.blocks {
		m.h.RemoveBlock(v.dbType, v.conn)
	}

	m.blocks = make(map[string]*stub, len(crud))

	// Create a new crud blocks
	for dbType, v := range crud {
		// Initialise a new block
		c, err := m.h.InitBlock(utils.DBType(dbType), v.Enabled, v.Conn)
		m.blocks[dbType] = &stub{c: c, conn: v.Conn, dbType: utils.DBType(dbType)}

		if err != nil {
			log.Println("Error connecting to " + dbType + " : " + err.Error())
			return err
		} else {
			log.Println("Successfully connected to " + dbType)
		}
	}
	return nil
}

// SetHooks sets the internal hooks
func (m *Module) SetHooks(hooks *model.CrudHooks, metricHook model.MetricCrudHook) {
	m.hooks = hooks
	m.metricHook = metricHook
}

type stub struct {
	conn   string
	c      driver.Crud
	dbType utils.DBType
}

func (m *Module) getCrudBlock(dbType string) (driver.Crud, error) {
	if crud, p := m.blocks[dbType]; p {
		return crud.c, nil
	}

	return nil, utils.ErrDatabaseConfigAbsent
}
