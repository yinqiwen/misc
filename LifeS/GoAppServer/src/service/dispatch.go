package service

import (
	"bytes"
	"code.google.com/p/goprotobuf/proto"
	"com_asp_proto"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

//	"util"
)

const MAGIC_HEADER uint16 = 0xFADA

func writeProtoEvent(proto_type com_asp_proto.MessageType, pb proto.Message, w http.ResponseWriter) {
	w.WriteHeader(200)
	var body bytes.Buffer
	tmp := MAGIC_HEADER
	binary.Write(&body, binary.BigEndian, &tmp)
	proto_type_int := int32(proto_type)
	binary.Write(&body, binary.BigEndian, &proto_type_int)
	var reserved uint32
	binary.Write(&body, binary.BigEndian, &reserved)
	buf, _ := proto.Marshal(pb)
	body.Write(buf)
	w.Write(body.Bytes())
}

func Dispatch(w http.ResponseWriter, r *http.Request) {
	if !strings.EqualFold(r.Method, "POST") {
		w.WriteHeader(400)
		w.Write([]byte("Only post method supported"))
		return
	}

	if r.ContentLength >= 1024*1024 || r.ContentLength <= 10 {
		w.WriteHeader(400)
		reason := fmt.Sprintf("[ERROR]Invalid request content-length:%d", r.ContentLength)
		w.Write([]byte(reason))
		log.Printf("%s", reason)
		return
	}

	var magic uint16
	err := binary.Read(r.Body, binary.BigEndian, &magic)
	if nil != err || magic != MAGIC_HEADER {
		w.WriteHeader(400)
		log.Printf("[ERROR]Read magic failed:%v-%x", err, magic)
		return
	}
	var proto_type int32
	err = binary.Read(r.Body, binary.BigEndian, &proto_type)
	if nil != err {
		w.WriteHeader(400)
		log.Printf("[ERROR]Read proto type failed:%v-%d", err, proto_type)
		return
	}
	var reserved uint32
	err = binary.Read(r.Body, binary.BigEndian, &reserved)
	if nil != err {
		w.WriteHeader(400)
		log.Printf("[ERROR]Read reserved failed:%v-%d", err, proto_type)
		return
	}
	var body []byte
	body, err = ioutil.ReadAll(r.Body)
	if nil != err {
		w.WriteHeader(400)
		log.Printf("[ERROR]Read rest body failed:%v", err)
		return
	}

	switch com_asp_proto.MessageType(proto_type) {
	case com_asp_proto.MessageType_GET_APPSTORE_URL_REQUEST:
		var req com_asp_proto.GetAppstoreURLRequest
		err = proto.Unmarshal(body, &req)
	case com_asp_proto.MessageType_GET_SERVER_LIST_REQUEST:
		var req com_asp_proto.GetServerListRequest
		err = proto.Unmarshal(body, &req)
		if nil == err {
			handleGetServerList(w, &req)
		}
	case com_asp_proto.MessageType_CANCEL_CAR_RENT_ORDER_REQUEST:
		var req com_asp_proto.CancelCarRentOrderRequest
		err = proto.Unmarshal(body, &req)
		if nil == err {
			handleCancelOrder(w, &req)
		}
	case com_asp_proto.MessageType_SUBMIT_CAR_RENT_ORDER_REQUEST:
		var req com_asp_proto.SubmitCarRentOrderRequest
		err = proto.Unmarshal(body, &req)
		if nil == err {
			handleSubmitOrder(w, &req)
		}
	case com_asp_proto.MessageType_QUERY_CAR_RENT_ORDER_STATUS_REQUEST:
		var req com_asp_proto.QueryCarRentOrderStatusRequest
		err = proto.Unmarshal(body, &req)
		if nil == err {
			handleQueryOrder(w, &req)
		}
	case com_asp_proto.MessageType_SUBMIT_ADVISE_REQUEST:
		var req com_asp_proto.SubmitAdviseRequest
		err = proto.Unmarshal(body, &req)
		if nil == err {
			handleSubmitAdvice(w, &req)
		}
	}

	if nil != err {
		w.WriteHeader(400)
		log.Printf("[ERROR]Failed to parse proto request:%v", err)
		return
	} else {
		log.Printf("Parse proto request:%d success", proto_type)
	}
}
