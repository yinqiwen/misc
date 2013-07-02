package service

import (
	"com_asp_proto"
	//"github.com/ziutek/mymysql/mysql"
	"net/http"
)

func handleGetServerList(w http.ResponseWriter, req *com_asp_proto.GetServerListRequest) {
	var res com_asp_proto.GetServerListResponse
	res.Items = make([]*com_asp_proto.ServerItem, 1)
	res.Items[0] = new(com_asp_proto.ServerItem)
	res.Items[0].Host = new(string)
	res.Items[0].Port = new(int32)
	res.Items[0].IspType = new(int32)
	res.Items[0].IspDesc = new(string)
	*(res.Items[0].Host) = "42.96.166.60"
	*(res.Items[0].Port) = 8080
	*(res.Items[0].IspType) = 1
	*(res.Items[0].Host) = "dual"
	writeProtoEvent(com_asp_proto.MessageType_GET_SERVER_LIST_RESPONSE, &res, w)
}

func handleSubmitAdvice(w http.ResponseWriter, req *com_asp_proto.SubmitAdviseRequest) {
	var res com_asp_proto.NotifyMessage
	if req.Email != nil {
		stmt, err := getUserDBConn().Prepare("update t_app_user set email=? where uid=?")
		if nil == err {
			stmt.Run(req.Email, req.Uid)
		}
	}

	stmt, err := getUserDBConn().Prepare("replace t_app_user_advice(uid,email,advice) values(?, ?, ?)")
	if nil == err {
		stmt.Run(req.Uid, req.Email, req.Advise)
	}
	res.Success = new(bool)
	*res.Success = true
	writeProtoEvent(com_asp_proto.MessageType_NOTIFY_MESSAGE, &res, w)
}
