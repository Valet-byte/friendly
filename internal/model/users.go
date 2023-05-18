package model

type User struct {
	Uid       string         `json:"uid"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	IsPremium bool           `json:"isBro"`
	IsEnable  bool           `json:"is_enable"`
	City      string         `json:"city"`
	Networks  SocialNetworks `json:"networks"`
}

func (u *User) Combine(userData *User) {
	if userData.Uid != "" {
		u.Uid = userData.Uid
	}
	if userData.Name != "" {
		u.Name = userData.Name
	}
	if userData.Email != "" {
		u.Email = userData.Email
	}
	if userData.City != "" {
		u.City = userData.City
	}
	if userData.IsPremium != u.IsPremium {
		u.IsPremium = userData.IsPremium
	}
	if userData.IsEnable != u.IsEnable {
		u.IsEnable = userData.IsEnable
	}
	if !userData.Networks.IsEmpty() {
		u.Networks = userData.Networks
	}
}

type RegistrationUserData struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserI interface {
	GetType() string
	GetId() int64
	GetName() string
}
