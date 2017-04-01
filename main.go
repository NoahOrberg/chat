package main

import (
	"log"
	"net/http"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler)ServeHTTP(w http.ResponseWriter, r http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func title(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		<html>
			<head>
				<title>CHAT</title>
			</head>
			<body>
				ちゃっと
			</body>
		</html>
	`))
}

func main() {
	http.HandleFunc("/", title)
	if err := http.ListenAndServe(":8181", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
