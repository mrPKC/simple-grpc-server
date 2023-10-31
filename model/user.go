package user

import (
	pb "github.com/Pratyush/User/user/protos"
)

func GetUser(id *int32) ([]*pb.User, *pb.User, *string) {
	users := map[int32]pb.User{
		1: pb.User{
			Id:      1,
			Fname:   "Anonymous 1",
			Phone:   "Anonymous 1",
			Height:  "5' 10",
			City:    "Delhi",
			Married: false,
		},
		2: pb.User{
			Id:      2,
			Fname:   "Anonymous 2",
			Phone:   "Anonymous 2",
			Height:  "5' 10",
			City:    "Delhi",
			Married: false,
		},
		3: pb.User{
			Id:      3,
			Fname:   "Anonymous 3",
			Phone:   "Anonymous 3",
			Height:  "5' 10",
			City:    "Delhi",
			Married: false,
		},
		4: pb.User{
			Id:      4,
			Fname:   "Anonymous 4",
			Phone:   "Anonymous 4",
			Height:  "5' 10",
			City:    "Delhi",
			Married: false,
		},
		5: pb.User{
			Id:      5,
			Fname:   "Anonymous 5",
			Phone:   "Anonymous 5",
			Height:  "5' 10",
			City:    "Delhi",
			Married: false,
		},
		6: pb.User{
			Id:      6,
			Fname:   "Anonymous 6",
			Phone:   "Anonymous 6",
			Height:  "5' 10",
			City:    "Delhi",
			Married: false,
		},
	}

	if id != nil {
		if user, ok := users[*id]; ok {
			return nil, &user, nil
		}
		err := "Invalid User ID"
		return nil, nil, &err
	}

	var response []*pb.User
	for _, value := range users {
		var user pb.User
		user = value
		response = append(response, &user)
	}
	return response, nil, nil
}
