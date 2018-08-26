package dao

type UserInfo struct {
    id        string
    name      string
    age       string
}

func (u *UserInfo) GetId() string {
    return u.id
}

func (u *UserInfo) SetId(id string) {
    u.id = id
}

func (u *UserInfo) GetName() string {
    return u.name
}

func (u *UserInfo) SetName(name string) {
    u.name = name
}

func (u *UserInfo) GetAge() string {
    return u.age
}

func (u *UserInfo) SetAge(age string) {
    u.age = age
}