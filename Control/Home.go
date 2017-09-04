package main

import (

  "net/http"

  "log"
  "strings"

  "html/template"
  "strconv"
)

type PageVariables struct {

  PageTitle        string

  PagePersons      []Person

  Person           Person
}

func main() {
    http.HandleFunc("/", LoadPerson)

    http.HandleFunc("/add", AddPerson)

    http.HandleFunc("/delete/", DeletePerson)

    http.HandleFunc("/update/", UpdatePerson)

    http.HandleFunc("/confirmUpdate", ConfirmUpdate)
    log.Fatal(http.ListenAndServe(":8080", nil))

}

func LoadPerson(w http.ResponseWriter, r *http.Request){

    Title := "Person"
    var arrayPerson []Person;
    arrayPerson = load();

    MyPageVariables := PageVariables{

     PageTitle: Title,

     PagePersons : arrayPerson,

    }
    t, err := template.ParseFiles("View/select.html") //parse the html file homepage.html

    if err != nil { // if there is an error
     log.Print("template parsing error: ", err) // log it
    }

    err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps

    if err != nil { // if there is an error
     log.Print("template executing error: ", err) //log it
    }
}

func AddPerson(w http.ResponseWriter, r *http.Request){

    r.ParseForm()

    var person Person;
    person.Name  = r.Form.Get("pname")
    i32, err := strconv.ParseInt(r.Form.Get("page"), 10, 32)
    i := int(i32)
    person.Age   = i
    person.Phone = r.Form.Get("pphone")

    var Title string


    var check = insert(person.Name, person.Age, person.Phone)
    // generate page by passing page variables into template
    if(check){
    Title = "Add Success!"
    }else{
     Title = "Add Failed!"
    }
    MyPageVariables := PageVariables{

      PageTitle: Title,
      Person   : person,
    }
    t, err := template.ParseFiles("View/success.html") //parse the html file homepage.html

    if err != nil { // if there is an error
    log.Print("template parsing error: ", err) // log it
    }
    err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps

    if err != nil { // if there is an error
    log.Print("template executing error: ", err) //log it
    }

}

func DeletePerson(w http.ResponseWriter, r *http.Request){
    var Title string

    id := strings.TrimPrefix(r.URL.Path, "/delete/")
    i32, err := strconv.ParseInt(id, 10, 32)
    if err != nil { // if there is an error
       log.Print("casting error: ", err) // log it
    }
    i := int(i32)
    var check = delete(i)
    if(check){
        Title = "Delete Success!"
    }else{
        Title = "Delete Failed!"
    }
    MyPageVariables := PageVariables{
        PageTitle : Title,
    }
    t, err := template.ParseFiles("View/success.html") //parse the html file homepage.html

    if err != nil { // if there is an error
      log.Print("template parsing error: ", err) // log it
    }
    err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps

    if err != nil { // if there is an error
      log.Print("template executing error: ", err) //log it
    }
}

func UpdatePerson(w http.ResponseWriter, r *http.Request){
    var person Person;
    id := strings.TrimPrefix(r.URL.Path, "/update/")
    i32, err := strconv.ParseInt(id, 10, 32)
    if err != nil { // if there is an error
       log.Print("casting error: ", err) // log it
    }
    i := int(i32)
    person = getPerson(i)
    MyPageVariables := PageVariables{
        Person : person,
    }
    if(person.Name != ""){
        t, err := template.ParseFiles("View/update.html") //parse the html file homepage.html

            if err != nil { // if there is an error
              log.Print("template parsing error: ", err) // log it
            }
            err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps

            if err != nil { // if there is an error
              log.Print("template executing error: ", err) //log it
            }
    }else{
        Title := "Person not found!"
        MyPageVariables := PageVariables{
            PageTitle : Title,
        }
        t, err := template.ParseFiles("View/success.html") //parse the html file homepage.html

            if err != nil { // if there is an error
                log.Print("template parsing error: ", err) // log it
            }
            err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps

            if err != nil { // if there is an error
                log.Print("template executing error: ", err) //log it
            }
    }

}

func ConfirmUpdate(w http.ResponseWriter, r *http.Request){
    r.ParseForm()
    var person Person;
    person.Name  = r.Form.Get("pname")
    i32, err := strconv.ParseInt(r.Form.Get("page"), 10, 32)
    i := int(i32)
    person.Age   = i
    person.Phone = r.Form.Get("pphone")
    id32, err := strconv.ParseInt(r.Form.Get("pid"), 10, 32)
    id := int(id32)
    person.Id = id

    var Title string

    var check = updatePerson(person)
    if(check){
       Title = "Update Success!"
    }else{
        Title = "Update Failed!"
    }
    MyPageVariables := PageVariables{
       PageTitle: Title,
    }
    // generate page by passing page variables into template

    t, err := template.ParseFiles("View/success.html") //parse the html file homepage.html

    if err != nil { // if there is an error
        log.Print("template parsing error: ", err) // log it
    }
    err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps

    if err != nil { // if there is an error
        log.Print("template executing error: ", err) //log it
    }
}