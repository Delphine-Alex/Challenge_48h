package models

type Jointure struct {
	Uuid             string `json:"uuid"`
	Title            string `json:"title"`
	TypeInfo         string `json:"type_info"`
	Location         string `json:"location"`
	Start            string `json:"start"`
	End              string `json:"end"`
	Uuid_organizator string `json:"uuid_orga"`
	Uuid_orga        string `json:"uuid_orga_bis"`
	Name             string `json:"name"`
	Referent         string `json:"referent"`
	Email            string `json:"email"`
	Username         string `json:"username"`
	Password         string `json:"password"`
}
