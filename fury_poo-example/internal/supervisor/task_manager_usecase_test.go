package supervisor

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"users_example/internal/supervisor/developer"
	"users_example/internal/supervisor/developer/task"
)

func Test_useCaseTaskManagement_Schedule(t *testing.T) {
	type fields struct {
		taskPublisher task.Publisher
		devRepo       developer.Repository
	}
	type args struct {
		id   string
		task task.Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "test-fail-get-dev",
			fields: fields{
				taskPublisher: nil,
				devRepo: testDevRepo{
					mockGet: func(id string) (*developer.Developer, error) {
						return nil, errors.New("error al hacer un get desde el repo de devs")
					},
				},
			},
			args: args{
				id: "test-1",
				task: task.Task{
					Name:      "random_for_test",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					Status:    task.StatusPending,
				},
			},
			wantErr: true,
			err:     errors.New("error al hacer un get desde el repo de devs"),
		},
		{
			name: "test-fail-get-dev-not_exists",
			fields: fields{
				taskPublisher: nil,
				devRepo: testDevRepo{
					mockGet: func(id string) (*developer.Developer, error) {
						return nil, nil
					},
				},
			},
			args: args{
				id: "test-1",
				task: task.Task{
					Name:      "random_for_test",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					Status:    task.StatusPending,
				},
			},
			wantErr: true,
			err:     errors.New("user_id:test-1 not found, task cannot be scheduled"),
		},
		{
			name: "test-fail-save",
			fields: fields{
				taskPublisher: nil,
				devRepo: testDevRepo{
					mockGet: func(id string) (*developer.Developer, error) {
						return &developer.Developer{
							ID:        "test-1",
							Name:      "random-test",
							Team:      "random-team",
							Seniority: developer.SemiSenior,
						}, nil
					},
					mockSave: func(d *developer.Developer) error {
						return errors.New("error guardando el developer con su nueva tarea")
					},
				},
			},
			args: args{
				id: "test-1",
				task: task.Task{
					Name:      "random_for_test",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					Status:    task.StatusPending,
				},
			},
			wantErr: true,
			err:     errors.New("error guardando el developer con su nueva tarea"),
		},
		{
			name: "test-fail-success",
			fields: fields{
				taskPublisher: nil,
				devRepo: testDevRepo{
					mockGet: func(id string) (*developer.Developer, error) {
						return &developer.Developer{
							ID:        "test-1",
							Name:      "random-test",
							Team:      "random-team",
							Seniority: developer.SemiSenior,
						}, nil
					},
					mockSave: func(d *developer.Developer) error {
						return nil
					},
				},
			},
			args: args{
				id: "test-1",
				task: task.Task{
					Name:      "random_for_test",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					Status:    task.StatusPending,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := useCaseTaskManagement{
				taskPublisher: tt.fields.taskPublisher,
				devRepo:       tt.fields.devRepo,
			}
			if _, err := u.ScheduleTask(tt.args.id, ""); (err != nil) != tt.wantErr {
				t.Errorf("ScheduleTask() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.err, err)
			}
		})
	}
}

type testPublisher struct {
	mockPublish func(ownerId string, t task.Task) error
}

func (t2 testPublisher) Publish(ownerId string, t task.Task) error {
	return t2.mockPublish(ownerId, t)
}

type testDevRepo struct {
	mockSave func(d *developer.Developer) error
	mockGet  func(id string) (*developer.Developer, error)
}

func (t testDevRepo) Save(d *developer.Developer) error {
	return t.mockSave(d)
}

func (t testDevRepo) Get(id string) (*developer.Developer, error) {
	return t.mockGet(id)
}

func (t testDevRepo) Delete(id string) error {
	panic("implement me")
}

func (t testDevRepo) SearchByStatus(status task.Status) ([]developer.Developer, error) {
	panic("implement me")
}
