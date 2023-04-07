package models

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"time"
)

type Template struct {
	*template.Template
}

type HtmlTemplate struct {
	Index      Template
	Category   Template
	Custom     Template
	Detail     Template
	Writing    Template
	Pigeonhole Template
	Login      Template
}

func (t *Template) ExecuteTemp(wr io.Writer, data interface{}) {
	err := t.Template.Execute(wr, data)
	t.ExecuteError(wr, err)
}

func (t *Template) ExecuteError(wr io.Writer, err error) {
	if err != nil {
		_, err := wr.Write([]byte(err.Error()))
		if err != nil {
			log.Println("ExecuteError:", err)
			return
		}
	}
}

func InitTemplate(dir string) (HtmlTemplate, error) {
	temp, err := readTemplate(
		[]string{"index", "category", "custom", "detail", "writing", "pigeonhole", "login"},
		dir)
	var htmlTemplate HtmlTemplate
	if err != nil {
		return htmlTemplate, err
	}
	htmlTemplate.Index = temp[0]
	htmlTemplate.Category = temp[1]
	htmlTemplate.Custom = temp[2]
	htmlTemplate.Detail = temp[3]
	htmlTemplate.Writing = temp[4]
	htmlTemplate.Pigeonhole = temp[5]
	htmlTemplate.Login = temp[6]
	return htmlTemplate, err
}

func IsODD(num int) bool {
	return num%2 == 0
}

func GetNextName(name []string, index int) string {
	return name[index+1]
}

func Date(layout string) string {
	return time.Now().Format(layout)
}

func DateDay(date time.Time) string {
	return date.Format("2006-01-02")
}
func readTemplate(temp []string, dir string) ([]Template, error) {
	var templateList []Template
	for _, v := range temp {
		viewName := v + ".html"
		t := template.New(viewName)
		home := dir + "home.html"
		header := dir + "layout/header.html"
		footer := dir + "layout/footer.html"
		personal := dir + "layout/personal.html"
		post := dir + "layout/post-list.html"
		pagination := dir + "layout/pagination.html"
		t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date, "dateDay": DateDay})
		t, err := t.ParseFiles(dir+viewName, home, header, footer, personal, post, pagination)
		if err != nil {
			log.Fatal(fmt.Sprintf("模板%s解析失败", viewName), err)
			return nil, err
		}

		var tb Template
		tb.Template = t
		templateList = append(templateList, tb)
		// templateList = append(templateList, Template{t})
	}
	return templateList, nil

}
