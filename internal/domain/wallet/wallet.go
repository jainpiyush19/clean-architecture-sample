package wallet

import "time"

//Info wallet info of user
type Info struct {
	UserID    int64
	Amount    float64
	UpdatedAt time.Time
}
