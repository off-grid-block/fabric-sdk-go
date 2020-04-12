/*
Copyright TCS All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package sigindy
import "testing"


//Test empty input to sigIndy function
func TestEmptyInput(t *testing.T) {
	var a []byte
	var b string

	emptyResult, err := IndySign(a, b)
	if emptyResult != nil {
		t.Errorf("failed, expected nil, but got %v and error is %v", emptyResult, err)
	} else {
		t.Logf("success, expected nil and got %v and error is %v", emptyResult, err)
	}
}

//Test empty did to sigIndy function
func TestEmptyDid(t *testing.T) {
	var a []byte
	var b string
	a = []byte("wqBOesOfW8OjN8KiaBh8wrvDmsKvwrsQw5NNwpfCl8OVwo3DtMODfMOuw53CjhjCqQJYxaERQ1fDkMOSDcOZw4Nsw4HDrWR7w5jDoMKLwq3CocKpVMKpw4TDsDHCtRB/w4bDoMKtAQ==") 
	emptyResult, err := IndySign(a, b)
	if emptyResult != nil {
		t.Errorf("failed, expected nil, but got %v and error is %v", emptyResult, err)
	} else {
		t.Logf("success, expected nil and got %v and error is %v", emptyResult, err)
	}
}

//Test invalid DID size
func TestInvalidDid(t *testing.T) {
	var digest []byte
	var did string
	digest = []byte("input value")
	did = "123456789012345"
	result, err := IndySign(digest, did)
	if result != nil {
		t.Errorf("failed, expected nil, but got %v and error is %v", result, err)
	} else {
		t.Logf("success, expected nil and got %v and error is %v", result, err)
	}
}

//Tests success when correct arguments are given
func TestSuccess(t *testing.T) {
	var digest []byte
	var did string
	digest = []byte("test") //modify the input
	did = "8KHMLmGrxuy1yJ2r7eM3xW"
	result, err := IndySign(digest, did)
	if result == nil {
		t.Errorf("failed, expected signature, but got %v and error is %v", result, err)
	} else {
		t.Logf("success, expected signature and got %v and error is %v", result, err)
	}
}
