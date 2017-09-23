package base

//var logger = Mylog()
import (
	"fmt"
	"strconv"
	"crypto/md5"
	"encoding/hex"
	"os"
)

//sleep time
func Wait(time int) {
	if time <= 0 {
		return
	} else {
		//		logger.Debugf("-- stop ", time, " second ")
		Sleep(time)
	}
}

//string to int
func Atoi(strn string) int{
	num, err := strconv.Atoi(strn)
	if err != nil{
		fmt.Println("atoi error")
	} 
	return num
}

//int to string
func Itoa(num int) string{
	str := strconv.Itoa(num)
	return str
}


func Md5(str string) string{
	h := md5.New()
	h.Write( []byte(str) )
	re := hex.EncodeToString(h.Sum(nil))
	return re
}

//文件是否存在
func FileExist(finame string) bool{
	_, err := os.Stat(finame)
	if err == nil {
		return true
	} else {
		return false
	}
}