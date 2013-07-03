package service

import (
	"com_asp_proto"
	"fmt"
	"log"
	"net/http"
)

const kStatusSubmited = 1
const kStatusConfirmed = 2
const kStatusCanceld = 3

var kStatusDesc []string = make([]string, 3)

func handleSubmitOrder(w http.ResponseWriter, req *com_asp_proto.SubmitCarRentOrderRequest) {
	var res com_asp_proto.SubmitCarRentOrderResponse
	res.OrderId = req.Order.OrderId
	res.OrderStatus = new(com_asp_proto.OrderStatus)
	*(res.OrderStatus) = com_asp_proto.OrderStatus(kStatusSubmited)
	sql := `insert into t_orders(orderId, rentType, carType, startDate, contactName,
	        contactPhone, contactEmail, pickUpAt, destination, flightNumber,
	        planRentHours, invoiceTitle, status, status_desc,need_english_driver)
	         values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?, ?, ?, ?)`
	stmt, err := getOrderDBConn().Prepare(sql)
	if nil != err {
		log.Printf("Failed to prepare %s", sql)
	} else {
		status := kStatusSubmited
		statusstr := "Submited"
		need_driver := 1
		if !req.GetOrder().GetNeedEnglishDriver() {
			need_driver = 0
		}
		stmt.Bind(req.Order.OrderId, req.Order.RentTypeId,
			req.Order.CarTypeId, req.Order.StartDate, req.Order.ContactName,
			req.Order.ContactPhone, req.Order.ContactEmail, req.Order.PickUpAt,
			req.Order.Destination, req.Order.FlightNumber, req.Order.PlanRentHours, req.Order.InvoiceTitle,
			&status, &statusstr, &need_driver)
		_, err = stmt.Run()
		if nil != err {
			log.Printf("Failed to exec sql: %s for reason:%v", sql, err)
			*(res.OrderStatus) = com_asp_proto.OrderStatus(kStatusCanceld)
		} else {
			subject := fmt.Sprintf("New Order:%s submitted!", req.Order.GetOrderId())
			content := fmt.Sprintf("Order Detail Reference\n")
			content = fmt.Sprintf("%s orderId: %s\n", content, req.Order.GetOrderId())
			content = fmt.Sprintf("%s rentType: %d\n", content, req.Order.GetRentTypeId())
			content = fmt.Sprintf("%s carType: %d\n", content, req.Order.GetRentTypeId())
			content = fmt.Sprintf("%s startDate: %s\n", content, req.Order.GetStartDate())
			content = fmt.Sprintf("%s contactName: %s\n", content, req.Order.GetContactName())
			content = fmt.Sprintf("%s contactPhone: %s\n", content, req.Order.GetContactPhone())
			content = fmt.Sprintf("%s contactEmail: %s\n", content, req.Order.GetContactEmail())
			content = fmt.Sprintf("%s pickUpAt: %s\n", content, req.Order.GetPickUpAt())
			content = fmt.Sprintf("%s destination: %s\n", content, req.Order.GetDestination())
			content = fmt.Sprintf("%s flightNumber: %s\n", content, req.Order.GetFlightNumber())
			content = fmt.Sprintf("%s planRentHours: %s\n", content, req.Order.GetPlanRentHours())
			content = fmt.Sprintf("%s invoiceTitle: %s\n", content, req.Order.GetInvoiceTitle())

			need_english_driver := "false"
			if req.Order.GetNeedEnglishDriver() {
				need_english_driver = "true"
			}
			content = fmt.Sprintf("%s need_english_driver: %s\n", content, need_english_driver)
			content = fmt.Sprintf("%s=================================================\n", content)
			content = fmt.Sprintf("%s Please check database for more information.\n", content)
			sendOrderMail(subject, content)
		}
	}
	writeProtoEvent(com_asp_proto.MessageType_SUBMIT_CAR_RENT_ORDER_RESPONSE, &res, w)
}

func handleCancelOrder(w http.ResponseWriter, req *com_asp_proto.CancelCarRentOrderRequest) {
	var res com_asp_proto.CancelCarRentOrderResponse
	sql := "update t_orders set status = ?, status_desc = ? where orderId=?"
	stmt, err := getOrderDBConn().Prepare(sql)
	if nil != err {
		log.Printf("Failed to prepare %s", sql)
	} else {
		status := kStatusCanceld
		statusstr := "Canceled"
		stmt.Bind(&status, &statusstr, req.OrderId)
		_, err = stmt.Run()
		if nil != err {
			log.Printf("Failed to exec sql: %s for reason:%v", sql, err)
		} else {
			rows, _, err := getOrderDBConn().Query("select rentType, carType, startDate, contactName, contactPhone, contactEmail, pickUpAt, destination, flightNumber, planRentHours, invoiceTitle, need_english_driver from t_orders where orderId='%s'", req.GetOrderId())
			if nil == err && len(rows) == 1 {
				subject := fmt.Sprintf("Order:%s has been canceld!", req.GetOrderId())
				content := fmt.Sprintf("Order Detail Reference\n")
				content = fmt.Sprintf("%s orderId: %s\n", content, req.GetOrderId())
				content = fmt.Sprintf("%s rentType: %s\n", content, rows[0].Str(0))
				content = fmt.Sprintf("%s carType: %s\n", content, rows[0].Str(1))
				content = fmt.Sprintf("%s startDate: %s\n", content, rows[0].Str(2))
				content = fmt.Sprintf("%s contactName: %s\n", content, rows[0].Str(3))
				content = fmt.Sprintf("%s contactPhone: %s\n", content, rows[0].Str(4))
				content = fmt.Sprintf("%s contactEmail: %s\n", content, rows[0].Str(5))
				content = fmt.Sprintf("%s pickUpAt: %s\n", content, rows[0].Str(6))
				content = fmt.Sprintf("%s destination: %s\n", content, rows[0].Str(7))
				content = fmt.Sprintf("%s flightNumber: %s\n", content, rows[0].Str(8))
				content = fmt.Sprintf("%s planRentHours: %s\n", content, rows[0].Str(9))
				content = fmt.Sprintf("%s invoiceTitle: %s\n", content, rows[0].Str(10))
				content = fmt.Sprintf("%s need_english_driver: %s\n", content, rows[0].Str(11))
				content = fmt.Sprintf("%s=================================================\n", content)
				content = fmt.Sprintf("%s Please check database for more information.\n", content)
				sendOrderMail(subject, content)
			}
		}
	}
	res.OrderId = req.OrderId
	res.OrderStatus = new(com_asp_proto.OrderStatus)
	*(res.OrderStatus) = com_asp_proto.OrderStatus(kStatusCanceld)
	writeProtoEvent(com_asp_proto.MessageType_CANCEL_CAR_RENT_ORDER_RESPONSE, &res, w)
}

func handleQueryOrder(w http.ResponseWriter, req *com_asp_proto.QueryCarRentOrderStatusRequest) {
	var res com_asp_proto.QueryCarRentOrderStatusResponse
	rows, result, err := getOrderDBConn().Query("select status from t_orders where orderId='%s'", req.GetOrderId())
	res.OrderId = req.OrderId
	res.OrderStatus = new(com_asp_proto.OrderStatus)
	*(res.OrderStatus) = com_asp_proto.OrderStatus(kStatusCanceld)
	if nil != err {
		log.Printf("Failed to query order:%s for reson:%v", req.GetOrderId(), err)
	} else {
		log.Printf("Query result:%d", len(rows))
		if len(rows) == 1 {
			s := result.Map("status")
			*(res.OrderStatus) = com_asp_proto.OrderStatus(rows[0].Int(s))
		}
	}
	writeProtoEvent(com_asp_proto.MessageType_QUERY_CAR_RENT_ORDER_STATUS_RESPONSE, &res, w)
}
