package models

import (
	"encoding/json"
	"log"
)

type User struct {
	id           string
	username     string
	access_level int
}

// NewUser creates a new user
func NewUser(id string, username string, access_level int) *User {
	return &User{
		id:           id,
		username:     username,
		access_level: access_level,
	}
}

// MarshalJSON marshalls the user into a json object as a []byte
func (u *User) MarshalJSON() ([]byte, error) {
	log.Println("Marshalling user")
	return json.Marshal(struct {
		Id           string `json:"id"`
		Username     string `json:"username"`
		Access_level int    `json:"access_level"`
	}{
		Id:           u.id,
		Username:     u.username,
		Access_level: u.access_level,
	})
}

// UnmarshalJSON unmarshalls the user from a json object
func (u *User) UnmarshalJSON(data []byte) error {
	tmp := struct {
		Id           string `json:"id"`
		Username     string `json:"username"`
		Access_level int    `json:"access_level"`
	}{
		Id:           u.id,
		Username:     u.username,
		Access_level: u.access_level,
	}

	err := json.Unmarshal(data, &tmp)

	u.id = tmp.Id
	u.username = tmp.Username
	u.access_level = tmp.Access_level

	log.Println("Unmarshalling user")
	log.Println(u.id, u.username, u.access_level)

	return err
}

func (u *User) GetId() string {
	return u.id
}

func (u *User) GetUsername() string {
	return u.username
}

func (u *User) SetUsername(username string) {
	u.username = username
}

func (u *User) SetId(id string) {
	u.id = id
}

func (u *User) GetAccessLevel() int {
	return u.access_level
}

func (u *User) SetAccessLevel(access_level int) {
	u.access_level = access_level
}
