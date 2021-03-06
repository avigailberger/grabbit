package saga

import (
	"database/sql"
	"reflect"

	"github.com/rhinof/grabbit/gbus"
)

//Store abtracts the way sagas get persisted
type Store interface {
	RegisterSagaType(saga gbus.Saga)
	GetSagaByID(tx *sql.Tx, sagaID string) (*Instance, error)
	GetSagasByType(tx *sql.Tx, sagaType reflect.Type) ([]*Instance, error)
	SaveNewSaga(tx *sql.Tx, sagaType reflect.Type, newInstance *Instance) error
	UpdateSaga(tx *sql.Tx, instance *Instance) error
	DeleteSaga(tx *sql.Tx, instance *Instance) error
}
