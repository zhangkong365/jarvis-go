package header

import (
	"github.com/thingworks/jarvis-go/core/process/domain/counter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JarvisSheetHeader struct {
	Id string `bson:"_id,omitempty"`
}

type JarvisHeaderEntryEntity struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
}

type JarvisHeaderEntries struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	SheetId    string
	HeaderId   string
	BatchId    string
	BatchIndex int32
	Start      int64
	End        int64
	total      int32
	Digest     string
	Counter    counter.DependencyCounter
	Rows       []JarvisHeaderEntryEntity
}

func (entries *JarvisHeaderEntries) CollectionName() string {
	return "header.entries"
}

func (entries *JarvisHeaderEntries) ObjectId() primitive.ObjectID {
	return entries.Id
}

func (entries *JarvisHeaderEntries) Init() {
}
