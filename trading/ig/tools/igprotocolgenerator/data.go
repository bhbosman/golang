package main

var data = []WebRequestData{
	//
	// Sesssion
	//
	WebRequestData{
		Version:    2,
		Method:     "POST",
		StructName: "Authentication",
		Command:    "session",
		RequestParams: FunctionDataRequest{
			HasInputParams: true},
		ResponseParams: FunctionDataResponse{
			CloseBody:           true,
			SuccessCode:         200,
			CreateSessionParams: true}},
	WebRequestData{
		Version:    1,
		Method:     "DELETE",
		StructName: "Logout",
		Command:    "session",
		RequestParams: FunctionDataRequest{
			HasInputParams: false},
		ResponseParams: FunctionDataResponse{
			CloseBody:           true,
			SuccessCode:         204,
			CreateSessionParams: false}},
	WebRequestData{
		Version:    1,
		Method:     "PUT",
		StructName: "AccountSwitch",
		Command:    "session",
		RequestParams: FunctionDataRequest{
			HasInputParams: true},
		ResponseParams: FunctionDataResponse{
			CloseBody:           true,
			SuccessCode:         200,
			CreateSessionParams: false}},
	//
	// Market
	//
	
	WebRequestData{
		HelpURL:    []string{
			"http://labs.ig.com/rest-trading-api-reference/service-detail?id=270",
			"http://labs.ig.com/rest-trading-api-reference/service-detail?id=267"},
		Version:    1,
		Method:     "GET",
		StructName: "MarketNavigation",
		Command:    "marketnavigation",
		RequestParams: FunctionDataRequest{
			HasURLPostFix: true,
			HasInputParams: false},
		ResponseParams: FunctionDataResponse{
			CloseBody:           true,
			SuccessCode:         200,
			CreateSessionParams: false}},
}
