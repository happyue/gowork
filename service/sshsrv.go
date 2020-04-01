package service

import (
	"sync"

	"gowork/log"
	"gowork/model"
)

// Ssh service.
var Ssh = &sshService{
	mutex: &sync.Mutex{},
}

type sshService struct {
	mutex *sync.Mutex
}

func (srv *sshService) GetMachineListASC(size, page, pageSize int, userID uint) (ret []*model.Machine, count int, err error) {
	offset := (page - 1) * pageSize
	if err := db.Model(&model.Machine{}).Where("`user_id` = ?", userID).Order("`id` ASC").Count(&count).Offset(offset).
		Limit(size).Find(&ret).Error; err != nil {
		log.Errorf("GetMachineList failed: " + err.Error())
		return ret, count, err
	}

	return ret, count, nil
}

func (srv *sshService) GetMachineListDESC(size, page, pageSize int, userID uint) (ret []*model.Machine, count int, err error) {
	offset := (page - 1) * pageSize
	if err := db.Model(&model.Machine{}).Where("`user_id` = ?", userID).Order("`id` DESC").Count(&count).Offset(offset).
		Limit(size).Find(&ret).Error; err != nil {
		log.Errorf("GetMachineList failed: " + err.Error())
		return ret, count, err
	}

	return ret, count, nil
}

func (srv *sshService) DeleMachine(machine model.Machine) error {

	if err := db.Where("`user_id` = ? AND `id` = ?", machine.UserID, machine.ID).Delete(&machine).Error; err != nil {
		return err
	}

	return nil
}

func (srv *sshService) FindMachineByID(uid, id uint) (*model.Machine, error) {
	ret := &model.Machine{}
	if err := db.Where("`user_id` = ? AND `id` = ?", uid, id).Find(&ret).Error; err != nil {
		log.Errorf("FindMachineByID failed: " + err.Error())
		return ret, err
	}

	return ret, nil
}

func (srv *sshService) AddMachine(machine model.Machine) error {

	if err := db.Create(&machine).Error; err != nil {
		return err
	}

	return nil
}

func (srv *sshService) AddSSHLog(sshLog model.SSHLog) error {

	if err := db.Create(&sshLog).Error; err != nil {
		return err
	}

	return nil
}

func (srv *sshService) UpdateMachine(machine model.Machine) error {
	// log.Debug(machine.Port)

	if err := db.Model(&machine).Where("`user_id` = ? AND `id` = ?", machine.UserID, machine.ID).
		Updates(map[string]interface{}{"name": machine.Name, "host": machine.Host, "port": machine.Port,
			"user": machine.User, "password": machine.Password, "type": machine.Type}).Error; err != nil {
		return err
	}

	return nil
}
