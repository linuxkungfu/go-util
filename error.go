package util

type ErrorCode int

const (
	Code_Ok            ErrorCode = 0
	Code_Http_Ok       ErrorCode = 200
	Code_Http_NotFound ErrorCode = 404
	Code_Unknown_Error ErrorCode = 999
	// web server error code
	Code_User_Name_Exist     ErrorCode = 1000
	Code_User_Name_Not_Exist ErrorCode = 1001
	Code_Open_File_Failed    ErrorCode = 1002
	Code_Json_Parse_Failed   ErrorCode = 1003
	Code_Load_Config_Failed  ErrorCode = 1004

	Code_Miss_Parameters   ErrorCode = 1005
	Code_Path_Config_Error ErrorCode = 1006

	Code_Trace_Event_Is_Empty ErrorCode = 2001

	// server inner error
	Code_Struct_Member_Error          ErrorCode = 5001
	Code_Peer_RTP_NOT_Replicate       ErrorCode = 5002
	Code_Create_Transport_Error       ErrorCode = 5003
	Code_Start_Transport_Error        ErrorCode = 5004
	Code_Signalling_Replicate_Error   ErrorCode = 5005
	Code_Http_Error                   ErrorCode = 5006
	Code_Http_Res_Error               ErrorCode = 5007
	Code_Http_Res_Data_Error          ErrorCode = 5008
	Code_Http_Res_Code_Error          ErrorCode = 5009
	Code_Build_Rtp_Error              ErrorCode = 5010
	Code_Signalling_UnReplicate_Error ErrorCode = 5011

	// websocket error
	Code_WebSocket_Unauth              ErrorCode = 6001
	Code_WebSocket_MissParams          ErrorCode = 6002
	Code_WebSocket_Auth_Failed         ErrorCode = 6003
	Code_WebSocket_Illegal_Token       ErrorCode = 6004
	Code_WebSocket_Invalid_Auth_Params ErrorCode = 6005
)

var errorMsgMap = map[ErrorCode]string{
	Code_Ok:                  "ok",
	Code_Http_Ok:             "http ok",
	Code_Http_NotFound:       "not found",
	Code_Unknown_Error:       "unknown error",
	Code_User_Name_Exist:     "user name is exist",
	Code_User_Name_Not_Exist: "user name is not exist",
	Code_Open_File_Failed:    "open file failed",
	Code_Json_Parse_Failed:   "json parse failed",
	Code_Load_Config_Failed:  "config load failed",
	Code_Miss_Parameters:     "miss required patameters",
	Code_Path_Config_Error:   "router path config error",

	Code_Trace_Event_Is_Empty: "trace event requires at least one",

	Code_Struct_Member_Error:        "struct mumber error",
	Code_Peer_RTP_NOT_Replicate:     "peer rtp not replicate",
	Code_Create_Transport_Error:     "create plain transport failed",
	Code_Start_Transport_Error:      "start transport failed",
	Code_Signalling_Replicate_Error: "replicate signalling failed",
	Code_Http_Error:                 "http request error",
	Code_Http_Res_Error:             "http response error",
	Code_Http_Res_Data_Error:        "http response data error",
	Code_Http_Res_Code_Error:        "http response buz code error",
	Code_Build_Rtp_Error:            "build rtp error",

	Code_WebSocket_Unauth:              "unauthorized error",
	Code_WebSocket_MissParams:          "miss require parameters",
	Code_WebSocket_Auth_Failed:         "authentication failed",
	Code_WebSocket_Illegal_Token:       "illegal token",
	Code_WebSocket_Invalid_Auth_Params: "invalid auth parameters",
}

func ErrorCodeToMsg(ec ErrorCode) string {
	value, ok := errorMsgMap[ec]
	if ok {
		return value
	} else {
		return errorMsgMap[Code_Unknown_Error]
	}
}
