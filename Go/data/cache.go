package data

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
}

func (c *CachedUsers) UserExists(uID string) bool {
	for _, v := range c.Users {
		if v == uID {
			return true
		}
	}
	return false
}
