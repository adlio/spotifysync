package main

import (
	"fmt"

	"github.com/zmb3/spotify"
)

func main() {
	auth := spotify.NewAuthenticator("", spotify.ScopeUserReadPrivate)
	// if you didn't store your ID and secret key in SPOTIFY_ID, SPOTIFY_KEY
	// you can set them manually here
	// auth.SetAuthInfo(clientID, secretKey)
	fmt.Println("vim-go")
}
