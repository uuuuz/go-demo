package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 单个文件夹计算
func main() {
	//rootPath := "/Users/wxm/workspace-demo/goproject/demo_test"
	rootPath := "./coverage_demo"
	//reportDir, codeDir := "cam/ana_code_cvg", "cam/back"
	reportDir, codeDir := rootPath+"/report", rootPath
	// 报告目录
	err := os.MkdirAll(reportDir, os.ModePerm)
	if err != nil {
		fmt.Println("MkdirAll err:", err.Error())
		return
	}
	// 扫描根目录
	dir, err := ioutil.ReadDir(codeDir)
	if err != nil {
		fmt.Println("ReadDir err:", err.Error())
		return
	}
	for _, elem := range dir {
		if elem.IsDir() && !strings.HasPrefix(elem.Name(), ".") {
			if err = analysisCodeCoverage(reportDir+"/"+elem.Name()+".out", reportDir+"/"+elem.Name()+"-ana.txt", codeDir+"/"+elem.Name()); err != nil {
				fmt.Println("analysisCodeCoverage err:", err)
				continue
			}
		}
	}
}

type coverageInfo struct {
	fileName   string
	total      int
	coverage   []int
	unCoverage []int
}

func analysisCodeCoverage(reportFile, analysisFile, codeFolder string) error {
	// 执行 go test
	if err := runGoTest(reportFile, codeFolder); err != nil {
		return err
	}
	// 读取 reportFolder 解析出未被覆盖的部分
	coverageDetails, err := analysisReportFolder(reportFile)
	if err != nil {
		return err
	}
	fmt.Println(coverageDetails)
	// 解析覆盖率 - 根据生成的文件，解析那些方法的覆盖率低
	data, err := pkgAuthor(codeFolder, coverageDetails)
	if err != nil {
		return err
	}
	// 生成解析文件
	return genAnalysisFile(analysisFile, data)
}

func genAnalysisFile(analysisFile string, data []string) error {
	f, err := os.Create(analysisFile)
	if err != nil {
		return err
	}
	defer f.Close()
	for i := range data {
		_, err = f.WriteString(data[i] + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func runGoTest(reportFolder, codeFolder string) error {
	//_, err := execCmd("go", []string{"test", "-gcflags=all=-l", "-coverpkg", codeFolder + "/...",
	//	"-coverprofile", reportFolder, "-timeout", "300s", codeFolder + "/..."})
	//if err != nil {
	//	return fmt.Errorf("%s; %s", codeFolder, err)
	//}
	//return nil

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	var stderr bytes.Buffer
	cmd := exec.CommandContext(ctx, "go", "test", "-gcflags=all=-l", "-coverpkg", codeFolder+"/...",
		"-coverprofile", reportFolder, "-timeout", "300s", codeFolder+"/...")
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s; %s; %s", codeFolder, err, stderr.String())
	}
	return nil
}

func runGitBlame(codeFile string) ([]byte, error) {
	//outData, err := execCmd("git", []string{"blame", codeFile})
	//if err != nil {
	//	return nil, fmt.Errorf("%s; %s", codeFile, err)
	//}
	//return outData, nil

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	var stderr, stdout bytes.Buffer
	cmd := exec.CommandContext(ctx, "git", "blame", codeFile)
	cmd.Stderr, cmd.Stdout = &stderr, &stdout
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("%s; %s; %s", codeFile, err, stderr.String())
	}
	return stdout.Bytes(), nil
}

func execCmd(cmdName string, args []string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	var stderr, stdout bytes.Buffer
	cmd := exec.CommandContext(ctx, cmdName, args...)
	cmd.Stderr, cmd.Stdout = &stderr, &stdout
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("%s; %s", err, stderr.String())
	}
	return stdout.Bytes(), nil
}

func analysisReportFolder(reportFolder string) (map[string]coverageInfo, error) {
	data, err := ioutil.ReadFile(reportFolder)
	if err != nil {
		return nil, err
	}
	res := make(map[string]coverageInfo)
	lines := strings.Split(string(data), "\n")
	for _, line := range lines[1:] {
		if len(line) == 0 {
			continue
		}
		arr1 := strings.Split(line, ":")
		fileName, details := arr1[0], arr1[1]
		arr2 := strings.Split(details, " ")
		start_end, isCovered := arr2[0], arr2[2] == "1"
		arr3 := strings.Split(start_end, ",")
		startLine := strings.Split(arr3[0], ".")[0]
		endLine := strings.Split(arr3[1], ".")[0]
		startLineNum, _ := strconv.Atoi(startLine)
		endLineNum, _ := strconv.Atoi(endLine)
		var lineIndex []int
		for i := startLineNum; i < endLineNum+1; i++ {
			lineIndex = append(lineIndex, i)
		}
		ci, exist := res[fileName]
		if !exist {
			ci = coverageInfo{fileName: fileName}
		}
		if isCovered {
			ci.coverage = append(ci.coverage, lineIndex...)
		} else {
			ci.unCoverage = append(ci.unCoverage, lineIndex...)
		}
		res[fileName] = ci
	}
	fmt.Println(res)
	for k, v := range res {
		// 去重
		v.coverage = duplicateElem(v.coverage)
		v.unCoverage = duplicateElem(v.unCoverage)
		// 排序
		sort.Slice(v.coverage, func(i, j int) bool {
			return v.coverage[i] < v.coverage[j]
		})
		sort.Slice(v.unCoverage, func(i, j int) bool {
			return v.unCoverage[i] < v.unCoverage[j]
		})
		// 求总
		v.total = len(v.coverage) + len(v.unCoverage)
		res[k] = v
	}
	return res, nil
}

func duplicateElem(data []int) []int {
	uniqueMap := make(map[int]struct{})
	for i := range data {
		uniqueMap[i] = struct{}{}
	}
	res := make([]int, 0, len(uniqueMap))
	for k := range uniqueMap {
		res = append(res, k)
	}
	return res
}

func pkgAuthor(codeFolder string, ciMap map[string]coverageInfo) ([]string, error) {
	var res []string
	files, err := ioutil.ReadDir(codeFolder)
	if err != nil {
		return nil, err
	}
	for i := range files {
		if files[i].IsDir() {
			data, err := pkgAuthor(codeFolder+"/"+files[i].Name(), ciMap)
			if err != nil {
				return nil, err
			}
			res = append(res, data...)
			continue
		}
		fileName := codeFolder + "/" + files[i].Name()
		outData, err := runGitBlame(fileName)
		if err != nil {
			return nil, err
		}
		// 00000000 (Not Committed Yet 2023-01-17 18:50:25 +0800  1) package main
		// f6348d18c92 (Youze Tang 2020-02-04 21:51:39 +0900  1) # Crypto Asset Management
		arr1 := strings.Split(string(outData), "\n")
		for j := range arr1 {
			_, after, _ := strings.Cut(arr1[j], "(")
			before, _, _ := strings.Cut(after, ")")
			arr2 := strings.Split(before, " ")
			author, lineStr := arr2[0], arr2[len(arr2)-1]
			lineNum, _ := strconv.Atoi(lineStr)
			if includeTarget(ciMap[fileName].unCoverage, lineNum) {
				res = append(res, fmt.Sprintf("author: %s; line: %d; not covered!", author, lineNum))
			}
		}
	}
	return res, nil
}

func includeTarget(src []int, tar int) bool {
	for i := range src {
		if src[i] == tar {
			return true
		}
	}
	return false
}
