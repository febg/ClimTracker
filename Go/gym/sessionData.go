package gym

type SessionData struct {
}

type ClimbingData struct {
	Data []DayData
}

type DayData struct {
	Date string `json:"Date"`
	V1   string `json:"V1"`
	V2   string `json:"V2"`
	V3   string `json:"V3"`
	V4   string `json:"V4"`
	V5   string `json:"V5"`
	V6   string `json:"V6"`
}

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

func (c *ClimbingData) Append(d DayData) error {
	c.Data = append(c.Data, d)
	return nil
}
