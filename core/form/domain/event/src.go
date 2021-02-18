package event

import (
	"github.com/thingworks/jarvis-go/core/form/domain/sheet"
	"github.com/thingworks/jarvis-go/core/form/domain/sheet/header"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SheetDomainEventType string

const (
	HeaderRemoved    SheetDomainEventType = "HeaderRemoved"
	SheetDeleted     SheetDomainEventType = "SheetDeleted"
	IndicatorDeleted SheetDomainEventType = "IndicatorDeleted"
)

type SheetDelete struct {
	Id            primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Type          SheetDomainEventType `json:"type,omitempty"`
	SheetId       string               `json:"sheetId,omitempty"`
	UserId        string               `json:"userId,omitempty"`
	DomainVersion int64                `json:"domainVersion,omitempty"`
	EventTime     int64                `json:"eventTime,omitempty"`
	Sheet         sheet.JarvisSheet    `json:"sheet,omitempty"`
}

func (sheetDel *SheetDelete) CollectionName() string {
	return "sheet_event"
}

func (sheetDel *SheetDelete) TypeAlias() string {
	return "sheet.event.sheetDeleted"
}

func (sheetDel *SheetDelete) ObjectId() primitive.ObjectID {
	return sheetDel.Id
}

func (sheetDel *SheetDelete) Init() {
	sheetDel.Type = SheetDeleted
}

type HeaderRemove struct {
	Id            primitive.ObjectID       `bson:"_id,omitempty" json:"id,omitempty"`
	SheetId       string                   `json:"sheetId,omitempty"`
	Type          SheetDomainEventType     `json:"type,omitempty"`
	UserId        string                   `json:"userId,omitempty"`
	DomainVersion int64                    `json:"domainVersion,omitempty"`
	EventTime     int64                    `json:"eventTime,omitempty"`
	Header        header.JarvisSheetHeader `json:"header,omitempty"`
	Class         string                   `bson:"_class" json:"typeAlias,omitempty"`
}

func (headerRm *HeaderRemove) CollectionName() string {
	return "sheet_event"
}

func (headerRm *HeaderRemove) TypeAlias() string {
	return "sheet.event.headerRemoved"
}

func (headerRm *HeaderRemove) ObjectId() primitive.ObjectID {
	return headerRm.Id
}

func (headerRm *HeaderRemove) Init() {
	headerRm.Type = HeaderRemoved
}
