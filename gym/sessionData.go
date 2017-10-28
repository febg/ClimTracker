package gym

type SessionData struct {
}

type ClimbingData struct {
	Data  []DayData
	OData []OverallData
}

type PublicClimbingData struct {
	OData []OverallData
}

type DayData struct {
	Index string `json:"index"`
	Date  string `json:"Date"`
	UId   string `json:"uID"`
	V1    string `json:"V1"`
	V2    string `json:"V2"`
	V3    string `json:"V3"`
	V4    string `json:"V4"`
	V5    string `json:"V5"`
	V6    string `json:"V6"`
}
type OverallData struct {
	SDate  string `json:"StatsDate"`
	MV1    string `json:"MV1"`
	MV2    string `json:"MV2"`
	MV3    string `json:"MV3"`
	MV4    string `json:"MV4"`
	MV5    string `json:"MV5"`
	MV6    string `json:"MV6"`
	Total  string `json:"Total"`
	PDate  string `json:"Pullup Date"`
	PCount string `json:"PullCount"`
	PMax   string `jason:"PullMax"`
}

type PullUpData struct {
	Index   string `json:"Index"`
	Date    string `json:"Date"`
	Count   string `json:"Count"`
	MaxPull string `jason:"MaxPull"`
}

type Date struct {
	Day  string `json:"Date"`
	Data []Data `json:"Day Data"`
}

type Data struct {
	Index string `json:"index"`
	UId   string `json:"uID"`
	V1    string `json:"V1"`
	V2    string `json:"V2"`
	V3    string `json:"V3"`
	V4    string `json:"V4"`
	V5    string `json:"V5"`
	V6    string `json:"V6"`
}

func (c *Date) Append(d Data) error {
	c.Data = append(c.Data, d)

	return nil
}

func (c *Date) AppendExistingDate(d Data) {

}

func (c *ClimbingData) Append(d DayData) error {
	c.Data = append(c.Data, d)
	return nil
}
