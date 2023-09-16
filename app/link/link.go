package link

import "usergraphql/app/user"

type Link struct {
	Id      string     `json:"id,omitempty"`
	Title   string     `json:"title,omitempty"`
	Address string     `json:"address,omitempty"`
	User    *user.User `json:"user,omitempty"`
}
