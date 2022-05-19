package atcoder

import (
	"fmt"
	"strconv"
)

func getRatingKey(uid string) string {
	return fmt.Sprintf("%v_%v", RatingKey, uid)
}

func getRankKey(uid string) string {
	return fmt.Sprintf("%v_%v", rankKey, uid)
}

func getContestSumKey(uid string) string {
	return fmt.Sprintf("%v_%v", contestSumKey, uid)
}

func getSubmissionKey(cid string, pid string) string {
	return fmt.Sprintf("%v_%v_%v", submissionKey, cid, pid)
}

func getPageUrl(page int) string {
	return "https://atcoder.jp/contests/archive?page=" + strconv.Itoa(page)
}

func getSubmissionPageUrl(cid string, uid string) string {
	return "https://atcoder.jp/contests/" + cid + "/submissions?f.User=" + uid + "&f.Status=AC"
}

//getAtCoderBaseUrl 获取个人主页URL
func getAtCoderBaseUrl(atCoderId string) string {
	return "https://atcoder.jp/users/" + atCoderId
}
