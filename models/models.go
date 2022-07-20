package models

//Fair a struct used to define "feira livre"
type Fair struct {
	Id            int    `json:"id,omitempty"`
	Longitude     string `json:"longitude,omitempty"`
	Latitude      string `json:"latitude,omitempty"`
	SetCen        string `json:"set_cen,omitempty"`
	AreaP         string `json:"area_p,omitempty"`
	CodDist       string `json:"cod_dist,omitempty"`
	District      string `json:"district,omitempty"`
	CodSubPref    string `json:"cod_sub_pref,omitempty"`
	SubPref       string `json:"sub_pref,omitempty"`
	RegionFive    string `json:"region_Five,omitempty"`
	RegionEight   string `json:"region_Eight,omitempty"`
	NameFair      string `json:"name_fair,omitempty"`
	Record        string `json:"record,omitempty"`
	Street        string `json:"street,omitempty"`
	Number        string `json:"number,omitempty"`
	Neighbourhood string `json:"neighbourhood,omitempty"`
	Reference     string `json:"reference,omitempty"`
}
