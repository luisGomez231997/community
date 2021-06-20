package main

import {
    "log"
    "github.com/luisGomez231997/community/database"
    "github.com/luisGomez231997/community/handlers"
}

func main() {
    if database.conectarBD() == 0 {
        log.Fatal("Sin conexion a la base de datos")
        return
    }
    handlers.handlers()

}
