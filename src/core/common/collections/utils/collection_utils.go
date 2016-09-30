package utils

type emptyStruct struct {
}

func ConvertArrayToMap(inArray []interface{}) map[interface{}]int {
	arrMap := make(map[interface{}]int, inArray.Size())
	for index, element := range inArray {
		arrMap[element] = index
	}
	return arrMap
}
