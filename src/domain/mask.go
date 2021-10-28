package domain

func (u *User) Masked() MaskedUser {
	return MaskedUser{
		ID:          u.ID,
		UserName:    u.UserName,
		Image:       u.Image,
		DisplayName: u.DisplayName,
		Contacts:    u.Contacts,
		Creator:     u.Creator,
	}
}
