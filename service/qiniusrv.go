package service

import (
	"sync"

	"gowork/log"
	"gowork/model"
)

// Qiniu service.
var Qiniu = &qiniuService{
	mutex: &sync.Mutex{},
}

type qiniuService struct {
	mutex *sync.Mutex
}

func (srv *qiniuService) GetQiniuMediaList(size int, userID uint) (ret []*model.QiniuStorage, err error) {
	if err := db.Where("`user_id` = ?", userID).Order("`id` ASC").Limit(size).Find(&ret).Error; err != nil {
		log.Errorf("GetQiniuMediaList failed: " + err.Error())
		return ret, err
	}

	return ret, nil
}

func (srv *qiniuService) DeleQiniuMedia(file model.QiniuStorage) error {

	if err := db.Where("`user_id` = ? AND `uid` = ?", file.UserID, file.UID).Delete(&file).Error; err != nil {
		return err
	}

	return nil
}

func (srv *qiniuService) AddQiniuMedia(file model.QiniuStorage) error {

	if err := db.Create(&file).Error; err != nil {
		return err
	}

	return nil
}
