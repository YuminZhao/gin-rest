package validate

var Message = map[string](map[string]string){
	"usercreate_name": map[string]string{
		"required": "请输入姓名",
		"unique":   "姓名格式错误",
	},
}
