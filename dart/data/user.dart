// --greet--

class LoginResp{
	
	/// 
	final String token;
	
	/// 
	final String uid;
	
	LoginResp({ 
		this.token,
		this.uid,
	});
	factory LoginResp.fromJson(Map<String,dynamic> m) {
		return LoginResp(
			token: m['token'],
			uid: m['uid'],
		);
	}
	Map<String,dynamic> toJson() {
		return { 
			'token': token,
			'uid': uid,
		};
	}
}

class LoginReq{
	
	/// 
	final String username;
	
	/// 
	final String password;
	
	LoginReq({ 
		this.username,
		this.password,
	});
	factory LoginReq.fromJson(Map<String,dynamic> m) {
		return LoginReq(
			username: m['username'],
			password: m['password'],
		);
	}
	Map<String,dynamic> toJson() {
		return { 
			'username': username,
			'password': password,
		};
	}
}

class UserInfoReq{
	
	/// 
	final String uid;
	
	UserInfoReq({ 
		this.uid,
	});
	factory UserInfoReq.fromJson(Map<String,dynamic> m) {
		return UserInfoReq(
			uid: m['uid'],
		);
	}
	Map<String,dynamic> toJson() {
		return { 
			'uid': uid,
		};
	}
}

class UserInfoResp{
	
	/// 
	final String uid;
	
	/// 
	final String name;
	
	/// 
	final String email;
	
	UserInfoResp({ 
		this.uid,
		this.name,
		this.email,
	});
	factory UserInfoResp.fromJson(Map<String,dynamic> m) {
		return UserInfoResp(
			uid: m['uid'],
			name: m['name'],
			email: m['email'],
		);
	}
	Map<String,dynamic> toJson() {
		return { 
			'uid': uid,
			'name': name,
			'email': email,
		};
	}
}

