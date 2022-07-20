package models

import "gopkg.in/validator.v2"

//Fair a struct used to define "feira livre"
type Fair struct {
	Id            int    `json:"id,omitempty"`
	Longitude     string `json:"longitude,omitempty" validate:"nonzero"`
	Latitude      string `json:"latitude,omitempty" validate:"nonzero"`
	SetCen        string `json:"set_cen,omitempty" validate:"nonzero"`
	AreaP         string `json:"area_p,omitempty"  validate:"nonzero"`
	CodDist       string `json:"cod_dist,omitempty"  validate:"nonzero"`
	District      string `json:"district,omitempty"  validate:"nonzero"`
	CodSubPref    string `json:"cod_sub_pref,omitempty"  validate:"nonzero"`
	SubPref       string `json:"sub_pref,omitempty" validate:"nonzero"`
	RegionFive    string `json:"region_Five,omitempty" validate:"nonzero"`
	RegionEight   string `json:"region_Eight,omitempty" validate:"nonzero"`
	NameFair      string `json:"name_fair,omitempty" validate:"nonzero"`
	Record        string `json:"record,omitempty"  validate:"nonzero"`
	Street        string `json:"street,omitempty"  validate:"nonzero"`
	Number        string `json:"number,omitempty"  validate:"nonzero"`
	Neighbourhood string `json:"neighbourhood,omitempty" validate:"nonzero"`
	Reference     string `json:"reference,omitempty" validate:"nonzero"`
}

func ValidateFair(fair *Fair) error {
	err := validator.Validate(fair)
	if err != nil {
		return err
	}
	return nil
}
