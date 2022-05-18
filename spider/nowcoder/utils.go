package nowcoder

import "fmt"

// @Author: Feng
// @Date: 2022/5/16 17:48

func getRatingKey(uid string) string {
	return fmt.Sprintf("%v_%v", ratingKey, uid)
}

func getRankingKey(uid string) string {
	return fmt.Sprintf("%v_%v", rankingKey, uid)
}

func getContestAmountKey(uid string) string {
	return fmt.Sprintf("%v_%v", contestAmountKey, uid)
}

func getPassAmountKey(uid string) string {
	return fmt.Sprintf("%v_%v", passAmountKey, uid)
}

//getContestProfileUrl 获取牛客竞赛区个人主页URL
func getContestProfileUrl(nowCoderId string) string {
	return "https://ac.nowcoder.com/acm/contest/profile/" + nowCoderId
}

//getContestPracticeUrl 获取牛客竞赛区个人练习URL
func getContestPracticeUrl(nowCoderId string) string {
	return getContestProfileUrl(nowCoderId) + "/practice-coding"
}

//getNowCoderContestBaseFindRule 获取牛客竞赛区基础的
func getNowCoderContestBaseFindRule(keyWord string) string {
	return fmt.Sprintf(".my-state-item:contains(%v) .state-num", keyWord)
}
