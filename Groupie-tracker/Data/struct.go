package Data

type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Location     string   `json:"locations"`
	ConcertDates string   `json:"conertDates"`
	Relations    string   `json:"relations"`
}
type DatesInfo struct {
	Id   int      `json:"id"`
	Date []string `json:"dates"`
}
type Dates struct {
	Index []DatesInfo `json:"index"`
}
type Locations struct {
	Index []LocationsInfo `json:"index"`
}
type LocationsInfo struct {
	Id       int      `json:"id"`
	Location []string `json:"locations"`
	Date     string   `json:"dates"`
}
type Relation struct {
	Index []RelationInfo `json:"index"`
}
type RelationInfo struct {
	Id              int                 `json:"id"`
	Dates_Locations map[string][]string `json:"datesLocations"`
}
