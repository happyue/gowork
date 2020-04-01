package service

import (
	"sync"

	"gowork/log"
	"gowork/model"
)

// Todo service.
var Todo = &todoService{
	mutex: &sync.Mutex{},
}

type todoService struct {
	mutex *sync.Mutex
}

func (srv *todoService) GetTodoList(size int, userID uint) (ret []*model.Todo, err error) {
	if err := db.Where("`user_id` = ?", userID).Order("`id` ASC").Limit(size).Find(&ret).Error; err != nil {
		log.Errorf("get todo list failed: " + err.Error())
		return ret, err
	}

	return ret, nil
}

func (srv *todoService) DeleTodo(todo model.Todo) error {

	// log.Debugf("dele todo [%#v]", todo)
	if err := db.Where("`user_id` = ? AND `id` = ?", todo.UserID, todo.ID).Delete(&todo).Error; err != nil {
		return err
	}

	return nil
}

func (srv *todoService) AddTodo(todo model.Todo) error {

	// log.Debugf("add todo [%#v]", todo)
	if err := db.Create(&todo).Error; err != nil {
		return err
	}

	return nil
}

func (srv *todoService) UpdateTodo(todo model.Todo) error {
	// 	if err := db.Where("`id2` = ? AND `type` = ? AND `blog_id` = ?", archiveID, model.CorrelationArticleArchive, blogID).

	// log.Debugf("update todo [%#v]", todo)
	if err := db.Model(&todo).Where("`user_id` = ? AND `id` = ?", todo.UserID, todo.ID).Updates(map[string]interface{}{"text": todo.Text, "done": todo.Done}).Error; err != nil {
		return err
	}

	return nil
}
