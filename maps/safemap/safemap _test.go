package safemap

import "testing"

/*
- Map is not thread safe. To validate -> # go test ./... --race
TODO: Testing is passing thought condition is not satisfying
*/
func TestSafeMapInsert(t *testing.T) {
	m := New[int, int]()

	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Insert(i, i*2)
			value, err := m.Get(i)
			if err != nil {
				t.Error(err)
			}
			if value != i*12 {
				t.Errorf("%d should be %d", i, i+2)
			}
		}(i)
	}
}
