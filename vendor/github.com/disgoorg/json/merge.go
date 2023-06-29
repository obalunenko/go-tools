package json

// SimpleMerge merges two JSON objects into one. If a key exists in both objects, the value from the second object is used.
func SimpleMerge(data1 []byte, data2 []byte) ([]byte, error) {
	m1 := map[string]any{}
	if err := Unmarshal(data1, &m1); err != nil {
		return nil, err
	}
	m2 := map[string]any{}
	if err := Unmarshal(data2, &m2); err != nil {
		return nil, err
	}
	for k, v := range m2 {
		m1[k] = v
	}
	return Marshal(m1)
}
