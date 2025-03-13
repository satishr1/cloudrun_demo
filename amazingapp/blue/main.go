// package main

// import (
//     "html/template"
//     "log"
//     "net/http"
// )

// func HomePage(w http.ResponseWriter, r *http.Request) {
//     t, _ := template.ParseFiles("homepage.html")
//     t.Execute(w, nil) 
// }

// func main() {
//     http.HandleFunc("/", HomePage)
//     log.Fatal(http.ListenAndServe(":8080", nil))
// }


package main

import (
        "html/template"
        "log"
        "net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
        tmpl, err := template.ParseFiles("homepage.html") // Replace "homepage.html" with your template file
        if err != nil {
                log.Printf("Error parsing template: %v", err)
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
                return
        }

        err = tmpl.Execute(w, nil) // Execute the loaded template
        if err != nil {
                log.Printf("Error executing template: %v", err)
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
                return
        }
}

func main() {
        http.HandleFunc("/", HomePage)
        log.Println("Server listening on port 8080")
        log.Fatal(http.ListenAndServe(":8080", nil))
}