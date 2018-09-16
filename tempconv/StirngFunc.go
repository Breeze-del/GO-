package tempconv

import "strings"

//判断s字符串是否含有prefix的前缀
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

//判断s字符串是否含有suffix后缀
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

//判断是不是substr 是不是s的子串
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

//去掉像是路径的前缀和像是文件格式的后缀 a/b/c.go --> c
func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] =='/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] =='.' {
			s = s[:i]
			break
		}
	}
	return s
}

//调用string.LastIndex库函数
func simpleBasename(s string) string {
	//如果没有找到"/" 那么slash值为-1 否则返回“/”的位置下标
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[: dot]
	}
	return s
}
//[]string 字符串切片slice是不允许改变值得，  这是因为为了保护两个切片访问
//同一哥底层string得时候，安全相同  但是是可以在中间加上字符得

//没三个字符加上一个","
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	//递归
	return comma(s[:n-3]+","+s[n-3:])
}

func stringPrefix(x string) string {
	return x
}