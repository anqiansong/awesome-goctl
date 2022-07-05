package "type desc here"

import com.google.gson.Gson

object typeTitleHereApi{
	
	data class LoginResp(
		val token: String,
		val uid: String
	)
	data class LoginReq(
		val username: String,
		val password: String
	)
	data class UserInfoReq(
		val uid: String
	)
	data class UserInfoResp(
		val uid: String,
		val name: String,
		val email: String
	)
	
	suspend fun postUserLogin(
		req:LoginReq,
		onOk: ((LoginResp) -> Unit)? = null,
        onFail: ((String) -> Unit)? = null,
        eventually: (() -> Unit)? = null
    ){
        apiRequest("POST","/user/login",body=req, onOk = { 
            onOk?.invoke(Gson().fromJson(it,LoginResp::class.java))
        }, onFail = onFail, eventually =eventually)
    }
	suspend fun getUserInfo(
		req:UserInfoReq,
		onOk: ((UserInfoResp) -> Unit)? = null,
        onFail: ((String) -> Unit)? = null,
        eventually: (() -> Unit)? = null
    ){
        apiRequest("GET","/user/info",body=req, onOk = { 
            onOk?.invoke(Gson().fromJson(it,UserInfoResp::class.java))
        }, onFail = onFail, eventually =eventually)
    }
	
}
