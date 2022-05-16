package codeforces

import (
	"XCPCer_board/model"
	"testing"
)

//检查函数，若不一致返回1
func isIntMsgDifferent(funcRet, ans map[string]int) bool {
	for k, v := range ans {
		if _, r := funcRet[k]; r == false || v != funcRet[k] {
			return true
		}
	}
	return false
}
func isStringMsgDifferent(funcRet, ans map[string]string) bool {
	for k, v := range ans {
		if _, r := funcRet[k]; r == false || v != funcRet[k] {
			return true
		}
	}
	return false
}

//判断并输出错误
func checkIntError(t *testing.T, uid string, tp string, all func(uid string) (map[string]int, error),
	acInt map[string]int) {
	if ret, err := all(uid); isIntMsgDifferent(ret, acInt) {
		if err != nil {
			t.Errorf("Error of %v in all msg: %v", tp, err)
		}
		t.Errorf("Error of %v in all msg\n ret= %v  \nbut the ans is %v", tp, ret, acInt)
	}
}
func checkStrError(t *testing.T, uid string, tp string, all func(uid string) (map[string]string, error),
	acInt map[string]string) {
	if ret, err := all(uid); isStringMsgDifferent(ret, acInt) {
		if err != nil {
			t.Errorf("Error of %v in all msg: %v", tp, err)
		}
		t.Errorf("Error of %v in str msg\n ret= %v  \nbut the ans is %v", tp, ret, acInt)
	}
}

//测试codeforces
func cfTest(t *testing.T) {
	//开始测试
	checkIntError(t, model.TestCodeForcesIdLYF, "CFInt", ScrapeAll, map[string]int{
		lastMonthPassAmount:  0,
		maxRatingKey:         1837,
		ratingKey:            1742,
		problemPassAmountKey: 350,
	})
	checkStrError(t, model.TestCodeForcesIdLYF, "CFString", ScrapeStr, map[string]string{
		rankingNameKey: "Expert ",
	})

}

func TestCF(t *testing.T) {
	cfTest(t)
}
