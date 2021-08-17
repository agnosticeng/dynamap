package dynamap

import (
	"errors"
	"strconv"
)

var (
	ErrWrongNodeType        = errors.New("WRONG_NODE_TYPE")
	ErrWrongPathSegmentType = errors.New("WRONG_PATH_SEGMENT_TYPE")
	ErrImpossibleState      = errors.New("IMPOSSIBLE_STATE")
)

func SPathToPath(path ...string) []interface{} {
	var iPath = make([]interface{}, len(path))

	for i, pathSegment := range path {
		numIdx, err := strconv.ParseInt(pathSegment, 10, 64)

		if err != nil {
			iPath[i] = pathSegment
		} else {
			iPath[i] = numIdx
		}
	}

	return iPath
}

func Set(node interface{}, v interface{}, path ...interface{}) (interface{}, error) {
	if len(path) == 0 {
		return v, nil
	}

	var (
		pathSegment = path[0]
		newPath     = path[1:]
	)

	switch pathSegment := pathSegment.(type) {
	case int:
		var (
			slice []interface{}
			ok    bool
			err   error
		)

		if node == nil {
			slice = make([]interface{}, pathSegment+1)
		} else {
			slice, ok = node.([]interface{})

			if !ok {
				return nil, ErrWrongNodeType
			}

			if len(slice) < (pathSegment + 1) {
				newSlice := make([]interface{}, pathSegment+1)
				copy(newSlice, slice)
				slice = newSlice
			}
		}

		slice[pathSegment], err = Set(slice[pathSegment], v, newPath...)

		if err != nil {
			return nil, err
		}

		return slice, nil

	case string:
		var (
			m   map[string]interface{}
			ok  bool
			err error
		)

		if node == nil {
			m = make(map[string]interface{})
		} else {
			m, ok = node.(map[string]interface{})

			if !ok {
				return nil, ErrWrongNodeType
			}
		}

		m[pathSegment], err = Set(m[pathSegment], v, newPath...)

		if err != nil {
			return nil, err
		}

		return m, nil

	default:
		return nil, ErrWrongPathSegmentType
	}

	return nil, ErrImpossibleState
}

func SSet(node interface{}, v interface{}, path ...string) (interface{}, error) {
	return Set(node, v, SPathToPath(path...))
}

func Get(node interface{}, path ...interface{}) (interface{}, error) {
	if len(path) == 0 {
		return node, nil
	}

	var (
		pathSegment = path[0]
		newPath     = path[1:]
	)

	switch pathSegment := pathSegment.(type) {
	case int:
		var (
			slice []interface{}
			ok    bool
		)

		if node == nil {
			return nil, nil
		}

		slice, ok = node.([]interface{})

		if !ok {
			return nil, ErrWrongNodeType
		}

		if len(slice) < (pathSegment + 1) {
			return nil, nil
		}

		return Get(slice[pathSegment], newPath...)

	case string:
		var (
			m  map[string]interface{}
			ok bool
		)

		if node == nil {
			return nil, nil
		}

		m, ok = node.(map[string]interface{})

		if !ok {
			return nil, ErrWrongNodeType
		}

		return Get(m[pathSegment], newPath...)

	default:
		return nil, ErrWrongPathSegmentType
	}

	return nil, ErrImpossibleState
}

func SGet(node interface{}, path ...string) (interface{}, error) {
	return Get(node, SPathToPath(path...))
}
