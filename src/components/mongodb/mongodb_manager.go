package mongodb

import (
	"reflect"
)

// mongoMap map to store cache interface
var mongoMap = make(map[string]MDBInterface)

// Set() stores the key with given type post init check
func Set(key string, conf *MDBConfig, myType reflect.Type) error {
	if val, ok := reflect.New(myType).Elem().Interface().(MDBInterface); ok {
		if _, ok = mongoMap[key]; ok {
			return getErrObj(ErrKeyPresent, "given key:"+key)
		}
		// check error for initialization
		if err := val.Init(conf); err != nil {
			return getErrObj(ErrInitialization, err.Error())
		}
		// store the new key
		mongoMap[key] = val
		return nil
	} else {
		return getErrObj(ErrWrongType, myType.String()+":does not implement MDBInterface")
	}

}

// Get() - returns the mongodb interface for given key
func Get(key string) (MDBInterface, error) {
	if val, ok := mongoMap[key]; !ok {
		return nil, getErrObj(ErrKeyNotPresent, "given key:"+key)
	} else {
		return val, nil
	}
}
