package model
import "time"

// SSHLog ssh log struct
type SSHLog struct {
	Model
	MachineID uint      `gorm:"index" json:"machine_id" form:"machine_id"`
	SSHUser   string    `json:"ssh_user" comment:"ssh账号"`
	ClientIP  string    `json:"client_ip" form:"client_ip"`
	StartedAt time.Time `json:"started_at" form:"started_at"`
	Status    uint      `json:"status" comment:"0-未标记 2-正常 4-警告 8-危险 16-致命"`
	Remark    string    `json:"remark"`
	Log       string    `gorm:"type:text" json:"log"`
	UserID uint   `json:"userID"`
}
