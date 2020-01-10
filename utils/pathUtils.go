package utils

// 获取根路径
func RootPath() string  {
	// os.Args第一个参数是文件名
	//path, err := exec.LookPath(os.Args[0])
	//if err != nil {
	//	log.Panicf("获取基础文件错误:%s",err.Error())
	//}
	//lastIndex := strings.LastIndex(path, "\\")
	//// 切根路径
	//root := path[:lastIndex+1]
	return "D:\\code\\gin\\"
}
