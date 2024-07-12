// Package handler implements handler of worker
package handler

import (
	"context"
	"encoding/json"
	"swclabs/swipecore/internal/core/service/accountmanagement"

	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/lib/worker"

	"github.com/hibiken/asynq"
)

type AccountManagement struct {
	accountmanagement.Task
	handler accountmanagement.IAccountManagement
}

func NewAccountManagementConsume(handler accountmanagement.IAccountManagement) *AccountManagement {
	return &AccountManagement{
		handler: handler,
	}
}

func (account *AccountManagement) HandleSignUp() (string, worker.HandleFunc) {
	return worker.GetTaskName(account.SignUp),
		func(_ context.Context, task *asynq.Task) error {
			var data domain.SignUpSchema
			if err := json.Unmarshal(task.Payload(), &data); err != nil {
				return err
			}
			return account.handler.SignUp(context.Background(), data)
		}
}

func (account *AccountManagement) HandleOAuth2SaveUser() (string, worker.HandleFunc) {
	return worker.GetTaskName(account.OAuth2SaveUser),
		func(_ context.Context, task *asynq.Task) error {
			var data domain.OAuth2SaveUser
			if err := json.Unmarshal(task.Payload(), &data); err != nil {
				return err
			}
			return account.handler.OAuth2SaveUser(context.Background(), data)
		}
}

func (account *AccountManagement) HandleUpdateUserInfo() (string, worker.HandleFunc) {
	return worker.GetTaskName(account.UpdateUserInfo),
		func(_ context.Context, task *asynq.Task) error {
			var data domain.UserUpdate
			if err := json.Unmarshal(task.Payload(), &data); err != nil {
				return err
			}
			return account.handler.UpdateUserInfo(context.Background(), data)
		}
}
