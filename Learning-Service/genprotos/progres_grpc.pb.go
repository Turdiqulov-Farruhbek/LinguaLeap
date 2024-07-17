// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: progres.proto

package genprotos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ProgressService_SetGoal_FullMethodName             = "/progress.ProgressService/SetGoal"
	ProgressService_UpdateGoal_FullMethodName          = "/progress.ProgressService/UpdateGoal"
	ProgressService_GetLanguageProgress_FullMethodName = "/progress.ProgressService/GetLanguageProgress"
	ProgressService_GetDailyProgress_FullMethodName    = "/progress.ProgressService/GetDailyProgress"
	ProgressService_GetWeeklyProgress_FullMethodName   = "/progress.ProgressService/GetWeeklyProgress"
	ProgressService_GetMonthlyProgress_FullMethodName  = "/progress.ProgressService/GetMonthlyProgress"
	ProgressService_GetAchievements_FullMethodName     = "/progress.ProgressService/GetAchievements"
	ProgressService_GetLeaderboard_FullMethodName      = "/progress.ProgressService/GetLeaderboard"
	ProgressService_GetSkills_FullMethodName           = "/progress.ProgressService/GetSkills"
	ProgressService_GetGoals_FullMethodName            = "/progress.ProgressService/GetGoals"
	ProgressService_GetStatistics_FullMethodName       = "/progress.ProgressService/GetStatistics"
	ProgressService_DeleteGoal_FullMethodName          = "/progress.ProgressService/DeleteGoal"
)

// ProgressServiceClient is the client API for ProgressService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProgressServiceClient interface {
	SetGoal(ctx context.Context, in *Goal, opts ...grpc.CallOption) (*GoalResponse, error)
	UpdateGoal(ctx context.Context, in *UpdateGoalRequest, opts ...grpc.CallOption) (*GoalResponse, error)
	GetLanguageProgress(ctx context.Context, in *GetLanguageProgressRequest, opts ...grpc.CallOption) (*ProgressResponse, error)
	GetDailyProgress(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*DailyProgressResponse, error)
	GetWeeklyProgress(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*WeeklyProgressResponse, error)
	GetMonthlyProgress(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*MonthlyProgressResponse, error)
	GetAchievements(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*AchievementsResponse, error)
	GetLeaderboard(ctx context.Context, in *GetLeaderboardRequest, opts ...grpc.CallOption) (*LeaderboardResponse, error)
	GetSkills(ctx context.Context, in *GetSkillsRequest, opts ...grpc.CallOption) (*SkillsResponse, error)
	GetGoals(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GoalsResponse, error)
	GetStatistics(ctx context.Context, in *GetStatisticsRequest, opts ...grpc.CallOption) (*StatisticsResponse, error)
	DeleteGoal(ctx context.Context, in *DeleteGoalRequest, opts ...grpc.CallOption) (*GoalResponse, error)
}

type progressServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProgressServiceClient(cc grpc.ClientConnInterface) ProgressServiceClient {
	return &progressServiceClient{cc}
}

func (c *progressServiceClient) SetGoal(ctx context.Context, in *Goal, opts ...grpc.CallOption) (*GoalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GoalResponse)
	err := c.cc.Invoke(ctx, ProgressService_SetGoal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressServiceClient) UpdateGoal(ctx context.Context, in *UpdateGoalRequest, opts ...grpc.CallOption) (*GoalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GoalResponse)
	err := c.cc.Invoke(ctx, ProgressService_UpdateGoal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressServiceClient) GetLanguageProgress(ctx context.Context, in *GetLanguageProgressRequest, opts ...grpc.CallOption) (*ProgressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProgressResponse)
	err := c.cc.Invoke(ctx, ProgressService_GetLanguageProgress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressServiceClient) GetDailyProgress(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*DailyProgressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DailyProgressResponse)
	err := c.cc.Invoke(ctx, ProgressService_GetDailyProgress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressServiceClient) GetWeeklyProgress(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*WeeklyProgressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WeeklyProgressResponse)
	err := c.cc.Invoke(ctx, ProgressService_GetWeeklyProgress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressServiceClient) GetMonthlyProgress(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*MonthlyProgressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MonthlyProgressResponse)
	err := c.cc.Invoke(ctx, ProgressService_GetMonthlyProgress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressServiceClient) GetAchievements(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*AchievementsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AchievementsResponse)
	err := c.cc.Invoke(ctx, ProgressService_GetAchievements_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressServiceClient) GetLeaderboard(ctx context.Context, in *GetLeaderboardRequest, opts ...grpc.CallOption) (*LeaderboardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LeaderboardResponse)
	err := c.cc.Invoke(ctx, ProgressService_GetLeaderboard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressServiceClient) GetSkills(ctx context.Context, in *GetSkillsRequest, opts ...grpc.CallOption) (*SkillsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SkillsResponse)
	err := c.cc.Invoke(ctx, ProgressService_GetSkills_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressServiceClient) GetGoals(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GoalsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GoalsResponse)
	err := c.cc.Invoke(ctx, ProgressService_GetGoals_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressServiceClient) GetStatistics(ctx context.Context, in *GetStatisticsRequest, opts ...grpc.CallOption) (*StatisticsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatisticsResponse)
	err := c.cc.Invoke(ctx, ProgressService_GetStatistics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressServiceClient) DeleteGoal(ctx context.Context, in *DeleteGoalRequest, opts ...grpc.CallOption) (*GoalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GoalResponse)
	err := c.cc.Invoke(ctx, ProgressService_DeleteGoal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProgressServiceServer is the server API for ProgressService service.
// All implementations must embed UnimplementedProgressServiceServer
// for forward compatibility
type ProgressServiceServer interface {
	SetGoal(context.Context, *Goal) (*GoalResponse, error)
	UpdateGoal(context.Context, *UpdateGoalRequest) (*GoalResponse, error)
	GetLanguageProgress(context.Context, *GetLanguageProgressRequest) (*ProgressResponse, error)
	GetDailyProgress(context.Context, *GetUserRequest) (*DailyProgressResponse, error)
	GetWeeklyProgress(context.Context, *GetUserRequest) (*WeeklyProgressResponse, error)
	GetMonthlyProgress(context.Context, *GetUserRequest) (*MonthlyProgressResponse, error)
	GetAchievements(context.Context, *GetUserRequest) (*AchievementsResponse, error)
	GetLeaderboard(context.Context, *GetLeaderboardRequest) (*LeaderboardResponse, error)
	GetSkills(context.Context, *GetSkillsRequest) (*SkillsResponse, error)
	GetGoals(context.Context, *GetUserRequest) (*GoalsResponse, error)
	GetStatistics(context.Context, *GetStatisticsRequest) (*StatisticsResponse, error)
	DeleteGoal(context.Context, *DeleteGoalRequest) (*GoalResponse, error)
	mustEmbedUnimplementedProgressServiceServer()
}

// UnimplementedProgressServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProgressServiceServer struct {
}

func (UnimplementedProgressServiceServer) SetGoal(context.Context, *Goal) (*GoalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGoal not implemented")
}
func (UnimplementedProgressServiceServer) UpdateGoal(context.Context, *UpdateGoalRequest) (*GoalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGoal not implemented")
}
func (UnimplementedProgressServiceServer) GetLanguageProgress(context.Context, *GetLanguageProgressRequest) (*ProgressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLanguageProgress not implemented")
}
func (UnimplementedProgressServiceServer) GetDailyProgress(context.Context, *GetUserRequest) (*DailyProgressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDailyProgress not implemented")
}
func (UnimplementedProgressServiceServer) GetWeeklyProgress(context.Context, *GetUserRequest) (*WeeklyProgressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeeklyProgress not implemented")
}
func (UnimplementedProgressServiceServer) GetMonthlyProgress(context.Context, *GetUserRequest) (*MonthlyProgressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMonthlyProgress not implemented")
}
func (UnimplementedProgressServiceServer) GetAchievements(context.Context, *GetUserRequest) (*AchievementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAchievements not implemented")
}
func (UnimplementedProgressServiceServer) GetLeaderboard(context.Context, *GetLeaderboardRequest) (*LeaderboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeaderboard not implemented")
}
func (UnimplementedProgressServiceServer) GetSkills(context.Context, *GetSkillsRequest) (*SkillsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSkills not implemented")
}
func (UnimplementedProgressServiceServer) GetGoals(context.Context, *GetUserRequest) (*GoalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoals not implemented")
}
func (UnimplementedProgressServiceServer) GetStatistics(context.Context, *GetStatisticsRequest) (*StatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatistics not implemented")
}
func (UnimplementedProgressServiceServer) DeleteGoal(context.Context, *DeleteGoalRequest) (*GoalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGoal not implemented")
}
func (UnimplementedProgressServiceServer) mustEmbedUnimplementedProgressServiceServer() {}

// UnsafeProgressServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProgressServiceServer will
// result in compilation errors.
type UnsafeProgressServiceServer interface {
	mustEmbedUnimplementedProgressServiceServer()
}

func RegisterProgressServiceServer(s grpc.ServiceRegistrar, srv ProgressServiceServer) {
	s.RegisterService(&ProgressService_ServiceDesc, srv)
}

func _ProgressService_SetGoal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Goal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).SetGoal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_SetGoal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).SetGoal(ctx, req.(*Goal))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressService_UpdateGoal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGoalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).UpdateGoal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_UpdateGoal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).UpdateGoal(ctx, req.(*UpdateGoalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressService_GetLanguageProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLanguageProgressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).GetLanguageProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_GetLanguageProgress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).GetLanguageProgress(ctx, req.(*GetLanguageProgressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressService_GetDailyProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).GetDailyProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_GetDailyProgress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).GetDailyProgress(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressService_GetWeeklyProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).GetWeeklyProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_GetWeeklyProgress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).GetWeeklyProgress(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressService_GetMonthlyProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).GetMonthlyProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_GetMonthlyProgress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).GetMonthlyProgress(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressService_GetAchievements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).GetAchievements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_GetAchievements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).GetAchievements(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressService_GetLeaderboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLeaderboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).GetLeaderboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_GetLeaderboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).GetLeaderboard(ctx, req.(*GetLeaderboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressService_GetSkills_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSkillsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).GetSkills(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_GetSkills_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).GetSkills(ctx, req.(*GetSkillsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressService_GetGoals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).GetGoals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_GetGoals_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).GetGoals(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressService_GetStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatisticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).GetStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_GetStatistics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).GetStatistics(ctx, req.(*GetStatisticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressService_DeleteGoal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGoalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).DeleteGoal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_DeleteGoal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).DeleteGoal(ctx, req.(*DeleteGoalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProgressService_ServiceDesc is the grpc.ServiceDesc for ProgressService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProgressService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "progress.ProgressService",
	HandlerType: (*ProgressServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetGoal",
			Handler:    _ProgressService_SetGoal_Handler,
		},
		{
			MethodName: "UpdateGoal",
			Handler:    _ProgressService_UpdateGoal_Handler,
		},
		{
			MethodName: "GetLanguageProgress",
			Handler:    _ProgressService_GetLanguageProgress_Handler,
		},
		{
			MethodName: "GetDailyProgress",
			Handler:    _ProgressService_GetDailyProgress_Handler,
		},
		{
			MethodName: "GetWeeklyProgress",
			Handler:    _ProgressService_GetWeeklyProgress_Handler,
		},
		{
			MethodName: "GetMonthlyProgress",
			Handler:    _ProgressService_GetMonthlyProgress_Handler,
		},
		{
			MethodName: "GetAchievements",
			Handler:    _ProgressService_GetAchievements_Handler,
		},
		{
			MethodName: "GetLeaderboard",
			Handler:    _ProgressService_GetLeaderboard_Handler,
		},
		{
			MethodName: "GetSkills",
			Handler:    _ProgressService_GetSkills_Handler,
		},
		{
			MethodName: "GetGoals",
			Handler:    _ProgressService_GetGoals_Handler,
		},
		{
			MethodName: "GetStatistics",
			Handler:    _ProgressService_GetStatistics_Handler,
		},
		{
			MethodName: "DeleteGoal",
			Handler:    _ProgressService_DeleteGoal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "progres.proto",
}
