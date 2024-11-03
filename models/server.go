package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Server struct {
	gorm.Model
	Type        string `gorm:"type:varchar(100);not null" json:"type"`
	SubType     string `gorm:"type:varchar(100)" json:"subType"`
	Version     string `gorm:"type:varchar(50)" json:"version"`
	File        string `gorm:"type:text" json:"file"`
	DisplaySize string `gorm:"type:varchar(20)" json:"displaySize"`
	ByteSize    uint   `gorm:"not null" json:"byteSize"`
	MD5         string `gorm:"type:char(32);not null" json:"md5"`
	Built       int64  `gorm:"not null" json:"built"`
	Stability   string `gorm:"type:varchar(50);not null;default:'stable'" json:"stability"`
}

func (s *Server) CreateRecord(db *gorm.DB) {
	db.Create(s)
}

func (s *Server) FindRecord(db *gorm.DB, server *Server, findString string, findValue string) {
	db.First(server, fmt.Sprintf("%s = ?", findString), findValue)
}
