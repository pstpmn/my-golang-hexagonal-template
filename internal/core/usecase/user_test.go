package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/pstpmn/my-golang-hexagonal-template/internal/core/domain"
	"github.com/pstpmn/my-golang-hexagonal-template/internal/core/port"
	"github.com/pstpmn/my-golang-hexagonal-template/internal/mocks"
	"github.com/stretchr/testify/mock"
)

func Test_user_GetAll(t *testing.T) {
	mockUsers := []domain.User{
		{ID: "8e6c6952-fcfe-4ffd-86b4-ce23c37538f0", Name: "Mr.Math Matic", Email: "math@fake.com", IsActive: true},
		{ID: "294bd4ab-4e9c-482f-ba73-e2ab3d20efe7", Name: "Mr.Eloy Musk", Email: "eloy@fake.com", IsActive: true},
	}

	type fields struct {
		userRepo port.IUserRepo
	}
	type args struct {
		pctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.User
		wantErr bool
	}{
		{
			name: "should return success",
			fields: fields{
				userRepo: func() port.IUserRepo {
					mockRepo := &mocks.IUserRepo{}
					mockRepo.On("FindAll", mock.Anything).Return(mockUsers, nil)
					return mockRepo
				}(),
			},
			args: args{
				pctx: context.Background(),
			},
			want:    mockUsers,
			wantErr: false,
		},
		{
			name: "should return empty list when no users exist",
			fields: fields{
				userRepo: func() port.IUserRepo {
					mockRepo := &mocks.IUserRepo{}
					mockRepo.On("FindAll", mock.Anything).Return([]domain.User{}, nil)
					return mockRepo
				}(),
			},
			args: args{
				pctx: context.Background(),
			},
			want:    []domain.User{},
			wantErr: false,
		},
		{
			name: "should return error when repository fails",
			fields: fields{
				userRepo: func() port.IUserRepo {
					mockRepo := &mocks.IUserRepo{}
					mockRepo.On("FindAll", mock.Anything).Return(nil, errors.New("database connection failed"))
					return mockRepo
				}(),
			},
			args: args{
				pctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return partial user list",
			fields: fields{
				userRepo: func() port.IUserRepo {
					partialUsers := []domain.User{mockUsers[0]}
					mockRepo := &mocks.IUserRepo{}
					mockRepo.On("FindAll", mock.Anything).Return(partialUsers, nil)
					return mockRepo
				}(),
			},
			args: args{
				pctx: context.Background(),
			},
			want:    []domain.User{mockUsers[0]},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &u{
				userRepo: tt.fields.userRepo,
			}
			got, err := u.GetAll(tt.args.pctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("user.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_u_GetUser(t *testing.T) {
	mockUser := &domain.User{
		ID:       "8e6c6952-fcfe-4ffd-86b4-ce23c37538f0",
		Name:     "Mr.Math Matic",
		Email:    "math@fake.com",
		IsActive: true,
	}

	type fields struct {
		userRepo port.IUserRepo
	}
	type args struct {
		pctx   context.Context
		userId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.User
		wantErr bool
	}{
		{
			name: "should return success",
			fields: fields{
				userRepo: func() port.IUserRepo {
					mockRepo := &mocks.IUserRepo{}
					mockRepo.On("FindOneById", mock.Anything, "8e6c6952-fcfe-4ffd-86b4-ce23c37538f0").
						Return(mockUser, nil)
					return mockRepo
				}(),
			},
			args: args{
				pctx:   context.Background(),
				userId: "8e6c6952-fcfe-4ffd-86b4-ce23c37538f0",
			},
			want:    mockUser,
			wantErr: false,
		},
		{
			name: "should return nil when user not found",
			fields: fields{
				userRepo: func() port.IUserRepo {
					mockRepo := &mocks.IUserRepo{}
					mockRepo.On("FindOneById", mock.Anything, "non-existent-id").
						Return(nil, domain.ErrUserNotFound)
					return mockRepo
				}(),
			},
			args: args{
				pctx:   context.Background(),
				userId: "non-existent-id",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return error when repository fails",
			fields: fields{
				userRepo: func() port.IUserRepo {
					mockRepo := &mocks.IUserRepo{}
					mockRepo.On("FindOneById", mock.Anything, "8e6c6952-fcfe-4ffd-86b4-ce23c37538f0").
						Return(nil, errors.New("database connection failed"))
					return mockRepo
				}(),
			},
			args: args{
				pctx:   context.Background(),
				userId: "8e6c6952-fcfe-4ffd-86b4-ce23c37538f0",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return inactive user",
			fields: fields{
				userRepo: func() port.IUserRepo {
					inactiveUser := &domain.User{
						ID:       "294bd4ab-4e9c-482f-ba73-e2ab3d20efe7",
						Name:     "Mr.Eloy Musk",
						Email:    "eloy@fake.com",
						IsActive: false,
					}
					mockRepo := &mocks.IUserRepo{}
					mockRepo.On("FindOneById", mock.Anything, "294bd4ab-4e9c-482f-ba73-e2ab3d20efe7").
						Return(inactiveUser, nil)
					return mockRepo
				}(),
			},
			args: args{
				pctx:   context.Background(),
				userId: "294bd4ab-4e9c-482f-ba73-e2ab3d20efe7",
			},
			want: &domain.User{
				ID:       "294bd4ab-4e9c-482f-ba73-e2ab3d20efe7",
				Name:     "Mr.Eloy Musk",
				Email:    "eloy@fake.com",
				IsActive: false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &u{
				userRepo: tt.fields.userRepo,
			}
			got, err := u.GetUser(tt.args.pctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("u.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("u.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
