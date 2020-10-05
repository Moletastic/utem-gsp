package models

type IProfile interface {
	GetID() uint
}

type TeacherProfile struct {
	ID        uint `json:"id"`
	EntryYear int  `json:"entry_year"`
}

func NewTeacherProfile(t *Teacher) TeacherProfile {
	p := TeacherProfile{
		ID:        t.ID,
		EntryYear: t.EntryYear,
	}
	return p
}

func (p *TeacherProfile) GetID() uint {
	return p.ID
}

type AdminProfile struct {
	ID        uint `json:"id"`
	EntryYear int  `json:"entry_year"`
}

func NewAdminProfile(a *Admin) AdminProfile {
	p := AdminProfile{
		ID:        a.ID,
		EntryYear: a.EntryYear,
	}
	return p
}

func (p *AdminProfile) GetID() uint {
	return p.ID
}

type ProfiledUser struct {
	User
	Profile IProfile `json:"profile"`
}

func NewProfiledUser(u *User, profile IProfile) ProfiledUser {
	p := ProfiledUser{
		User: *u,
	}
	p.Profile = profile
	return p
}