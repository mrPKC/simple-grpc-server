package user

import (
	userModel "github.com/Pratyush/User/model"
	pb "github.com/Pratyush/User/user/protos"
)

func GetUserById(ids []int32) *pb.GetUserByIdResponse {
	var response []*pb.User
	if ids != nil {
		for _, id := range ids {
			_, user, err := userModel.GetUser(&id)
			if err != nil {
				continue
			}
			response = append(response, user)
		}
	}
	return &pb.GetUserByIdResponse{User: response}
}

func GetUser() *pb.GetUsersResponse {
	var response []*pb.User
	users, _, err := userModel.GetUser(nil)
	if err != nil {
		return nil
	}
	response = append(response, users...)

	return &pb.GetUsersResponse{Users: response}
}

func GetUserByIdStream(id int32) *pb.GetUserStreamResponse {
	_, user, err := userModel.GetUser(&id)
	if err != nil {
		return nil
	}
	return &pb.GetUserStreamResponse{User: user}
}
