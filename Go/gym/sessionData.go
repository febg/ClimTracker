package gym

type SessionData struct {
}

type ClimbingData struct {
	Data []DayData
}

type DayData struct {
	UId  string `json:"uID"`
	Date string `json:"Date"`
	V1   string `json:"V1"`
	V2   string `json:"V2"`
	V3   string `json:"V3"`
	V4   string `json:"V4"`
	V5   string `json:"V5"`
	V6   string `json:"V6"`
}

func (c *ClimbingData) Append(d DayData) error {
	c.Data = append(c.Data, d)
	return nil
}
