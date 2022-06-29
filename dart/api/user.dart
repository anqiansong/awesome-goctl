import 'api.dart';
import '../data/greet.dart';

/// User

/// --/user/login--
///
/// 请求: LoginReq
/// 返回: LoginResp
Future apiUserLogin( LoginReq request,
    {Function(LoginResp) ok,
    Function(String) fail,
    Function eventually}) async {
  await apiPost('/user/login',request,
  	 ok: (data) {
    if (ok != null) ok(LoginResp.fromJson(data));
  }, fail: fail, eventually: eventually);
}

/// --/user/info--
///
/// 请求: UserInfoReq
/// 返回: UserInfoResp
Future apiUserInfo( 
    {Function(UserInfoResp) ok,
    Function(String) fail,
    Function eventually}) async {
  await apiGet('/user/info',
  	 ok: (data) {
    if (ok != null) ok(UserInfoResp.fromJson(data));
  }, fail: fail, eventually: eventually);
}

