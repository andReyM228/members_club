package service

import (
	"errors"
	"fmt"
	mc "members_club"
	"time"
)

func GetListOfMembers() *[]mc.Member {
	return mc.Members

}

func AddMember(name, email string) error {
	if err := chekEmail(email); err != nil {
		return err
	}
	e, m, d := time.Now().Date()
	date := fmt.Sprintf("%v.%v.%v", d, m, e)
	members := append(*mc.Members,
		mc.Member{
			Name:  name,
			Email: email,
			Date:  date})

	mc.Members = &members
	fmt.Println(name, email, date)
	return nil
}

func chekEmail(email string) error {
	for _, m := range *mc.Members {
		if m.Email == email {
			return errors.New(fmt.Sprintf("error Email"))
		}
	}
	return nil
}
