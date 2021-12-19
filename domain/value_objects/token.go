package value_objects

type Token struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	ATExpires    int64
	ATDuration   int
	RTExpires    int64
	RTDuration   int
	Username     string
}
