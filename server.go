//this is my first commit

package main

import (
    "encoding/json"
	"fmt"
	"log"
	"net/http"
    "os"
	"time"
)


func main() {
	http.HandleFunc("/time", main1)
    log.Fatal(http.ListenAndServe(":8201", nil))

}
type User struct {
   Time     time.Time
}

func main1(w http.ResponseWriter, r *http.Request) {

    u := &User{
        Time: time.Now(),
    }

    out, err := json.MarshalIndent(u, "", "  ")
    if err != nil {
        log.Println(err)
        os.Exit(1)
    }

    fmt.Println(string(out))
}
