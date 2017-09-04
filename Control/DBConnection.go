package main

import (
  "log"
  "fmt"
  "database/sql"
  	_ "github.com/go-sql-driver/mysql"


)

func connectDb() *sql.DB{

      //update connection string here
      db, err := sql.Open("mysql", "root:anhnh@tcp(localhost:3306)/demogo")
      if err != nil {
         fmt.Println("Some errors")

      }else{
          fmt.Println("Connect success")
      }
      err = db.Ping()
      if err != nil {
        fmt.Println("Some errors")

      }
      return db
}


func load() []Person {
      db := connectDb()
      var person Person;
      stmt, err := db.Prepare("select * from person")
      if err != nil {
        log.Fatal(err)
      }
      defer stmt.Close()
      rows, err := stmt.Query()
      var arrayPerson []Person

      for rows.Next() {
        err := rows.Scan(&person.Id, &person.Name ,&person.Age, &person.Phone)
        if err != nil {
            log.Fatal(err)
        }
        arrayPerson = append(arrayPerson, Person{Id:person.Id, Name:person.Name, Age:person.Age, Phone:person.Phone})
      }
      err = rows.Err()
      if err != nil {
        log.Fatal(err)
      }
      return arrayPerson;
}

func insert(name string, age int, phone string) bool{
     db := connectDb()
     stmt, err := db.Prepare("INSERT INTO person (name,age,phone) VALUES(?,?,?)")
     if err != nil {
         log.Fatal(err)
     }
     res, err := stmt.Exec(name,age,phone)
     defer stmt.Close()
     if err != nil {
         log.Fatal(err)
         return false;
     }
     lastId, err := res.LastInsertId()
     if err != nil {
         log.Fatal(err)
     }
     rowCnt, err := res.RowsAffected()
     if err != nil {
        log.Fatal(err)
     }

     log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
     return true;
}

func delete(id int) bool{
     db := connectDb()
     stmt, err := db.Prepare("Delete from person where id = ?")
       if err != nil {
         log.Fatal(err)
       }
       var person = getPerson(id);
       if(person.Name == ""){
            return false;
       }
       defer stmt.Close()
       stmt.Query(id)
     defer stmt.Close()
     if err != nil {
        log.Fatal(err)
        return false;
     }
     return true;
}

func getPerson(id int) Person{
       db := connectDb()
       var person Person;
       stmt, err := db.Prepare("select * from person where id = ?")
       if err != nil {
        log.Fatal(err)
       }
       defer stmt.Close()
       rows, err := stmt.Query(id)
       defer stmt.Close()
       for rows.Next() {
        err := rows.Scan(&person.Id, &person.Name ,&person.Age, &person.Phone)
        if err != nil {
            log.Fatal(err)
        }
       }
       err = rows.Err()
       if err != nil {
        log.Fatal(err)
       }

       return person;
}

func updatePerson(person Person) bool{
     db := connectDb()
    stmt, err := db.Prepare("Update person set name = ?, age = ?, phone = ? where id = ?")
    if err != nil {
        log.Fatal(err)
    }
    res, err := stmt.Exec(person.Name, person.Age, person.Phone, person.Id)
    defer stmt.Close()
    if err != nil {
        log.Fatal(err)
        return false;
    }

    lastId, err := res.LastInsertId()
    if err != nil {
        log.Fatal(err)
    }
    rowCnt, err := res.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(lastId, rowCnt)
    return true;
}