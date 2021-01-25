package count

// SliceItems returns a map with a count of how many items are in
// a slice
func SliceItems(s []interface{}) map[interface{}]int {
    m := make(map[interface{}]int)

    for _, item := range s {
        if _, ok := m[item]; ok {
            m[item]++
        } else {
            m[item] = 1
        }
    }

    return m
}
