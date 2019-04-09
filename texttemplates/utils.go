package texttemplates

import (
	"html/template"
	"io"
	"io/ioutil"
	"os"
)

type Paginator struct {
	Page     int
	Total    int
	Page_num int
	Pagesize int
	Data     []interface{}
	Body     []byte
}

type Person struct {
	Name string
	Age  int
}

func LoadPage(title string) (*Paginator, error) {
	filename := title + ".html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	ps := make([]interface{}, 2)

	ps[0] = Person{"zxk", 1}
	ps[1] = Person{"zzy", 2}
	return &Paginator{
		1,
		1,
		1,
		10,
		ps,
		body,
	}, nil
}

func RenderTemplate(w io.Writer, tmpl string, p *Paginator) {
	t, _ := template.ParseFiles(tmpl + ".html")
	_ = t.Execute(os.Stdout, p)
}
