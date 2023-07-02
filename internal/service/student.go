package service

import (
	"context"

	pb "quickstart/api/helloworld/v1"
	"quickstart/internal/biz"
)

type StudentService struct {
	pb.UnimplementedStudentServer
	uc *biz.StudentUsercase
}

func NewStudentService(uc *biz.StudentUsercase) *StudentService {
	return &StudentService{uc: uc}
}

func (s *StudentService) CreateStudent(ctx context.Context, req *pb.CreateStudentRequest) (*pb.CreateStudentReply, error) {
	return &pb.CreateStudentReply{}, nil
}
func (s *StudentService) UpdateStudent(ctx context.Context, req *pb.UpdateStudentRequest) (*pb.UpdateStudentReply, error) {
	return &pb.UpdateStudentReply{}, nil
}
func (s *StudentService) DeleteStudent(ctx context.Context, req *pb.DeleteStudentRequest) (*pb.DeleteStudentReply, error) {
	return &pb.DeleteStudentReply{}, nil
}
func (s *StudentService) GetStudent(ctx context.Context, req *pb.GetStudentRequest) (*pb.GetStudentReply, error) {
	return &pb.GetStudentReply{}, nil
}
func (s *StudentService) ListStudent(ctx context.Context, req *pb.ListStudentRequest) (*pb.ListStudentReply, error) {
	return &pb.ListStudentReply{}, nil
}
func (s *StudentService) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{}, nil
}
