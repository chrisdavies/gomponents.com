package main

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type PageProps struct {
	Title string
	Body  g.Node
}

func Page(title, description, path string, body g.Node) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       title,
		Description: description,
		Language:    "en",
		Head: []g.Node{
			Link(Rel("stylesheet"), Href("/styles/app.css"), Type("text/css")),
			Link(Rel("stylesheet"), Href("/styles/highlightjs.min.css"), Type("text/css")),
			Script(Src("/scripts/highlight.min.js")),
			Script(g.Raw("hljs.highlightAll();")),
		},
		Body: []g.Node{
			Navbar(path),
			Container(true,
				Prose(
					body,
				),
				Footer(Class("mt-32 prose prose-sm prose-indigo"),
					P(
						g.Text("made in 🇩🇰 by "),
						A(Href("https://www.maragu.dk"), g.Text("maragu")),
					),
					P(
						g.Text("maker of "),
						A(Href("https://www.golang.dk"), g.Text("online Go courses")),
					),
				),
			),
		},
	})
}

// Container restricts the width of the children, centers, and adds some padding.
func Container(padY bool, children ...g.Node) g.Node {
	return Div(c.Classes{"max-w-7xl mx-auto px-4 sm:px-6 lg:px-8": true, "py-4 sm:py-6 lg:py-8": padY}, g.Group(children))
}

func Navbar(path string) g.Node {
	return Nav(Class("bg-gray-700 mb-6"),
		Container(false,
			Div(Class("flex items-center space-x-4 sm:space-x-6 lg:space-x-8 h-16"),
				NavbarLink("/", "Home", path),
				NavbarLink("/plus/", "gomponents +", path),
			),
		),
	)
}

func NavbarLink(path, text, currentPath string) g.Node {
	return A(Href(path), g.Text(text),
		c.Classes{
			"text-sm font-medium focus:outline-none focus:text-white hover:text-white": true,
			"text-white":    path == currentPath,
			"text-gray-300": path != currentPath,
		},
	)
}

func Prose(children ...g.Node) g.Node {
	return Div(Class("prose lg:prose-lg xl:prose-xl prose-indigo"), g.Group(children))
}

func CodeBlock(text string) g.Node {
	return Pre(Code(Class("language-go"), g.Text(text)))
}

func BashBlock(text string) g.Node {
	return Pre(Code(Class("language-bash"), g.Text(text)))
}
