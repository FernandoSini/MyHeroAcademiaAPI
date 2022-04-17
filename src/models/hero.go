package models

type Hero struct {
	id        string      `json:"id,omitempty"`
	trueName  string      `json:"trueName"`
	lastName  string      `json:"lastName"`
	heroName  string      `json:"heroName"`
	heroRank  uint        `json:"heroRank,omitempty""`
	age       uint        `json:"age""`
	heroImage []HeroImage `json:"images,omitempty""`
}
