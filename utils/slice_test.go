package utils

import "testing"

func TestSlice(t *testing.T) {
	var data, err = ConvertToStringSlice([]float32{1.2, 2.3, 3.4, 4.5, 5, 6})
	t.Log(data)
	t.Log(err)
}
