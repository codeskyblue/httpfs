package httpfs

import "testing"

func TestOpen(t *testing.T) {
	file, err := Open("https://raw.githubusercontent.com/codeskyblue/httpfs/master/LICENSE")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(file.ModTime())
	t.Log(file.Size())
	buf := make([]byte, 50)
	t.Log(file.ReadAt(buf, 1))
	t.Log(string(buf))
}
