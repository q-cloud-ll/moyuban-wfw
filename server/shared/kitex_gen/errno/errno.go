// Code generated by thriftgo (0.3.2). DO NOT EDIT.

package errno

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type Err int64

const (
	Err_Success        Err = 0
	Err_NoRoute        Err = 1
	Err_NoMethod       Err = 2
	Err_BadRequest     Err = 10000
	Err_ParamsErr      Err = 10001
	Err_AuthorizeFail  Err = 10002
	Err_TooManyRequest Err = 10003
	Err_ServiceErr     Err = 20000
	Err_RPCUserSrvErr  Err = 30000
	Err_UserSrvErr     Err = 30001
)

func (p Err) String() string {
	switch p {
	case Err_Success:
		return "Success"
	case Err_NoRoute:
		return "NoRoute"
	case Err_NoMethod:
		return "NoMethod"
	case Err_BadRequest:
		return "BadRequest"
	case Err_ParamsErr:
		return "ParamsErr"
	case Err_AuthorizeFail:
		return "AuthorizeFail"
	case Err_TooManyRequest:
		return "TooManyRequest"
	case Err_ServiceErr:
		return "ServiceErr"
	case Err_RPCUserSrvErr:
		return "RPCUserSrvErr"
	case Err_UserSrvErr:
		return "UserSrvErr"
	}
	return "<UNSET>"
}

func ErrFromString(s string) (Err, error) {
	switch s {
	case "Success":
		return Err_Success, nil
	case "NoRoute":
		return Err_NoRoute, nil
	case "NoMethod":
		return Err_NoMethod, nil
	case "BadRequest":
		return Err_BadRequest, nil
	case "ParamsErr":
		return Err_ParamsErr, nil
	case "AuthorizeFail":
		return Err_AuthorizeFail, nil
	case "TooManyRequest":
		return Err_TooManyRequest, nil
	case "ServiceErr":
		return Err_ServiceErr, nil
	case "RPCUserSrvErr":
		return Err_RPCUserSrvErr, nil
	case "UserSrvErr":
		return Err_UserSrvErr, nil
	}
	return Err(0), fmt.Errorf("not a valid Err string")
}

func ErrPtr(v Err) *Err { return &v }
func (p *Err) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = Err(result.Int64)
	return
}

func (p *Err) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}
