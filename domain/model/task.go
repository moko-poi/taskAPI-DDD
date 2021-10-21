package model

import "errors"

// Task task の　構造体
type Task struct {
	ID      int
	Title   string
	Content string
}


//NewTask　task の コンストラクタ
func NewTask(title, content string) (*Task, error) {
	if title == "" {
		return nil, errors.New("titleを入力してください")
	}

	task := &Task{
		Title: title,
		Content: content,
	}

	return task, nil
}

//Set task のセッター
func (t *Task) Set(title, content string) (error) {
	if title == "" {
		return errors.New("titleを入力してください")
	}

	t.Title = title
	t.Content = content

	return nil
}

