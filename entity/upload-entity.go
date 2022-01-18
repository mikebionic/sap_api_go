package entity

type Upload struct {
	Name       string `json:"Name"`
	Path       string `json:"Path"`
	ImageGuid  string `json:"Guid"`
	Error      string `json:"Error"`
	TargetGuid string `json:"TargetGuid"`
}
