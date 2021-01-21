package main
 
import (
    "encoding/json"
    "io/ioutil"
    "fmt"
    "log"
    "net/http"
    "kkn.fi/base62"
    "github.com/gorilla/mux"
)

type Site struct {
    Id    int64 `json:"Id"`
    URL   string `json:"Url"`
}

type Url struct {
    URL   string `json:"Url"`
}

var Sites []Site

func main() { 
    handleRequests()
	// Encode 
	var urlVal int64 = 3781504209452600
	encodedURL := base62.Encode(urlVal) 
	fmt.Println(encodedURL)
 
	// Decode 
	byteURL, _ := base62.Decode(encodedURL) 
    fmt.Println(byteURL)
}

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "/api/v1/new endpoint to shorten URL")
    fmt.Fprintf(w, "/api/v1/:url endpoint to retrive URL")
}

func createNewURL(w http.ResponseWriter, r *http.Request) {   
    reqBody, _ := ioutil.ReadAll(r.Body)
    var site Site
    var url Url
    json.Unmarshal(reqBody, &url)

    site.Id = 1
    site.URL = url.URL

    Sites = append(Sites, site)

    file, _ := json.MarshalIndent(Sites, "", " ")
    _ = ioutil.WriteFile("test.json", file, 0644)
    
    json.NewEncoder(w).Encode(Sites)
    fmt.Println(site)
}

func retriveURL(w http.ResponseWriter, r *http.Request) {
    /*vars := mux.Vars(r)
    id := vars["id"]

    for Site := range Sites {
        if Site.URL == url {
            byteURL, _ := base62.Decode(Site.ShortURL) 
            fmt.Fprintf(w, byteURL)
        }
    }*/
}

func handleRequests() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/api/v1/new", createNewURL).Methods("POST")
    router.HandleFunc("/api/v1/:url", retriveURL).Methods("GET")
    router.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":25565", router))
}