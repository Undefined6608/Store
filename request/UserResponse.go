// Package request
/**
 * @projectName:    Store
 * @package:        request
 * @className:      UserResponse
 * @author:     张杰
 * @description:  TODO
 * @date:    2024/4/7 13:40
 * @version:    1.0
 */
package request

// UserResponse /** 返回用户信息
type UserResponse struct {
	UserName  string `json:"userName"`  // 用户名
	Email     string `json:"email"`     // 邮箱 不允许为空 唯一
	Phone     string `json:"phone"`     // 电话号码 不允许为空 唯一
	Gender    bool   `json:"gender"`    // 性别
	LimitType uint8  `json:"limitType"` // 用户类型
	Avatar    string `json:"avatar"`    // 头像 默认值
	Integral  uint32 `json:"integral"`  //用户积分
	Balance   uint32 `json:"balance"`   // 余额
	LikeNum   uint32 `json:"likeNum"`   // 赞
	DontLike  uint32 `json:"dontLike"`  // 踩
	UID       string `json:"uid"`       // 用户身份
}
