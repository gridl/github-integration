package user_manager_client

import "fmt"

const (
	userUrlStr = "/user"
)

type UserService struct {
	client *Client
}

func (u *UserService) Create(user User) (error) {
	fmt.Printf("Client: %+v", u.client)

	req, err := u.client.NewRequest(postMethod, userUrlStr, u)
	if err != nil {
		return err
	}

	_, err = u.client.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}

type User struct {
	Username *string  `json:"username"`
}

func NewUser (username string) *User {
	return &User{
		Username: &username,
	}
}