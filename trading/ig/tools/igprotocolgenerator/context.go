package main

import (
	"fmt"
)

// FunctionDataRequest ...
type FunctionDataRequest struct {
	HasInputParams bool
	HasURLPostFix bool
}

// FunctionDataResponse ...
type FunctionDataResponse struct {
	CreateSessionParams bool
	SuccessCode         int
	CloseBody           bool
}

// WebRequestData ...
type WebRequestData struct {
	StructName     string
	Command        string
	Method         string
	Version        int
	HelpURL        []string
	RequestParams  FunctionDataRequest
	ResponseParams FunctionDataResponse
}

// CreateInputParamCode ...
func (context WebRequestData) CreateInputParamCode() string {
	if context.RequestParams.HasInputParams {
		return `arrayOfBytes, _ := json.Marshal(inputData)`
	}
	return ""
}

// CreateInputParam ...
func (context WebRequestData) CreateInputParam() string {
	if context.RequestParams.HasInputParams || context.RequestParams.HasURLPostFix {
		return fmt.Sprintf("\ninputData %sRequest, ", context.StructName)
	}
	return ""
}

// CreateSessionParam ...
func (context WebRequestData) CreateSessionParam() string {
	if context.ResponseParams.CreateSessionParams {
		return `SessionKeys := SessionKeys{
		SecurityToken: response.Header.Get(XSecurityTokenConst), 
		CST: response.Header.Get(CstConst)}`
	}
	return ""
}

// AssignSessionParam ...
func (context WebRequestData) AssignSessionParam() string {
	if context.ResponseParams.CreateSessionParams {
		return "SessionKeys: SessionKeys,"
	}
	return ""
}

// AssignRequestBody ...
func (context WebRequestData) AssignRequestBody() string {
	if context.RequestParams.HasInputParams {
		return "Body:          ioutil.NopCloser(bytes.NewBuffer(arrayOfBytes)),"
	}
	return ""
}

// AssignRequestLength ...
func (context WebRequestData) AssignRequestLength() string {
	if context.RequestParams.HasInputParams {
		return "ContentLength: int64(len(arrayOfBytes)),"
	}
	return ""
}

// UseInputParam ...
func (context WebRequestData) UseInputParam() string {
	if context.RequestParams.HasInputParams {
		return "inputData, "
	}
	return ""
}

// UseInputParamBytes ...
func (context WebRequestData) UseInputParamBytes() string {
	if context.RequestParams.HasInputParams || context.RequestParams.HasURLPostFix{
		return "inputData, "
	}
	return ""
}

// CloseResponseBody ...
func (context WebRequestData) CloseResponseBody() string {
	if context.ResponseParams.CloseBody {
		return "defer response.Body.Close()"
	}
	return ""
}

// KeepResponseBodyAlive ...
func (context WebRequestData) KeepResponseBodyAlive() string {
	if context.ResponseParams.CloseBody {
		return ""
	}
	return ""
}

// CreateReaderParam ...
func (context WebRequestData) CreateReaderParam() string {
	if context.RequestParams.HasInputParams || context.RequestParams.HasURLPostFix{
		return fmt.Sprintf("\ninputData %sRequest,", context.StructName)
	}
	return ""
}



// CreateReaderParam ...
func (context WebRequestData) URLPostfix() string {
	if context.RequestParams.HasURLPostFix {
		return "inputData.URLPostFix()"	
	}
	return `""`
}






