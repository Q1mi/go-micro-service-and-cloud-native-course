package biz

// ReplyParam 商家回复评价的参数
type ReplyParam struct {
	ReviewID  int64
	StoreID   int64
	Content   string
	PicInfo   string
	VideoInfo string
}

// AuditParam 运营审核评价的参数
type AuditParam struct {
	ReviewID  int64
	OpUser    string
	OpReason  string
	OpRemarks string
	Status    int32
}

// AppealParam 商家申诉评价的参数
type AppealParam struct {
	ReviewID  int64
	StoreID   int64
	Reason    string
	Content   string
	PicInfo   string
	VideoInfo string
	OpUser    string
}

// AuditAppealParam O端审核商家申诉的参数
type AuditAppealParam struct {
	AppealID int64
	OpUser   string
	OpReason string
	Status   int32
}
