package models

type Model struct {
	BaseModel
	Name      string `json:"name" gorm:"not null;unique:idx_org_model_name"`
	Wiki      string `json:"wiki"`
	OrgID     uint   `json:"org_id" gorm:"not null;unique:idx_org_model_name"`
	CreatedBy uint   `json:"created_by" gorm:"not null"`
	UpdatedBy uint   `json:"updated_by"`
	IsPublic  bool   `json:"is_public" default:"false"`

	Org           Organization `gorm:"foreignKey:OrgID"`
	CreatedByUser User         `gorm:"foreignKey:CreatedBy"`
	UpdatedByUser User         `gorm:"foreignKey:UpdatedBy"`

	Branches []ModelBranch `gorm:"foreignKey:ModelID"`
	Users    []User        `gorm:"many2many:model_users;"`
}

type ModelUser struct {
	ModelID uint `json:"model_id" gorm:"primaryKey"`
	UserID  uint `json:"user_id" gorm:"primaryKey"`
	Role    string
}

type ModelBranch struct {
	BaseModel
	Name      string `json:"name" gorm:"not null;unique:idx_model_branch"`
	ModelID   uint   `json:"model_id" gorm:"not null;unique:idx_model_branch"`
	IsDefault bool   `json:"is_default" default:"false"`

	Model Model `gorm:"foreignKey:ModelID"`

	Versions []ModelVersion `gorm:"foreignKey:BranchID"`
}

type ModelVersion struct {
	BaseModel
	Version  string `json:"version" gorm:"not null;unique:idx_branch_version"`
	BranchID uint   `json:"branch_id" gorm:"not null;unique:idx_branch_version"`
	Hash     string `json:"hash" gorm:"not null;unique"`
	PathID   uint   `json:"path_id"`

	Branch ModelBranch `gorm:"foreignKey:BranchID"`
	Path   Path        `gorm:"foreignKey:PathID"`
}

type ModelReview struct {
	BaseModel
	FromBranchID uint   `json:"from_branch_id" gorm:"not null"`
	ToBranchID   uint   `json:"to_branch_id" gorm:"not null"`
	Title        string `json:"title" gorm:"not null"`
	Description  string `json:"description"`
	CreatedBy    uint   `json:"created_by" gorm:"not null"`
	AssignedTo   uint   `json:"assigned_to"`
	IsComplete   bool   `json:"is_complete" default:"false"`
	IsAccepted   bool   `json:"is_accepted" default:"false"`

	FromBranch     ModelBranch `gorm:"foreignKey:FromBranchID"`
	ToBranch       ModelBranch `gorm:"foreignKey:ToBranchID"`
	CreatedByUser  User        `gorm:"foreignKey:CreatedBy"`
	AssignedToUser User        `gorm:"foreignKey:AssignedTo"`
}

// Response models

type ModelNameResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ModelResponse struct {
	ID        uint                 `json:"id"`
	Name      string               `json:"name"`
	Wiki      string               `json:"wiki"`
	Org       OrganizationResponse `json:"org"`
	CreatedBy UserHandleResponse   `json:"created_by"`
	UpdatedBy UserHandleResponse   `json:"updated_by"`
	IsPublic  bool                 `json:"is_public"`
}

type ModelUserResponse struct {
	User UserHandleResponse `json:"user"`
	Role string             `json:"role"`
}

type ModelBranchNameResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ModelBranchResponse struct {
	ID        uint              `json:"id"`
	Name      string            `json:"name"`
	Model     ModelNameResponse `json:"model"`
	IsDefault bool              `json:"is_default"`
}

type ModelVersionNameResponse struct {
	ID      uint   `json:"id"`
	Version string `json:"version"`
}

type ModelVersionResponse struct {
	ID      uint                    `json:"id"`
	Version string                  `json:"version"`
	Branch  ModelBranchNameResponse `json:"branch"`
	Hash    string                  `json:"hash"`
	Path    PathResponse            `json:"path"`
}

type ModelReviewResponse struct {
	ID          uint                `json:"id"`
	FromBranch  ModelBranchResponse `json:"from_branch"`
	ToBranch    ModelBranchResponse `json:"to_branch"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	CreatedBy   UserHandleResponse  `json:"created_by"`
	AssignedTo  UserHandleResponse  `json:"assigned_to"`
	IsComplete  bool                `json:"is_complete"`
	IsAccepted  bool                `json:"is_accepted"`
}
