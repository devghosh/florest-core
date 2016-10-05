package sqldb

import (
	"reflect"
)

// sdbMap map to store cache interface
var sdbMap = make(map[string]SDBInterface)

// Set() stores the key with given type post init check
func Set(key string, conf *SDBConfig, myType reflect.Type) error {
	if val, ok := reflect.New(myType).Elem().Interface().(SDBInterface); ok {
		if _, ok = sdbMap[key]; ok {
			return getErrObj(ErrKeyPresent, "given key:"+key)
		}
		// check error for initialization
		if err := val.Init(conf); err != nil {
			return getErrObj(ErrInitialization, err.Error())
		}
		// store the new key
		sdbMap[key] = val
		return nil
	} else {
		return getErrObj(ErrWrongType, myType.String()+":does not implement SDBInterface")
	}

}

// Get() - returns the sql db interface for given key
func Get(key string) (SDBInterface, error) {
	if val, ok := sdbMap[key]; !ok {
		return nil, getErrObj(ErrKeyNotPresent, "given key:"+key)
	} else {
		return val, nil
	}
}
