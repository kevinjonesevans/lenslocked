package main

import (
    "html/template"
    "os"
)

func main() {
    t, err := template.ParseFiles("hello.gohtml")
    if err != nil {
        panic(err)
    }

    data := struct {
        Name string
        Age int
    }{"John Smith", 30}

    err = t.Execute(os.Stdout, data)
    if err != nil {
        panic(err)
    }
}