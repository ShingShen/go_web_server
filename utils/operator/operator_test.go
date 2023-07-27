package operator

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	mock_sqloperator "server/tests/mocks"
)

func TestIfElse(t *testing.T) {
	ifElseTrueResforString := IfElse(true, "trueValue", "falseValue")
	if ifElseTrueResforString == "trueValue" {
		t.Logf("ifElseTrueResforString passed. Expected: %v, Got: %v", "trueValue", ifElseTrueResforString)
	} else {
		t.Errorf("ifElseTrueResforString failed. Expected: %v, Got: %v", "trueValue", ifElseTrueResforString)
	}

	ifElseFalseResforString := IfElse(false, "trueValue", "falseValue")
	if ifElseFalseResforString == "falseValue" {
		t.Logf("ifElseFalseResforString passed. Expected: %v, Got: %v", "falseValue", ifElseFalseResforString)
	} else {
		t.Errorf("ifElseFalseResforString failed. Expected: %v, Got: %v", "trueValue", ifElseFalseResforString)
	}

	ifElseTrueResforNil := IfElse(true, nil, "falseValue")
	if ifElseTrueResforNil == nil {
		t.Logf("ifElseTrueResforNil passed. Expected: %v, Got: %v", nil, ifElseTrueResforNil)
	} else {
		t.Errorf("ifElseTrueResforNil failed. Expected: %v, Got: %v", nil, ifElseTrueResforNil)

	}

	ifElseFalseResforNil := IfElse(false, "trueValue", nil)
	if ifElseFalseResforNil == nil {
		t.Logf("ifElseFalseResforNil passed. Expected: %v, Got: %v", nil, ifElseFalseResforNil)
	} else {
		t.Errorf("ifElseFalseResforNil failed. Expected: %v, Got: %v", nil, ifElseFalseResforNil)

	}

	ifElseTrueResforInt := IfElse(true, 37, "falseValue")
	if ifElseTrueResforInt == 37 {
		t.Logf("ifElseTrueResforInt passed. Expected: %v, Got: %v", 37, ifElseTrueResforInt)
	} else {
		t.Errorf("ifElseTrueResforInt failed. Expected: %v, Got: %v", 37, ifElseTrueResforInt)

	}

	ifElseFalseResforInt := IfElse(false, "trueValue", 19)
	if ifElseFalseResforInt == 19 {
		t.Logf("ifElseFalseResforInt passed. Expected: %v, Got: %v", 19, ifElseFalseResforInt)
	} else {
		t.Errorf("ifElseFalseResforInt failed. Expected: %v, Got: %v", 19, ifElseFalseResforInt)

	}

	ifElseTrueResforFloat := IfElse(true, 37.23, "falseValue")
	if ifElseTrueResforFloat == 37.23 {
		t.Logf("ifElseTrueResforFloat passed. Expected: %v, Got: %v", 37.23, ifElseTrueResforFloat)
	} else {
		t.Errorf("ifElseTrueResforFloat failed. Expected: %v, Got: %v", 37.23, ifElseTrueResforFloat)

	}

	ifElseFalseResforFloat := IfElse(false, "trueValue", 19.1)
	if ifElseFalseResforFloat == 19.1 {
		t.Logf("ifElseFalseResforFloat passed. Expected: %v, Got: %v", 19.1, ifElseFalseResforFloat)
	} else {
		t.Errorf("ifElseFalseResforFloat failed. Expected: %v, Got: %v", 19.1, ifElseFalseResforFloat)

	}
}

func TestCreatingDataList(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRows := mock_sqloperator.NewMockISqlRows(ctrl)

	// creatingDataList Col Err
	mockRows.EXPECT().Columns().Return(nil, fmt.Errorf("Columns error"))
	creatingDataListColErr, err := CreatingDataList(mockRows)
	if err != nil {
		t.Logf("creatingDataListColErr passed: %v, %v", creatingDataListColErr, err)
	} else {
		t.Errorf("creatingDataListColErr failed: %v, %v", creatingDataListColErr, err)
	}

	// creatingDataList
	mockRows.EXPECT().Columns().Return([]string{
		"table_id",
		"table_content",
		"created_time",
		"updated_time",
	}, nil)
	mockRows.EXPECT().Next().Return(false)
	creatingDataList, err := CreatingDataList(mockRows)
	if err == nil {
		t.Logf("creatingDataList passed: %v, %v", creatingDataList, err)
	} else {
		t.Errorf("creatingDataList failed: %v, %v", creatingDataList, err)
	}
}
