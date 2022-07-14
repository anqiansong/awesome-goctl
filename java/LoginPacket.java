package com.xhb.logic.http.packet.User;

import com.xhb.core.packet.HttpPacket;
import com.xhb.core.network.HttpRequestClient;
import com.xhb.logic.http.packet.User.model.*;

/*
    "用户登录"	
*/
public class LoginPacket extends HttpPacket<LoginResp> {
	
	public LoginPacket(LoginReq request) {
		super(request);
		this.request = request;
    }

	@Override
    public HttpRequestClient.Method requestMethod() {
        return HttpRequestClient.Method.POST;
    }

	@Override
    public String requestUri() {
        return "user/login";
    }
}
