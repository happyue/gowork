package controller

import (
	"math"
	"net/http"

	"gowork/model"
	"gowork/service"
	"gowork/utils/utsessions"

	"gowork/controller/common"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//GetTodoList 获取todolist
func GetTodoList(c *gin.Context) {
	SendErrJSON := common.SendErrJSON
	session := utsessions.GetSession(c)

	var todolist []*model.Todo
	var err error
	todolist, err = service.Todo.GetTodoList(math.MaxInt64, session.User.ID)
	if err != nil {
		SendErrJSON("获取待办事项错误", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": model.ErrorCode.SUCCESS,
		"msg":   "success",
		"data":  todolist,
	})
	return
}

//AddTodo 增加todo
func AddTodo(c *gin.Context) {
	SendErrJSON := common.SendErrJSON
	session := utsessions.GetSession(c)

	type Todo struct {
		Text string `json:"text"`
		Done bool   `json:"done"`
	}
	var todoBind Todo

	if err := c.ShouldBindWith(&todoBind, binding.JSON); err != nil {
		SendErrJSON("服务器错误，解析todo失败！", c)
		return
	}
	// log.Debugf("add todoBind [%#v]", todoBind)

	var todo model.Todo

	todo.Text = todoBind.Text
	todo.Done = todoBind.Done
	todo.UserID = session.User.ID

	err := service.Todo.AddTodo(todo)
	if err != nil {
		SendErrJSON("新增待办事项错误！", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": model.ErrorCode.SHOWSUCCESS,
		"msg":   "成功新增待办事项，要记得尽快完成哦～",
	})
	return

}

//DeleTodo 设置todolist
func DeleTodo(c *gin.Context) {
	SendErrJSON := common.SendErrJSON
	session := utsessions.GetSession(c)

	type Todo struct {
		ID   uint   `json:"id"`
		Text string `json:"text"`
		Done bool   `json:"done"`
	}
	var todoBind Todo

	if err := c.ShouldBindWith(&todoBind, binding.JSON); err != nil {
		SendErrJSON("参数无效", c)
		return
	}

	var todo model.Todo

	todo.ID = todoBind.ID
	todo.Text = todoBind.Text
	todo.Done = todoBind.Done
	todo.UserID = session.User.ID

	err := service.Todo.DeleTodo(todo)
	if err != nil {
		SendErrJSON("删除待办事项错误！", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": model.ErrorCode.SHOWSUCCESS,
		"msg":   "成功删除一条待办事项！",
	})
	return

}

//UpdateTodo 设置todolist
func UpdateTodo(c *gin.Context) {
	SendErrJSON := common.SendErrJSON
	session := utsessions.GetSession(c)

	type Todo struct {
		ID   uint   `json:"id"`
		Text string `json:"text"`
		Done bool   `json:"done"`
	}
	var todoBind Todo

	if err := c.ShouldBindWith(&todoBind, binding.JSON); err != nil {
		SendErrJSON("参数无效", c)
		return
	}

	var todo model.Todo

	todo.ID = todoBind.ID
	todo.Text = todoBind.Text
	todo.Done = todoBind.Done
	todo.UserID = session.User.ID

	err := service.Todo.UpdateTodo(todo)
	if err != nil {
		SendErrJSON("更新待办事项错误！", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": model.ErrorCode.SUCCESS,
		"msg":   "success",
	})
	return

}
