package confession

import (
	"errors"
	"time"

	"fptugo/configs/db"
)

// Confession ...
type Confession struct {
	ID        int        `json:"id" gorm:"primary_key"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index"`

	Content  string `json:"content" gorm:"not null; type:text;"`
	Sender   string `json:"sender" gorm:"not null; type:varchar(250);"`
	PushID   string `json:"push_id" gorm:"type:varchar(250);"`
	Status   int    `json:"status" gorm:"not null; type:int(11);"`
	Approver int    `json:"approver" gorm:"type:int(11);"`
	Reason   string `json:"reason" gorm:"type:varchar(250);"`
	CfsID    int    `json:"cfs_id" gorm:"type:int(11);"`
}

// TableName set Confession's table name to be `confessions`
func (Confession) TableName() string {
	return "confessions"
}

// FetchAll ...
func (c *Confession) FetchAll(numLoad int) []Confession {
	db := db.GetDatabaseConnection()

	var confessions []Confession
	db.Order("id desc").Limit(numLoad).Find(&confessions)

	return confessions
}

// FetchByID ...
func (c *Confession) FetchByID() error {
	db := db.GetDatabaseConnection()

	if err := db.Where("id = ?", c.ID).Find(&c).Error; err != nil {
		return errors.New("Could not find the confession")
	}

	return nil
}

// Create ...
func (c *Confession) Create() error {
	db := db.GetDatabaseConnection()

	// Validate record
	if !db.NewRecord(c) { // => returns `true` as primary key is blank
		return errors.New("New records can not have primary key id")
	}

	if err := db.Create(&c).Error; err != nil {
		return errors.New("Could not create confession")
	}

	return nil
}

// Save ...
func (c *Confession) Save() error {
	db := db.GetDatabaseConnection()

	if db.NewRecord(c) {
		if err := db.Create(&c).Error; err != nil {
			return errors.New("Could not create confessions")
		}
	} else {
		if err := db.Save(&c).Error; err != nil {
			return errors.New("Could not update confessions")
		}
	}
	return nil
}

// FetchBySender ...
func (c *Confession) FetchBySender(sender string, numLoad int) []Confession {
	db := db.GetDatabaseConnection()

	var confessions []Confession
	db.Order("id desc").Limit(numLoad).Where("sender = ?", sender).Find(&confessions)

	return confessions
}

type overviewSpec struct {
	Total   int
	Pending int
	Reject  int
}

// FetchOverview ...
func (c *Confession) FetchOverview() (int, int, int) {
	db := db.GetDatabaseConnection()
	totalCount, pendingCount, rejectedCount := 0, 0, 0

	db.Model(&Confession{}).Count(&totalCount)
	db.Model(&Confession{}).Where("status = ?", 0).Count(&pendingCount)
	db.Model(&Confession{}).Where("status = ?", 2).Count(&rejectedCount)

	return totalCount, pendingCount, rejectedCount
}

// FetchApprovedConfession ...
func (c *Confession) FetchApprovedConfession(lastestID int, isAuthenticated bool) []Confession {
	db := db.GetDatabaseConnection()

	var isAuthenticatedQueryStr string
	if isAuthenticated {
		isAuthenticatedQueryStr = "OR status = 2"
	} else {
		isAuthenticatedQueryStr = ""
	}

	var confessions []Confession
	if lastestID == 0 {
		db.Where("status = 1" + isAuthenticatedQueryStr).Order("id desc").Limit(10).Find(&confessions)
	} else {
		db.Where("id < ? and status = 1"+isAuthenticatedQueryStr, lastestID).Order("id desc").Limit(10).Find(&confessions)
	}

	return confessions
}

// GetNextConfessionID ...
func (c *Confession) GetNextConfessionID() int {
	db := db.GetDatabaseConnection()
	db.Order("cfs_id desc").Take(&c)
	return c.CfsID + 1
}

func (c *Confession) setConfessionApproved(status int, approver int, cfsID int) {
	c.Status = status
	c.Approver = approver
	c.CfsID = cfsID
}

// ApproveConfession ...
func (c *Confession) ApproveConfession(approverID int) error {
	if err := c.FetchByID(); err != nil {
		return errors.New("Could not find the confession")
	}

	if c.Status != 0 {
		return errors.New("Status of confession must be pending to be approved")
	}

	confessions := new(Confession)

	c.setConfessionApproved(1, approverID, confessions.GetNextConfessionID())

	if err := c.Save(); err != nil {
		return errors.New("Unable to update approved confession`")
	}

	return nil
}

func (c *Confession) setConfessionRejected(status int, approver int, reason string) {
	c.Status = status
	c.Approver = approver
	c.Reason = reason
}

// RejectConfession ...
func (c *Confession) RejectConfession(approverID int, reason string) error {
	if err := c.FetchByID(); err != nil {
		return errors.New("Could not find the confession")
	}

	if c.Status != 0 {
		return errors.New("Status of confession must be pending to be rejected")
	}

	c.setConfessionRejected(2, approverID, reason)

	if err := c.Save(); err != nil {
		return errors.New("Unable to update approved confession`")
	}

	return nil
}

// SearchConfession ...
func (c *Confession) SearchConfession(keyword string) []Confession {
	db := db.GetDatabaseConnection()

	var confessions []Confession
	db.Order("id desc").Limit(50).Where("status = 1 AND content LIKE?", "%"+keyword+"%").Find(&confessions)

	return confessions
}

// SyncPushID ...
func (c *Confession) SyncPushID(sender string, pushID string) {
	db := db.GetDatabaseConnection()

	db.Model(&c).Where("sender = ?", sender).Update("push_id", pushID)
}
