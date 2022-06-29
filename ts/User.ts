import axiosutil from "axiosutil"
import * as components from "./UserComponents"
export * from "./UserComponents"

/**
 * @description "用户登录"
 * @param req
 */
export function login(req: components.LoginReq) {
	return axiosutil.post<components.LoginResp>("/user/login", req)
}

/**
 * @description "获取用户信息"
 * @param params
 */
export function userinfo(params: components.UserInfoReqParams) {
	return axiosutil.get<components.UserInfoResp>("/user/info", params)
}
