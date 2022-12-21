package adventofcode

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Problem2022Day07Part1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	var (
		paths      = make([]string, 0)
		fs         = make(map[string]interface{})
		fsAddDir   func(currentDir map[string]interface{}, paths []string) map[string]interface{}
		fsAddFile  func(currentDir map[string]interface{}, paths []string, size int) map[string]interface{}
		fsPrint    func(currentDir map[string]interface{}, indent int)
		fsSize     func(currentDir map[string]interface{}) int
		fsFindDirs func(currentDir map[string]interface{}) []map[string]interface{}
	)

	fsAddDir = func(currentDir map[string]interface{}, paths []string) map[string]interface{} {
		if len(paths) == 1 {
			if _, ok := currentDir[paths[0]]; !ok {
				currentDir[paths[0]] = make(map[string]interface{})
			}
			return currentDir
		}
		currentDir[paths[0]] = fsAddDir(currentDir[paths[0]].(map[string]interface{}), paths[1:])
		return currentDir
	}
	fsAddFile = func(currentDir map[string]interface{}, paths []string, size int) map[string]interface{} {
		if len(paths) == 1 {
			currentDir[paths[0]] = size
			return currentDir
		}
		currentDir[paths[0]] = fsAddFile(currentDir[paths[0]].(map[string]interface{}), paths[1:], size)
		return currentDir
	}
	fsPrint = func(currentDir map[string]interface{}, indent int) {
		for k, v := range currentDir {
			if dir, ok := v.(map[string]interface{}); ok {
				fmt.Printf("%s - %s (dir)\n", strings.Repeat("  ", indent), k)
				fsPrint(dir, indent+1)
			} else if size, ok := v.(int); ok {
				fmt.Printf("%s - %s (file, size=%d)\n", strings.Repeat("  ", indent), k, size)
			}
		}
	}
	fsSize = func(currentDir map[string]interface{}) int {
		size := 0
		for _, v := range currentDir {
			if dir, ok := v.(map[string]interface{}); ok {
				size += fsSize(dir)
			} else if s, ok := v.(int); ok {
				size += s
			}
		}
		return size
	}
	fsFindDirs = func(currentDir map[string]interface{}) []map[string]interface{} {
		dirs := make([]map[string]interface{}, 0)
		for _, v := range currentDir {
			if dir, ok := v.(map[string]interface{}); ok {
				dirs = append(dirs, dir)
				dirs = append(dirs, fsFindDirs(dir)...)
			}
		}
		return dirs
	}
	for {
		raw, _, err := in.ReadLine()
		if err == io.EOF {
			break
		}
		line := string(raw)
		if strings.HasPrefix(line, "$ cd ") {
			// change directory
			toDir := strings.Replace(line, "$ cd ", "", 1)
			if toDir == ".." {
				paths = paths[:len(paths)-1]
			} else {
				paths = append(paths, toDir)
			}
			fs = fsAddDir(fs, paths)
			continue
		} else if strings.HasPrefix(line, "$ ls") {
			// list
		} else if strings.HasPrefix(line, "dir ") {
			// result of ls
			// but dir
		} else {
			buffer := strings.Split(line, " ")
			size, _ := strconv.Atoi(buffer[0])
			fs = fsAddFile(fs, append(paths, buffer[1]), size)
		}
	}
	result := 0
	for _, dir := range fsFindDirs(fs) {
		dirSize := fsSize(dir)
		if dirSize > 100000 {
			continue
		}
		result += dirSize
	}
	fmt.Fprint(out, result)
}

func Problem2022Day07Part2(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	var (
		paths      = make([]string, 0)
		fs         = make(map[string]interface{})
		fsAddDir   func(currentDir map[string]interface{}, paths []string) map[string]interface{}
		fsAddFile  func(currentDir map[string]interface{}, paths []string, size int) map[string]interface{}
		fsPrint    func(currentDir map[string]interface{}, indent int)
		fsSize     func(currentDir map[string]interface{}) int
		fsFindDirs func(currentDir map[string]interface{}) []map[string]interface{}
	)

	fsAddDir = func(currentDir map[string]interface{}, paths []string) map[string]interface{} {
		if len(paths) == 1 {
			if _, ok := currentDir[paths[0]]; !ok {
				currentDir[paths[0]] = make(map[string]interface{})
			}
			return currentDir
		}
		currentDir[paths[0]] = fsAddDir(currentDir[paths[0]].(map[string]interface{}), paths[1:])
		return currentDir
	}
	fsAddFile = func(currentDir map[string]interface{}, paths []string, size int) map[string]interface{} {
		if len(paths) == 1 {
			currentDir[paths[0]] = size
			return currentDir
		}
		currentDir[paths[0]] = fsAddFile(currentDir[paths[0]].(map[string]interface{}), paths[1:], size)
		return currentDir
	}
	fsPrint = func(currentDir map[string]interface{}, indent int) {
		for k, v := range currentDir {
			if dir, ok := v.(map[string]interface{}); ok {
				fmt.Printf("%s - %s (dir)\n", strings.Repeat("  ", indent), k)
				fsPrint(dir, indent+1)
			} else if size, ok := v.(int); ok {
				fmt.Printf("%s - %s (file, size=%d)\n", strings.Repeat("  ", indent), k, size)
			}
		}
	}
	fsSize = func(currentDir map[string]interface{}) int {
		size := 0
		for _, v := range currentDir {
			if dir, ok := v.(map[string]interface{}); ok {
				size += fsSize(dir)
			} else if s, ok := v.(int); ok {
				size += s
			}
		}
		return size
	}
	fsFindDirs = func(currentDir map[string]interface{}) []map[string]interface{} {
		dirs := make([]map[string]interface{}, 0)
		for _, v := range currentDir {
			if dir, ok := v.(map[string]interface{}); ok {
				dirs = append(dirs, dir)
				dirs = append(dirs, fsFindDirs(dir)...)
			}
		}
		return dirs
	}
	for {
		raw, _, err := in.ReadLine()
		if err == io.EOF {
			break
		}
		line := string(raw)
		if strings.HasPrefix(line, "$ cd ") {
			// change directory
			toDir := strings.Replace(line, "$ cd ", "", 1)
			if toDir == ".." {
				paths = paths[:len(paths)-1]
			} else {
				paths = append(paths, toDir)
			}
			fs = fsAddDir(fs, paths)
			continue
		} else if strings.HasPrefix(line, "$ ls") {
			// list
		} else if strings.HasPrefix(line, "dir ") {
			// result of ls
			// but dir
		} else {
			buffer := strings.Split(line, " ")
			size, _ := strconv.Atoi(buffer[0])
			fs = fsAddFile(fs, append(paths, buffer[1]), size)
		}
	}
	const (
		totalDiskSpace     = 70000000
		requiredEmptySpace = 30000000
	)
	usedSpace := fsSize(fs)
	emptySpace := totalDiskSpace - usedSpace
	minSpaceToDelete := requiredEmptySpace - emptySpace
	result := 0
	if minSpaceToDelete > 0 {
		minFolderSizeToDelete := usedSpace
		for _, dir := range fsFindDirs(fs) {
			dirSize := fsSize(dir)
			if dirSize < minSpaceToDelete {
				continue
			}
			if dirSize < minFolderSizeToDelete {
				minFolderSizeToDelete = dirSize
			}
		}
		result = minFolderSizeToDelete
	}
	fmt.Fprint(out, result)
}
