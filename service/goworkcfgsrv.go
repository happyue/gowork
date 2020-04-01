package service

import (
	"sync"

	"gowork/log"
	"gowork/model"
)

// Cfg service.
var Cfg = &cfgService{
	mutex: &sync.Mutex{},
}

type cfgService struct {
	mutex *sync.Mutex
}

func (srv *cfgService) GetGwcfg() (*model.Gwcfg, error) {
	ret := &model.Gwcfg{}

	if err := db.Where("`id` = ?", 1).First(&ret).Error; err != nil {
		log.Errorf("GetRegStatus: " + err.Error())
		return ret, err
	}

	return ret, nil
}

func (srv *cfgService) SetRegStatus(cfg model.Gwcfg) error {

	if err := db.Model(&cfg).Where("`id` = ?", cfg.ID).Updates(map[string]interface{}{"reg_status": cfg.RegStatus}).Error; err != nil {
		return err
	}

	return nil
}
