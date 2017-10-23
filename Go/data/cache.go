package data

import "log"

type CachedUsers struct {
	Users []string
}

func InitializeCache() *CachedUsers {
	u := make([]string, 0)
	c := CachedUsers{
		Users: u,
	}
	return &c
}

func (c *CachedUsers) EmptyCache() {
	c.Users = c.Users[:0]
}

func (c *CachedUsers) AddUser(uID string) {

	c.Users = append(c.Users, uID)
	log.Printf("ADD: %v", c.Users)
}

func (c *CachedUsers) UserExists(uID string) bool {
	log.Printf("Cache: %v", c.Users)
	for _, v := range c.Users {
		log.Printf("test, %v", v)
		if v == uID {
			return true
		}
	}
	return false
}
