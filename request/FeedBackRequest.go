// Package request
/**
 * @projectName:    Store
 * @package:        request
 * @className:      FeedBackRequest
 * @author:     张杰
 * @description:  TODO
 * @date:    2024/5/24 9:07
 * @version:    1.0
 */
package request

type AddFeedbackRequest struct {
	Name    string `json:"name" binding:"required"`
	Phone   string `json:"phone" binding:"required,len=11"`
	Email   string `json:"email" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type GetFeedbackRequest struct {
}
