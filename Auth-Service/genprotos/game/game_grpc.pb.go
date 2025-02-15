// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: game.proto

package game

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
	GameService_GetLevels_FullMethodName         = "/game.GameService/GetLevels"
	GameService_StartLevel_FullMethodName        = "/game.GameService/StartLevel"
	GameService_CompleteLevel_FullMethodName     = "/game.GameService/CompleteLevel"
	GameService_GetChallenge_FullMethodName      = "/game.GameService/GetChallenge"
	GameService_SubmitChallenge_FullMethodName   = "/game.GameService/SubmitChallenge"
	GameService_GetSimulations_FullMethodName    = "/game.GameService/GetSimulations"
	GameService_GetSimulation_FullMethodName     = "/game.GameService/GetSimulation"
	GameService_RunSimulation_FullMethodName     = "/game.GameService/RunSimulation"
	GameService_GetLeaderboard_FullMethodName    = "/game.GameService/GetLeaderboard"
	GameService_GetAchievements_FullMethodName   = "/game.GameService/GetAchievements"
	GameService_GetQuantumCircuit_FullMethodName = "/game.GameService/GetQuantumCircuit"
)

// GameServiceClient is the client API for GameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameServiceClient interface {
	GetLevels(ctx context.Context, in *GetLevelsRequest, opts ...grpc.CallOption) (*GetLevelsResponse, error)
	StartLevel(ctx context.Context, in *StartLevelRequest, opts ...grpc.CallOption) (*StartLevelResponse, error)
	CompleteLevel(ctx context.Context, in *CompleteLevelRequest, opts ...grpc.CallOption) (*CompleteLevelResponse, error)
	GetChallenge(ctx context.Context, in *GetChallengeRequest, opts ...grpc.CallOption) (*Challenge, error)
	SubmitChallenge(ctx context.Context, in *SubmitChallengeRequest, opts ...grpc.CallOption) (*SubmitChallengeResponse, error)
	GetSimulations(ctx context.Context, in *GetSimulationsRequest, opts ...grpc.CallOption) (*GetSimulationsResponse, error)
	GetSimulation(ctx context.Context, in *GetSimulationRequest, opts ...grpc.CallOption) (*Simulation, error)
	RunSimulation(ctx context.Context, in *RunSimulationRequest, opts ...grpc.CallOption) (*SimulationResult, error)
	GetLeaderboard(ctx context.Context, in *GetLeaderboardRequest, opts ...grpc.CallOption) (*LeaderboardResponse, error)
	GetAchievements(ctx context.Context, in *GetAchievementsRequest, opts ...grpc.CallOption) (*AchievementsResponse, error)
	GetQuantumCircuit(ctx context.Context, in *GetQuantumCircuitRequest, opts ...grpc.CallOption) (*Circuit, error)
}

type gameServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGameServiceClient(cc grpc.ClientConnInterface) GameServiceClient {
	return &gameServiceClient{cc}
}

func (c *gameServiceClient) GetLevels(ctx context.Context, in *GetLevelsRequest, opts ...grpc.CallOption) (*GetLevelsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLevelsResponse)
	err := c.cc.Invoke(ctx, GameService_GetLevels_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) StartLevel(ctx context.Context, in *StartLevelRequest, opts ...grpc.CallOption) (*StartLevelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartLevelResponse)
	err := c.cc.Invoke(ctx, GameService_StartLevel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) CompleteLevel(ctx context.Context, in *CompleteLevelRequest, opts ...grpc.CallOption) (*CompleteLevelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CompleteLevelResponse)
	err := c.cc.Invoke(ctx, GameService_CompleteLevel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) GetChallenge(ctx context.Context, in *GetChallengeRequest, opts ...grpc.CallOption) (*Challenge, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Challenge)
	err := c.cc.Invoke(ctx, GameService_GetChallenge_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) SubmitChallenge(ctx context.Context, in *SubmitChallengeRequest, opts ...grpc.CallOption) (*SubmitChallengeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubmitChallengeResponse)
	err := c.cc.Invoke(ctx, GameService_SubmitChallenge_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) GetSimulations(ctx context.Context, in *GetSimulationsRequest, opts ...grpc.CallOption) (*GetSimulationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSimulationsResponse)
	err := c.cc.Invoke(ctx, GameService_GetSimulations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) GetSimulation(ctx context.Context, in *GetSimulationRequest, opts ...grpc.CallOption) (*Simulation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Simulation)
	err := c.cc.Invoke(ctx, GameService_GetSimulation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) RunSimulation(ctx context.Context, in *RunSimulationRequest, opts ...grpc.CallOption) (*SimulationResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SimulationResult)
	err := c.cc.Invoke(ctx, GameService_RunSimulation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) GetLeaderboard(ctx context.Context, in *GetLeaderboardRequest, opts ...grpc.CallOption) (*LeaderboardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LeaderboardResponse)
	err := c.cc.Invoke(ctx, GameService_GetLeaderboard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) GetAchievements(ctx context.Context, in *GetAchievementsRequest, opts ...grpc.CallOption) (*AchievementsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AchievementsResponse)
	err := c.cc.Invoke(ctx, GameService_GetAchievements_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) GetQuantumCircuit(ctx context.Context, in *GetQuantumCircuitRequest, opts ...grpc.CallOption) (*Circuit, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Circuit)
	err := c.cc.Invoke(ctx, GameService_GetQuantumCircuit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServiceServer is the server API for GameService service.
// All implementations must embed UnimplementedGameServiceServer
// for forward compatibility
type GameServiceServer interface {
	GetLevels(context.Context, *GetLevelsRequest) (*GetLevelsResponse, error)
	StartLevel(context.Context, *StartLevelRequest) (*StartLevelResponse, error)
	CompleteLevel(context.Context, *CompleteLevelRequest) (*CompleteLevelResponse, error)
	GetChallenge(context.Context, *GetChallengeRequest) (*Challenge, error)
	SubmitChallenge(context.Context, *SubmitChallengeRequest) (*SubmitChallengeResponse, error)
	GetSimulations(context.Context, *GetSimulationsRequest) (*GetSimulationsResponse, error)
	GetSimulation(context.Context, *GetSimulationRequest) (*Simulation, error)
	RunSimulation(context.Context, *RunSimulationRequest) (*SimulationResult, error)
	GetLeaderboard(context.Context, *GetLeaderboardRequest) (*LeaderboardResponse, error)
	GetAchievements(context.Context, *GetAchievementsRequest) (*AchievementsResponse, error)
	GetQuantumCircuit(context.Context, *GetQuantumCircuitRequest) (*Circuit, error)
	mustEmbedUnimplementedGameServiceServer()
}

// UnimplementedGameServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGameServiceServer struct {
}

func (UnimplementedGameServiceServer) GetLevels(context.Context, *GetLevelsRequest) (*GetLevelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLevels not implemented")
}
func (UnimplementedGameServiceServer) StartLevel(context.Context, *StartLevelRequest) (*StartLevelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartLevel not implemented")
}
func (UnimplementedGameServiceServer) CompleteLevel(context.Context, *CompleteLevelRequest) (*CompleteLevelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteLevel not implemented")
}
func (UnimplementedGameServiceServer) GetChallenge(context.Context, *GetChallengeRequest) (*Challenge, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChallenge not implemented")
}
func (UnimplementedGameServiceServer) SubmitChallenge(context.Context, *SubmitChallengeRequest) (*SubmitChallengeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitChallenge not implemented")
}
func (UnimplementedGameServiceServer) GetSimulations(context.Context, *GetSimulationsRequest) (*GetSimulationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSimulations not implemented")
}
func (UnimplementedGameServiceServer) GetSimulation(context.Context, *GetSimulationRequest) (*Simulation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSimulation not implemented")
}
func (UnimplementedGameServiceServer) RunSimulation(context.Context, *RunSimulationRequest) (*SimulationResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunSimulation not implemented")
}
func (UnimplementedGameServiceServer) GetLeaderboard(context.Context, *GetLeaderboardRequest) (*LeaderboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeaderboard not implemented")
}
func (UnimplementedGameServiceServer) GetAchievements(context.Context, *GetAchievementsRequest) (*AchievementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAchievements not implemented")
}
func (UnimplementedGameServiceServer) GetQuantumCircuit(context.Context, *GetQuantumCircuitRequest) (*Circuit, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuantumCircuit not implemented")
}
func (UnimplementedGameServiceServer) mustEmbedUnimplementedGameServiceServer() {}

// UnsafeGameServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServiceServer will
// result in compilation errors.
type UnsafeGameServiceServer interface {
	mustEmbedUnimplementedGameServiceServer()
}

func RegisterGameServiceServer(s grpc.ServiceRegistrar, srv GameServiceServer) {
	s.RegisterService(&GameService_ServiceDesc, srv)
}

func _GameService_GetLevels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLevelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).GetLevels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_GetLevels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).GetLevels(ctx, req.(*GetLevelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_StartLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).StartLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_StartLevel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).StartLevel(ctx, req.(*StartLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_CompleteLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).CompleteLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_CompleteLevel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).CompleteLevel(ctx, req.(*CompleteLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_GetChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChallengeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).GetChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_GetChallenge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).GetChallenge(ctx, req.(*GetChallengeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_SubmitChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitChallengeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).SubmitChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_SubmitChallenge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).SubmitChallenge(ctx, req.(*SubmitChallengeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_GetSimulations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSimulationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).GetSimulations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_GetSimulations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).GetSimulations(ctx, req.(*GetSimulationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_GetSimulation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSimulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).GetSimulation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_GetSimulation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).GetSimulation(ctx, req.(*GetSimulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_RunSimulation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunSimulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).RunSimulation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_RunSimulation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).RunSimulation(ctx, req.(*RunSimulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_GetLeaderboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLeaderboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).GetLeaderboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_GetLeaderboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).GetLeaderboard(ctx, req.(*GetLeaderboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_GetAchievements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAchievementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).GetAchievements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_GetAchievements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).GetAchievements(ctx, req.(*GetAchievementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_GetQuantumCircuit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuantumCircuitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).GetQuantumCircuit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_GetQuantumCircuit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).GetQuantumCircuit(ctx, req.(*GetQuantumCircuitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GameService_ServiceDesc is the grpc.ServiceDesc for GameService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "game.GameService",
	HandlerType: (*GameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLevels",
			Handler:    _GameService_GetLevels_Handler,
		},
		{
			MethodName: "StartLevel",
			Handler:    _GameService_StartLevel_Handler,
		},
		{
			MethodName: "CompleteLevel",
			Handler:    _GameService_CompleteLevel_Handler,
		},
		{
			MethodName: "GetChallenge",
			Handler:    _GameService_GetChallenge_Handler,
		},
		{
			MethodName: "SubmitChallenge",
			Handler:    _GameService_SubmitChallenge_Handler,
		},
		{
			MethodName: "GetSimulations",
			Handler:    _GameService_GetSimulations_Handler,
		},
		{
			MethodName: "GetSimulation",
			Handler:    _GameService_GetSimulation_Handler,
		},
		{
			MethodName: "RunSimulation",
			Handler:    _GameService_RunSimulation_Handler,
		},
		{
			MethodName: "GetLeaderboard",
			Handler:    _GameService_GetLeaderboard_Handler,
		},
		{
			MethodName: "GetAchievements",
			Handler:    _GameService_GetAchievements_Handler,
		},
		{
			MethodName: "GetQuantumCircuit",
			Handler:    _GameService_GetQuantumCircuit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "game.proto",
}
