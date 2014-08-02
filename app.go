package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()

  m.Get("/", hello)

  m.Run()

}

func hello() string {
    return "hello, martini!"
}
