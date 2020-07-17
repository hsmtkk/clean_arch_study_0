package user

type User struct {
	id        int
	firstName string
	lastName  string
}

func New(id int, firstName, lastName string) User {
	return User{id: id, firstName: firstName, lastName: lastName}
}

func (u User) ID() int {
	return u.id
}

func (u User) FirstName() string {
	return u.firstName
}

func (u User) LastName() string {
	return u.lastName
}
