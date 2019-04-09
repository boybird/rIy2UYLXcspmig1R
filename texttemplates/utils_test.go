package texttemplates

import (
	"bytes"
	"fmt"
	"testing"
)

func TestLoadPage(t *testing.T) {
	p, err := LoadPage("utils_test")
	if err != nil {
		p = &Paginator{0, 0, 0, 10, nil, []byte{}}
	}
	buf := bytes.NewBufferString("")
	RenderTemplate(buf, "utils_test", p)
	fmt.Printf("%s", buf.String())
}
