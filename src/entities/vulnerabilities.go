package entities

type Vulnerability struct {
	Oid          string  `json:"oid"`
	Name         string  `json:"name"`
	Severity     float64 `json:"severity"`
	Qod          int     `json:"qod"`
	SolutionType string  `json:"solutionType"`
	AssetId      string  `json:"assetId"`
}
