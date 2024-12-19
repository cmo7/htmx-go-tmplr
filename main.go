package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/cmo7/htmx-go-tmplr/templates"
)

func main() {
	component := templates.Layout(
		templates.Stack([]templ.Component{
			templates.Hello("Marce"),
			templates.Hello("Marce"),
			templates.Hello("Marce"),
			templates.Mouse(),
		}),
		templates.Head(
			"HEY HO",
			nil,
			[]string{
				"/static/css/main.css",
			},
			[]string{
				"/static/js/htmx.min.js",
			},
		),
	)

	http.Handle("/", templ.Handler(component))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("/mouse_entered", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		div := templates.Stack([]templ.Component{
			templates.Hello("Marce"),
			templates.Hello("Marce"),
			templates.Hello("Marce"),
			templates.Mouse(),
		})
		div.Render(context.Background(), w)
		fmt.Printf("Mouse entered!\n")
	}))

	http.ListenAndServe(":8080", nil)
}
