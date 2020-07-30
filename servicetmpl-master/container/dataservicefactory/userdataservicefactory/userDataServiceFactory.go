package userdataservicefactory

import (
	"github.com/jfeng45/servicetmpl/config"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/dataservice"
)

var udsFbMap = map[string]userDataServiceFbInterface{
	config.SQLDB:   &sqlUserDataServiceFactory{},
	config.COUCHDB: &couchdbUserDataServiceFactory{},
}

// The builder interface for factory method pattern
// Every factory needs to implement Build method
type userDataServiceFbInterface interface {
	Build(container.Container, *config.DataConfig) (dataservice.UserDataInterface, error)
}

// GetDataServiceFb is accessors for factoryBuilderMap
func GetUserDataServiceFb(key string) userDataServiceFbInterface {
	return udsFbMap[key]
}
