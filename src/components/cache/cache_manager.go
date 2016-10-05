package cache

import (
	"reflect"
)

// cacheMap map to store cache interface
var cacheMap = make(map[string]CInterface)

// Set() stores the key with given type post init check
func Set(key string, conf *Config, myType reflect.Type) error {
	if val, ok := reflect.New(myType).Elem().Interface().(CInterface); ok {
		if _, ok = cacheMap[key]; ok {
			return getErrObj(ErrKeyPresent, "given key:"+key)
		}
		// check error for initialization
		if err := val.Init(conf); err != nil {
			return getErrObj(ErrInitialization, err.Error())
		}
		// store the new key
		cacheMap[key] = val
		return nil
	} else {
		return getErrObj(ErrWrongType, myType.String()+":does not implement CInterface")
	}

}

// Get() - returns the cache interface for given key
func Get(key string) (CInterface, error) {
	if val, ok := cacheMap[key]; !ok {
		return nil, getErrObj(ErrKeyNotPresent, "given key:"+key)
	} else {
		return val, nil
	}
}
