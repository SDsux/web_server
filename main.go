package main

/*
(1) render array
$ curl http://localhost:9999/date
<html>
<body>
    <p>hello, gee</p>
    <p>Date: 2019-08-17</p>
</body>
</html>
*/

/*
(2) custom render function
$ curl http://localhost:9999/students
<html>
<body>
    <p>hello, gee</p>
    <p>0: Geektutu is 20 years old</p>
    <p>1: Jack is 22 years old</p>
</body>
</html>
*/

/*
(3) serve static files
$ curl http://localhost:9999/assets/css/geektutu.css
p {
    color: orange;
    font-weight: 700;
    font-size: 20px;
}
*/

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"gee"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gee.Default()
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")
	//CSS文件用于描述网页或文档的样式和布局。
	//它定义了HTML元素的外观、颜色、字体、边框、背景等样式属性，用于控制网页的外观效果。

	//.tmpl文件是一种模板文件，用于生成其他文件或动态内容。
	//它包含特定格式和占位符，用于生成最终的文件或内容。通常用于生成HTML页面、邮件、配置文件等。
	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})
	/*time.Date函数用于创建一个指定日期和时间的time.Time对象
		参数说明：
	year：年份，如2022。
	month：月份，是time.Month类型的值，表示1月到12月。
	day：日期，表示一个月中的某一天。
	hour：小时，表示一天中的某个小时。
	min：分钟，表示小时中的某个分钟。
	sec：秒，表示分钟中的某个秒。
	nsec：纳秒，表示秒中的某个纳秒。
	loc：时区，表示日期和时间的时区信息。如果为nil，表示使用本地时区。
	time.Date函数返回一个time.Time类型的对象，表示指定的日期和时间。这里时区参数使用了time.UTC，表示使用UTC时区

	*/
	r.GET("/date", func(c *gee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title": "gee",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")
}
