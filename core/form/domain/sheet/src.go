package sheet

import (
	"github.com/thingworks/jarvis/common"
	"github.com/thingworks/jarvis/common/domain"
	"github.com/thingworks/jarvis/core/form/domain/sheet/header"
	"github.com/thingworks/jarvis/label"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HeaderCompiledInfo struct {
	ResourceId   string              `json:"resourceId,omitempty"`
	ResourceType domain.ResourceType `json:"resourceType:omitempty"`
	Cacheable    bool                `json:"cacheable"`
}

type JarvisSheet struct {
	Id                 primitive.ObjectID         `bson:"_id,omitempty" json:"id"`
	SheetId            string                     `json:"sheetId"`
	Name               string                     `json:"name,omitempty"`
	IndexFormula       string                     `json:"indexFormula,omitempty"`
	LatestCacheVersion int64                      `json:"latestCacheVersion,omitempty"`
	Code               string                     `json:"code"`
	DomainVersion      int64                      `json:"domainVersion"`
	NamespaceId        string                     `json:"namespaceId"`
	TenantId           string                     `json:"tenantId"`
	Label              label.Label                `json:"label"`
	Folder             common.FolderInfo          `json:"folder"`
	IndexCompiled      HeaderCompiledInfo         `json:"indexCompiled"`
	Headers            []header.JarvisSheetHeader `json:"headers"`
	HeadersV2          []header.JarvisSheetHeader `json:"headersV2"`
}

func (sheet *JarvisSheet) CollectionName() string {
	return "sheet"
}

func (sheet *JarvisSheet) TypeAlias() string {
	return "jarvis.sheet"
}

func (sheet *JarvisSheet) ObjectId() primitive.ObjectID {
	return sheet.Id
}
