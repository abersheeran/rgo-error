package rgo_error

import "testing"

func TestResult(t *testing.T) {
	if res := result_ok(); res.IsOk() {
		t.Log(res.Unwrap())
	} else {
		t.Error(res.Error())
	}

	if res := result_err(); res.IsErr() {
		t.Log(res.Error())
	} else {
		t.Error(res.Unwrap())
	}
}

func result_ok() (res Result[string]) {
	return res.Ok("hello world")
}

func result_err() (res Result[string]) {
	return res.Err("Error message")
}
