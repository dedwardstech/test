package count

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
