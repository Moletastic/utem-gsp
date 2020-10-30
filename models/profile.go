package models

type User struct {
	ID        int64   `json:"id"`
	AccountID int64   `json:"account_id"`
	Account   Account `json:"account"`
}

func NewUser(ac *Account) User {
	u := User{
		Account: *ac,
	}
	return u
}

func (u User) Bind(v interface{}) {
	v = User{}
}

func (u *User) GetID() int64 {
	return u.ID
}

func (u *User) GetUID() int64 {
	return u.ID
}

type IProfile interface {
	GetAccount() Account
}

type TeacherUser struct {
	User
	EntryYear int `json:"entry_year"`
}

func NewTeacherUser(t *Teacher) TeacherUser {
	p := TeacherUser{
		EntryYear: t.EntryYear,
	}
	p.ID = t.ID
	p.Account = t.Account
	p.AccountID = t.Account.ID
	return p
}

func (t *TeacherUser) GetAccount() Account {
	return t.Account
}

type AdminUser struct {
	User
	EntryYear int `json:"entry_year"`
}

func NewAdminUser(a *Admin) AdminUser {
	p := AdminUser{
		EntryYear: a.EntryYear,
	}
	p.ID = a.ID
	p.Account = a.Account
	p.AccountID = a.Account.ID
	return p
}

func (a *AdminUser) GetAccount() Account {
	return a.Account
}
