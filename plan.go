package main

func getPlanTypeZh(i string) string {
	switch i {
	case "free":
		return "个人免费版"
	case "personal_pro":
		return "个人专业版"
	case "personal_family":
		return "团队版"
	}
	return i
}
func isFree(i string) bool {
	return i == "free"
}
func isPersonalPro(i string) bool {
	return i == "personal_pro"
}
