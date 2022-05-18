package codeforces

//---------------------------------------------------------------------//
// 共用函数
//---------------------------------------------------------------------//

func getPersonPage(uid string) string {
	return "https://codeforces.com/profile/" + uid
}
func getPersonProblemPage(uid string) string {
	return "https://codeforces.com/submissions/" + uid
}
