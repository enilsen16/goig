package main

import (
	"github.com/ahmdrz/goinsta/v2"
	"log"
)

// TODO: Login
// Once logged in, iterate through user urls or usernames
// Get there last updated photo
// Like the photo

func main() {
	insta := goinsta.New("", "")
	if err := insta.Login(); err != nil {
		log.Println(err)
		return
	}
	defer insta.Logout()

	profile, err := insta.Profiles.ByName("")
	check(err)
	latest := profile.Feed()
	latest.Next(false)
	item := latest.Items[0]
	item.Like()
}
