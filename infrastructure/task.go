package infrastructure

import (
	"github.com/jinzhu/gorm"
	"go-layered-architecture/domain/model"
	"go-layered-architecture/domain/repository"
)

// TaskRepository task repositoryの構造体
type TaskRepository struct {
	Conn *gorm.DB
}

// NewTaskRepository task repositoryのコンストラクタ
func NewTaskRepository(conn *gorm.DB) repository.TaskRepository {
	return &TaskRepository{Conn: conn}
}

// Create taskの保存
func (tr *TaskRepository) Create(task *model.Task) (*model.Task, error) {
	err := tr.Conn.Create(&task).Error
	if err != nil {
		return  nil, err
	}

	return task, nil
}

// FindByID taskをIDで取得
func (tr *TaskRepository) FindByID(id int) (*model.Task, error) {
	task := &model.Task{ID: id}

	err := tr.Conn.First(&task).Error
	if err != nil {
		return nil, err
	}

	return task, nil
}

// Update taskの更新
func (tr *TaskRepository) Update(task *model.Task) (*model.Task, error) {
	err := tr.Conn.Model(&task).Update(&task).Error
	if err != nil {
		return  nil, err
	}

	return  task, nil
}

// Delete taskの削除
func (tr *TaskRepository) Delete(task *model.Task) error {
	err := tr.Conn.Delete(&task).Error

	if err != nil {
		return err
	}

	return  nil
}
