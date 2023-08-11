package errors

import (
	"errors"
	perrors "errors"
	"fmt"
	"testing"
)

func TestGetCode(t *testing.T) {
	code := 10101
	err := NewWithCode(code, "test code")
	werr := Wrap(err, "sss")
	if code != GetCode(werr) {
		t.Log("Failed to get code")
		t.Failed()
	}
}

func TestWrapErrorIs(t *testing.T) {
	var errFalse = NewWithCode(404, "False")
	var errOrigin = NewWithCode(200, "Origin")
	err := Wrap(errOrigin, "Wrap Test")
	if !errors.Is(err, errOrigin) {
		t.Log("Failed to errors.Is")
		t.Fail()
	}

	if errors.Is(err, errFalse) {
		t.Log("False Match @ errors.Is")
		t.Fail()
	}
}

func TestWrapErrorAs(t *testing.T) {
	var errFalse = errors.New("sdfsdf")
	var errOrigin = NewWithCode(200, "Origin")
	err := Wrap(errOrigin, "Wrap Test")
	// if !errors.As(err, &errOrigin) {
	// 	t.Log("Failed to errors.As")
	// 	t.Fail()
	// }
	// errOrigin = NewWithCode(200,"Origin")
	if errors.As(err, &errFalse) {
		t.Log("False Match @ errors.As")
		t.Fail()
	}
}

func TestWrapddd(t *testing.T) {
	var errO = errors.New("sdfsdf")
	var errW = Wrap(errO, "xxx")
	fmt.Println(errors.Is(errW, errO))
	if !errors.Is(errW, errO) {
		t.Log("errW is not err")
		t.Failed()
	}
}

func TestWithError(t *testing.T) {
	var errFalse = NewWithCode(404, "False")
	var errOrigin = NewWithCode(200, "test")
	var errTest2 = NewWithCode(201, "test2")
	err := With(errOrigin, errTest2)
	if !perrors.As(err, &errOrigin) {
		t.Log("Failed to errors.As")
		t.Fail()
	}
	if !perrors.Is(err, errOrigin) {
		t.Log("Failed to errors.Is")
		t.Fail()
	}
	if !perrors.As(err, &errTest2) {
		t.Log("Failed to errors.As")
		t.Fail()
	}
	if !perrors.Is(err, errTest2) {
		t.Log("Failed to errors.Is")
		t.Fail()
	}
	if perrors.As(err, &errFalse) {
		t.Log("False Match @ errors.As")
		t.Fail()
	}
	if perrors.Is(err, errFalse) {
		t.Log("False Match @ errors.Is")
		t.Fail()
	}
}
