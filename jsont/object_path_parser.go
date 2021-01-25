package jsont

import "strings"

type pathParseFn func(interface{}) (interface{}, error)

func parsePathValue(m map[string]interface{}, propertyPath string) (val interface{}, err error) {
	filter := parsePath(propertyPath)

	val = m
	for _, fn := range filter {
		val, err = fn(val)
		if err != nil {
			return nil, err
		}
	}

	return val, nil
}

// returns a filter chain that parses a property path string
func parsePath(propertyPath string) []pathParseFn {
	pathParts := strings.Split(propertyPath, ".")
	filter := make([]pathParseFn, len(pathParts))

	for i, part := range pathParts {
		filter[i] = get(part)
	}

	return filter
}

func get(key string) pathParseFn {
	return func(v interface{}) (interface{}, error) {
		m, ok := v.(map[string]interface{})
		if !ok {
			return nil, ErrPathIndexFailed
		}

		_, ok = m[key]
		if !ok {
			return nil, ErrPropertyDoesNotExist
		}

		return m[key], nil
	}
}
