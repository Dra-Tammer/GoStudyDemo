package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func StoreFiles(SortedArray []int, SortName string) {
	//拿到当下的时间，拼接到文件名的后面，防止重名
	timeNow := time.Now().Unix()
	//拼接文件的存储路径
	fileStr := "D:/WorkPlace/GoWorkPlace/src/StudyDemo/CurriculumDesign/Sort/ResultFile/" +
		SortName + "-" + strconv.FormatInt(timeNow, 10) + ".txt"
	//采用只读的方式访问文件，如果没有这个文件，就创建一个文件
	file, err := os.OpenFile(fileStr, os.O_WRONLY|os.O_CREATE, 06666)
	if err != nil {
		log.Println(err)
		return
	}
	//及时关闭，防止内存泄漏
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	//写入时使用带缓存的*Writer
	writer := bufio.NewWriter(file)
	str := fmt.Sprintln(SortedArray)
	_, WriterErr := writer.WriteString(str)
	if WriterErr != nil {
		fmt.Println(err)
	}
	//因为writer是带缓存的，因此在调用writerstring方法时，
	//内容是写入到缓存的，所有需要调用flush方法，将缓存的数据真正写入到文件中，否则会丢失数据
	fileStoreErr := writer.Flush()
	if err != nil {
		log.Println(fileStoreErr)
	}
	log.Printf("%s文件写入成功", SortName)
}
