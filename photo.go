package photoprism

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Photo struct {
	gorm.Model
	Title          string
	Description    string `gorm:"type:text;"`
	Artist         string
	Keywords       string
	Colors         string
	VibrantColor   string
	MutedColor     string
	TakenAt        time.Time
	CanonicalName  string
	PerceptualHash string
	Tags           []Tag `gorm:"many2many:photo_tags;"`
	Files          []File
	Albums         []Album `gorm:"many2many:album_photos;"`
	Camera         *Camera
	CameraID       uint
	Lat            float64
	Long           float64
	Location       *Location
	LocationID     uint
	Favorite       bool
	Private        bool
	Deleted        bool
}