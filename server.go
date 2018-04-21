package main

import (
        "fmt"
        "net/http"
        "html/template"
        "io/ioutil"

        "goji.io"
        "goji.io/pat"

        "github.com/Joker/jade"
)


type Person struct {
    Name   string
    Age    int
    Emails []string
    Jobs   []*Job
}

type Job struct {
    Employer string
    Role     string
}

func jobs(w http.ResponseWriter, r *http.Request) {

    buf, err := ioutil.ReadFile("templates/jobs.jade")
    if err != nil {
        fmt.Printf("\nReadFile error: %v", err)
        return
    }
    jadeTpl, err := jade.Parse("jade_tp", string(buf))
    if err != nil {
        fmt.Printf("\nParse error: %v", err)
        return
    }
    // fmt.Printf("%s", hpp.PrPrint(jadeTpl))

    //

    job1 := Job{Employer: "Monash B", Role: "Honorary"}
    job2 := Job{Employer: "Box Hill", Role: "Head of HE"}

    person := Person{
        Name:   "jan",
        Age:    50,
        Emails: []string{"jan@newmarch.name", "jan.newmarch@gmail.com"},
        Jobs:   []*Job{&job1, &job2},
    }

    //

    goTpl, err := template.New("html").Parse(jadeTpl)
    if err != nil {
        fmt.Printf("\nTemplate parse error: %v", err)
        return
    }
    err = goTpl.Execute(w, person)
    if err != nil {
        fmt.Printf("\nExecute error: %v", err)
        return
    }
}

// func jobs(w http.ResponseWriter, r *http.Request) {
//     dat, err := ioutil.ReadFile("templates/main.jade")
//     if err != nil {
//         fmt.Printf("ReadFile error: %v", err)
//         return
//     }

//     tmpl, err := jade.Parse("name_of_tpl", string(dat))
//     if err != nil {
//         fmt.Printf("Parse error: %v", err)
//         return
//     }
//     fmt.Fprintf(w, tmpl)
// }

func main() {
    mux := goji.NewMux()
    mux.HandleFunc(pat.Get("/jobs"), jobs)

    http.ListenAndServe("localhost:8000", mux)
}