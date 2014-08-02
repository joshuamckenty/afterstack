package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()
  m.Use(render.Renderer())
  m.Get("/", func(r render.Render) {
        r.HTML(200, "index", nil)
  })

  m.Get("/hello", hello)
  m.Run()

}

func hello() string {
    return "hello, martini!"
}
