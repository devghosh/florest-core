package utils

import (
	"github.com/jabong/florest-core/src/common/collections"
	"github.com/jabong/florest-core/src/common/collections/utils/comparators"
)

func GetStringComparator() collections.Comparator {
	return comparators.NewStringComparator()
}

func GetIntComparator() collections.Comparator {
	return comparators.NewIntComparator()
}

func GetInt8Comparator() collections.Comparator {
	return comparators.NewInt8Comparator()
}

func GetInt16Comparator() collections.Comparator {
	return comparators.NewInt16Comparator()
}

func GetInt32Comparator() collections.Comparator {
	return comparators.NewInt32Comparator()
}

func GetInt64Comparator() collections.Comparator {
	return comparators.NewInt64Comparator()
}

func GetUIntComparator() collections.Comparator {
	return comparators.NewUIntComparator()
}

func GetUInt8Comparator() collections.Comparator {
	return comparators.NewUInt8Comparator()
}

func GetUInt16Comparator() collections.Comparator {
	return comparators.NewUInt16Comparator()
}

func GetUInt32Comparator() collections.Comparator {
	return comparators.NewUInt32Comparator()
}

func GetUInt64Comparator() collections.Comparator {
	return comparators.NewUInt64Comparator()
}

func GetByteComparator() collections.Comparator {
	return comparators.NewByteComparator()
}

func GetRuneComparator() collections.Comparator {
	return comparators.NewRuneComparator()
}
