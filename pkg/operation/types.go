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

// 文章发表和草稿参数
type articleArgs struct {
	ID       *int   `json:"id"`
	Cover    string `json:"cover"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	IsDraft  bool   `json:"isDraft"`
	Category string `json:"category"`
	Tags     string `json:"tags"`
	CID      int    `json:"cID"`
}
