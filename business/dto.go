package business

type UserID uint
type ProfileID uint

type Profile struct {
	ID       ProfileID
	PhotoURL string
}

type User struct {
	ID       UserID
	Username string
	Email    string
	IsActive bool
}

type UserProfileAggregation struct {
	User    User
	Profile Profile
}

type PlainStruct struct {
	UserID    UserID
	ProfileID ProfileID
	Username  string
	PhotoURL  string
}

func (upa UserProfileAggregation) ToPlainStruct() PlainStruct {
	return PlainStruct{
		UserID:    upa.User.ID,
		ProfileID: upa.Profile.ID,
		Username:  upa.User.Username,
		PhotoURL:  upa.Profile.PhotoURL,
	}
}
