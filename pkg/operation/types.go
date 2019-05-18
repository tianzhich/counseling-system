package operation

// 流程处理参数
type processArgs struct {
	StartTime     *string `json:"startTime"`
	Location      *string `json:"location"`
	CancelReason1 *string `json:"cancelReason1"`
	CancelReason2 *string `json:"cancelReason2"`
	RatingScore   *int    `json:"ratingScore"`
	RatingText    *string `json:"ratingText"`
	Letter        *string `json:"letter"`
}

// 新增问答Form
type askForm struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	IsAnony bool   `json:"isAnony"`
	Tags    string `json:"tags"`
}
