package structs

type VData struct {
	VSType      string `json:"type"`
	Value       uint64 `json:"value"`
	Flag        string `json:"flag"`
	Description string `json:"description"`
}
