package main

import (
    "fmt" // пакет для форматированного ввода вывода
    "net/http" // пакет для поддержки HTTP протокола
    "strings" // пакет для работы с  UTF-8 строками
    "log" // пакет для логирования
)

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() //анализ аргументов,
    fmt.Println(r.Form)  // ввод информации о форме на стороне сервера
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    var name string
    var question string
    var answer string

    for k, v := range r.Form {
        switch k {
            case "name": name = strings.Join(v, ", ")
            case "question": question = strings.Join(v, ", ")
            case "answer": answer = strings.Join(v, ", ")
        }
    }

    fmt.Fprintf(w, "Привет, %s!\n", name) // отправляем данные на клиентскую сторону

        if question == "q1" {
            switch answer {
            case "a1":
                fmt.Fprintf(w, "Нет. Неверно.") 
            case "a2":
                fmt.Fprintf(w, "Правильный ответ.") 
            case "a3":
                fmt.Fprintf(w, "Ты идиот?!") 
            }
        }
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, r.URL.Path[1:])
}

func main() {
    //http.HandleFunc("/hello", HomeRouterHandler) // установим роутер
    http.HandleFunc("/test", HomeRouterHandler) // установим роутер
    http.HandleFunc("/", IndexHandler) // установим роутер
    err := http.ListenAndServe(":80", nil) // задаем слушать порт
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
