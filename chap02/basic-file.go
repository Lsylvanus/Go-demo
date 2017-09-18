package main

import (
	"os"
	"path/filepath"
	"time"
	"strings"
	"log"
	"io/ioutil"
	"os/exec"
	"path"
)

const (
	IsDirectory = iota
	IsRegular
	IsSymlink
)

type sysFile struct {
	fType  int
	fName  string
	fLink  string
	fSize  int64
	fMtime time.Time
	fPerm  os.FileMode
}
type F struct {
	files []*sysFile
}

func (self *F) visit(path string, f os.FileInfo, err error) error {
	if f == nil {
		return err
	}
	var tp int
	if f.IsDir() {
		tp = IsDirectory
	} else if (f.Mode() & os.ModeSymlink) > 0 {
		tp = IsSymlink
	} else {
		tp = IsRegular
	}
	inoFile := &sysFile{
		fName:  path,
		fType:  tp,
		fPerm:  f.Mode(),
		fMtime: f.ModTime(),
		fSize:  f.Size(),
	}
	self.files = append(self.files, inoFile)
	return nil
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}
//获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}

func main() {
	/*flag.Parse()
	root := flag.Arg(0)
	self := F{
		files: make([]*sysFile, 0),
	}
	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		return self.visit(path, f, err)
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	for _, v := range self.files {
		fmt.Println(v.fName, v.fSize)
	}*/

	/*var str1, str2 string
	str1 = getCurrentDirectory()
	fmt.Println(str1)
	str2 = getParentDirectory(str1)
	fmt.Println(str2)*/

	/*files, err := ListDir("D:\\_downs\\lsyl", ".exe")
	fmt.Println(files, err)
	files, err = WalkDir("D:\\work", ".doc")
	fmt.Println(files, err)*/

	execDirAbsPath, _ := os.Getwd()
	log.Println("执行程序所在目录的绝对路径　　　　　　　:", execDirAbsPath)

	execFileRelativePath, _ := exec.LookPath(os.Args[0])
	log.Println("执行程序与命令执行目录的相对路径　　　　:", execFileRelativePath)

	execDirRelativePath, _ := path.Split(execFileRelativePath)
	log.Println("执行程序所在目录与命令执行目录的相对路径:", execDirRelativePath)

	execFileAbsPath, _ := filepath.Abs(execFileRelativePath)
	log.Println("执行程序的绝对路径　　　　　　　　　　　:", execFileAbsPath)

	execDirAbsPath, _ = filepath.Abs(execDirRelativePath)
	log.Println("执行程序所在目录的绝对路径　　　　　　　:", execDirAbsPath)

	os.Chdir(execDirRelativePath) //进入目录
	enteredDirAbsPath, _ := os.Getwd()
	log.Println("所进入目录的绝对路径　　　　　　　　　　:", enteredDirAbsPath)
}
