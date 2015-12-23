  package main
          import (
                  “fmt” 
                  “log” 
                  “net/http” 
                  “strconv” 
                 )
                  
func main() { 

  http.HandleFunc(“/”, func(w http.ResponseWriter, r *http.Request) { 

    s := r.URL.Query().Get(“s”)
    i , _:= strconv.Atoi(s) 

    if(i != 0) {

      if i%400 == 0 {
          fmt.Fprintln(w, s + “ artık yıl”) 
        } else if i%100 == 0{
          fmt.Fprintln(w, s +” artık yıl”)
        } else if i%4==0 {
          fmt.Fprintln(w, s +” artık yıl”)
        } else {
          fmt.Fprintln(w, s +” artık yıl değil”)
        }
              }
      })

        log.Fatal(http.ListenAndServe(“:8080”, nil)) 
  
}
