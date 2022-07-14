package com.xhb.logic.http.packet.User;

import com.xhb.core.packet.HttpPacket;
import com.xhb.core.network.HttpRequestClient;
import com.xhb.logic.http.packet.User.model.*;

/*
    "获取用户信息"	
*/
public class UserinfoPacket extends HttpPacket<UserInfoResp> {
	
	public UserinfoPacket() {
		super(EmptyRequest.instance);
		
    }

	@Override
    public HttpRequestClient.Method requestMethod() {
        return HttpRequestClient.Method.GET;
    }

	@Override
    public String requestUri() {
        return "user/info";
    }
}
