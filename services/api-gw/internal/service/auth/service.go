package authservice

// type Service interface {
// 	LoginUserService(user *protobuf.CreateUserRequest) (*string, error)
// 	Close()
// }

// type AuthService struct {
// 	clientGRPC protobuf.UserServiceClient
// 	conn       *grpc.ClientConn
// }

// func New() *AuthService {
// 	conn, err := grpc.NewClient(config.USERS_ADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		logrus.Fatalf("did not connect: %v", err)
// 	}
// 	client := protobuf.NewUserServiceClient(conn)
// 	return &AuthService{
// 		clientGRPC: client,
// 		conn:       conn,
// 	}
// }

// func (s *AuthService) RegisterUserService(user *protobuf.CreateUserRequest) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	newUser, err := jsonutil.DecodeJson[protobuf.CreateUserRequest](r)
// 	if err != nil {
// 		jsonutil.EncodeJson(w, 400, "Invalid Credentails")
// 		return
// 	}
// 	response, err := s.clientGRPC.CreateUser(ctx, &newUser)
// 	if err != nil {
// 		status, msg := grpcutil.GRPCToHTTPError(err)
// 		jsonutil.EncodeJson(w, status, msg)
// 		return
// 	}
// 	if response.GetSuccess() {
// 		jsonutil.EncodeJson(w, 201, fmt.Sprintf("Registered user with id=%v", response.GetId()))
// 	} else {
// 		jsonutil.EncodeJson(w, 500, "Internal Server Error")
// 	}
// }

// func (s *AuthService) LoginUserService(user *protobuf.CreateUserRequest) (*string, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	existingUser, err := s.clientGRPC.GetUserByEmail(ctx, &protobuf.GetUserByEmailRequest{Email: user.GetEmail()})
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = bcrypt.CompareHashAndPassword([]byte(existingUser.GetUser().GetPassword()), []byte(user.GetPassword()))
// 	if err != nil {
// 		return nil, status.Errorf(codes.Unauthenticated, "Wrong password")
// 	}
// 	token, err := jwtutil.GenerateToken(int(existingUser.GetUser().GetId()), existingUser.GetUser().GetIsAdmin(), existingUser.GetUser().GetIsVerified())
// 	if err != nil {
// 		return nil, status.Errorf(codes.InvalidArgument, "Incorrect Credentials")
// 	}
// 	return token, nil
// }

// func (s *AuthService) Close() {
// 	err := s.conn.Close()
// 	if err != nil {
// 		logrus.WithError(err).Error("Couldn't close users service grpc connection")
// 	} else {
// 		logrus.Info("users service grpc conn closed")
// 	}
// }
