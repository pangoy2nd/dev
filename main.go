package main

import (
        "fmt"
        "net/http"

        "google.golang.org/appengine" // Required external App Engine library
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
        // if statement redirects all invalid URLs to the root homepage.
        // Ex: if URL is http://[YOUR_PROJECT_ID].appspot.com/FOO, it will be
        // redirected to http://[YOUR_PROJECT_ID].appspot.com.
        if r.URL.Path != "/" {
                http.Redirect(w, r, "/", http.StatusFound)
                return
        }

        fmt.Fprintln(w, "1 John 3:23-24 - And this is his commandment, That we should believe on the name of his Son Jesus Christ, and love one another, as he gave us commandment. And he that keepeth his commandments dwelleth in him, and he in him. And hereby we know that he abideth in us, by the Spirit which he hath given us. - www.BibleStory.info")
}

func main() {
        http.HandleFunc("/", indexHandler)
        appengine.Main() // Starts the server to receive requests
}
