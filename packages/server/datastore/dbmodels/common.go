package dbmodels

import uuid "github.com/satori/go.uuid"

type Activity struct {
	BaseModel   `gorm:"embedded"`
	UserUUID    uuid.UUID `json:"user_uuid" gorm:"type:uuid;not null"`
	Category    string    `json:"category"`
	Activity    string    `json:"activity"`
	ModelUUID   uuid.UUID `json:"model_uuid" gorm:"type:uuid;"`
	DatasetUUID uuid.UUID `json:"dataset_uuid" gorm:"type:uuid;"`

	User    User    `gorm:"foreignKey:UserUUID"`
	Model   Model   `gorm:"foreignKey:ModelUUID"`
	Dataset Dataset `gorm:"foreignKey:DatasetUUID"`
}

type Tag struct {
	ModelUUID        uuid.UUID `json:"model_uuid" gorm:"type:uuid;primaryKey"`
	DatasetUUID      uuid.UUID `json:"dataset_uuid" gorm:"type:uuid;primaryKey"`
	OrganizationUUID uuid.UUID `json:"organization_uuid" gorm:"type:uuid;not null;index:idx_org_tag,unique"`
	Tag              string    `json:"tag" gorm:"not null;index:idx_org_tag,unique"`

	Model   Model        `gorm:"foreignKey:ModelUUID"`
	Dataset Dataset      `gorm:"foreignKey:DatasetUUID"`
	Org     Organization `gorm:"foreignKey:OrganizationUUID"`
}

type Log struct {
	BaseModel          `gorm:"embedded"`
	Data               string    `json:"data"`
	ModelVersionUUID   uuid.UUID `json:"model_version_uuid" gorm:"type:uuid;"`
	DatasetVersionUUID uuid.UUID `json:"dataset_version_uuid" gorm:"type:uuid;"`

	ModelVersion   ModelVersion   `gorm:"foreignKey:ModelVersionUUID"`
	DatasetVersion DatasetVersion `gorm:"foreignKey:DatasetVersionUUID"`
}
