package common

import "time"

type DomainBase interface {
	CreatedBy() string
	CreationTime() time.Time
	ModifiedBy() string
	ModificationTime() string
}
