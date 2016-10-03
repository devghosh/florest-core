package utils

import (
	"github.com/jabong/florest-core/src/common/collections"
	"github.com/jabong/florest-core/src/common/collections/utils/comparators"
)

// Returns the string comparator
func GetStringComparator() collections.Comparator {
	return comparators.NewStringComparator()
}

// Returns the int comparator
func GetIntComparator() collections.Comparator {
	return comparators.NewIntComparator()
}

// Returns the int8 comparator
func GetInt8Comparator() collections.Comparator {
	return comparators.NewInt8Comparator()
}

// Returns the int16 comparator
func GetInt16Comparator() collections.Comparator {
	return comparators.NewInt16Comparator()
}

// Returns the int32 comparator
func GetInt32Comparator() collections.Comparator {
	return comparators.NewInt32Comparator()
}

// Returns the int64 comparator
func GetInt64Comparator() collections.Comparator {
	return comparators.NewInt64Comparator()
}

// Returns the uint comparator
func GetUIntComparator() collections.Comparator {
	return comparators.NewUIntComparator()
}

// Returns the uint8 comparator
func GetUInt8Comparator() collections.Comparator {
	return comparators.NewUInt8Comparator()
}

// Returns the uint16 comparator
func GetUInt16Comparator() collections.Comparator {
	return comparators.NewUInt16Comparator()
}

// Returns the uint32 comparator
func GetUInt32Comparator() collections.Comparator {
	return comparators.NewUInt32Comparator()
}

// Returns the uint64 comparator
func GetUInt64Comparator() collections.Comparator {
	return comparators.NewUInt64Comparator()
}

// Returns the byte comparator
func GetByteComparator() collections.Comparator {
	return comparators.NewByteComparator()
}

// Returns the rune comparator
func GetRuneComparator() collections.Comparator {
	return comparators.NewRuneComparator()
}
