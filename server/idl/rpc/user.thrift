namespace go user

include "../base/common.thrift"
include "../base/user.thrift"

struct RegisterRequest {
    1: string username
    2: string password
}

struct RegisterResponse {
    1: common.BaseResponse base_resp
}

struct PhoneLoginRequest {
    1: string mobile
    2: string code
}

struct PhoneLoginResponse {
    1: common.BaseResponse base_resp
    2: string token
}

service UserService {
    RegisterResponse Register(1: RegisterRequest req)
    PhoneLoginResponse PhoneLogin(1:PhoneLoginRequest req)
}