// Code generated by goctl. DO NOT EDIT.
package types

type UserRequest struct {
	Id string `path:"id"`
}

type UserResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserInfo struct {
	Id   string `json:"id,optional"`
	Name string `json:"name"`
}
