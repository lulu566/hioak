package gitlab

import (
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/hidevopsio/hioak/starter/scm"
	"github.com/jinzhu/copier"
)

type User struct {
	scm.User
	client UserInterface
}

func NewUser(c UserInterface) *User {
	return &User{
		client: c,
	}
}

func (s *User) GetUser(baseUrl, accessToken string) (*scm.User, error) {
	log.Debug("Session get user")
	user, _, err := s.client.CurrentUser()
	if err != nil {
		return nil, err
	}
	scmUser := &scm.User{}
	copier.Copy(scmUser, user)
	log.Debugf("get User : %v", user)
	return scmUser, nil

}
