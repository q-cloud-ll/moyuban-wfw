namespace go errno

enum Err {
    Success            = 0,
    NoRoute            = 1,
    NoMethod           = 2,
    BadRequest         = 10000,
    ParamsErr          = 10001,
    AuthorizeFail      = 10002,
    TooManyRequest     = 10003,
    ServiceErr         = 20000,
    RPCUserSrvErr      = 30000,
    UserSrvErr         = 30001,
}