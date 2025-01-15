package models

type Phone struct {
	PhoneName        string  `db:"phone_name"`
	Brand            string  `db:"brand"`
	OS               string  `db:"os"`
	Inches           float64 `db:"inches"`
	Resolution       string  `db:"resolution"`
	Battery          int     `db:"battery"`
	BatteryType      string  `db:"battery_type"`
	RAM              int     `db:"ram"`
	AnnouncementDate string  `db:"announcement_date"`
	Weight           int     `db:"weight"`
	Storage          int     `db:"storage"`
	Video720p        bool    `db:"video_720p"`
	Video1080p       bool    `db:"video_1080p"`
	Video4K          bool    `db:"video_4k"`
	Video8K          bool    `db:"video_8k"`
	Video30fps       bool    `db:"video_30fps"`
	Video60fps       bool    `db:"video_60fps"`
	Video120fps      bool    `db:"video_120fps"`
	Video240fps      bool    `db:"video_240fps"`
	Video480fps      bool    `db:"video_480fps"`
	Video960fps      bool    `db:"video_960fps"`
	PriceUSD         float64 `db:"price_usd"`
}
