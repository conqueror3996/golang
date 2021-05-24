package model

// OperatorRole OperatorRole's Entity
type OperatorRole struct {
	OperatorRoleID string `gorm:"column:operator_role_id"`
	CardMemberRole string `gorm:"column:card_member_role"`
	ClubRole       string `gorm:"column:club_role"`
	ClubOwnerRole  string `gorm:"column:club_owner_role"`
	MessageRole    string `gorm:"column:message_role"`
	OperatorRole   string `gorm:"column:operator_role"`
	CreateDate     string `gorm:"column:create_date"`
	CreatePgmID    string `gorm:"column:create_pgm_id"`
	UpdateDate     string `gorm:"column:update_date"`
	UpdatePgmID    string `gorm:"column:update_pgm_id"`
}

// OperatorRoleForPut OperatorRoleForPut's Entity
type OperatorRoleForPut struct {
	OperatorRoleID string `gorm:"column:operator_role_id"`
	CardMemberRole string `gorm:"column:card_member_role"`
	ClubRole       string `gorm:"column:club_role"`
	ClubOwnerRole  string `gorm:"column:club_owner_role"`
	MessageRole    string `gorm:"column:message_role"`
	OperatorRole   string `gorm:"column:operator_role"`
	CreateDate     string `gorm:"column:create_date"`
	CreatePgmID    string `gorm:"column:create_pgm_id"`
	UpdateDate     string `gorm:"column:update_date"`
	UpdatePgmID    string `gorm:"column:update_pgm_id"`
}

// OperatorRoleMeta OperatorRoleMeta's Entity
type OperatorRoleMeta struct {
}

// OperatorRoleCounts Out Entity
type OperatorRoleCounts struct {
	Count int `gorm:"column:count"`
}
