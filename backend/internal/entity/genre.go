package entity

import "encoding/json"

type Genre struct {
	MediaID uint   `gorm:"primaryKey"`
	Genre   string `gorm:"primaryKey"`
}

func (g Genre) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.Genre)
}

func (g *Genre) UnmarshalJSON(data []byte) error {
	var genreString string
	if err := json.Unmarshal(data, &genreString); err != nil {
		return err
	}
	g.Genre = genreString
	return nil
}
