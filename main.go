package main

import "fmt"
import "github.com/cyruzin/golang-tmdb"

tmdbClient, err := tmdb.Init(os.GetEnv("YOUR_APIKEY"))
if err != nil {
    fmt.Println(err)
}

// OPTIONAL: Setting a custom config for the http.Client.
// The default timeout is 10 seconds. Here you can set other
// options like Timeout and Transport.
customClient := http.Client{
    Timeout: time.Second * 5,
    Transport: &http.Transport{
        MaxIdleConns: 10,
        IdleConnTimeout: 15 * time.Second,
    }
}

tmdbClient.SetClientConfig(customClient)

// OPTIONAL: Enable this option if you're going to use endpoints
// that needs session id.
//
// You can read more about how this works:
// https://developers.themoviedb.org/3/authentication/how-do-i-generate-a-session-id
tmdbClient.SetSessionID(os.GetEnv("YOUR_SESSION_ID"))

// OPTIONAL (Recommended): Enabling auto retry functionality.
// This option will retry if the previous request fail (429 TOO MANY REQUESTS).
tmdbClient.SetClientAutoRetry()

movie, err := tmdbClient.GetMovieDetails(297802, nil)
if err != nil {
 fmt.Println(err)
}

fmt.Println(movie.Title)

func main() {
	fmt.Print("uhhh.. here?\n")
}
