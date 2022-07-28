package sequence

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type hunting interface {
	getXY(path int) (int, int)
	getPath(x int, y int) int
	addPath(path int, amountGold int)
	getAmountGold(path int) int
	wasNotHere(path int) bool
	couldHuntThere(path int) bool
	updateField(chunk string)
	cloneHunters(h hunter) []hunter
	iteration()
	addResult(result int)
	process() int
}

type hunter struct {
	allPath    []int
	curPath    int
	amountGold int
}

type field struct {
	cells string
	width int
}

type manager struct {
	hunters []hunter
	field   field
	result  int
}

func (f *field) getAmountGold(path int) int {
	ret, _ := strconv.Atoi(string(f.cells[path]))
	//fmt.Println(ret)
	return ret
}

func (f *field) couldHuntThere(path int) bool {
	return path >= 0 && path < len(f.cells) && f.cells[path] != '#'
}

func (f *field) getPath(x int, y int) int {
	//fmt.Println(x, y)
	if x < 0 || y < 0 || x >= f.width || y >= (len(f.cells)/f.width) {
		return -1
	} else if y == 0 {
		return x
	} else {
		return y*f.width + x
	}
}

func (f *field) getXY(path int) (int, int) {
	//fmt.Println(path)
	y := path / f.width
	x := path - f.width*y
	return x, y
}

func (h *hunter) addPath(path int, amountGold int) {
	h.amountGold += amountGold
	h.curPath = path
	//fmt.Println(h.allPath)
	h.allPath = append(h.allPath, path)
	//fmt.Println(h.allPath)
}

func (h *hunter) wasNotHere(path int) bool {
	for _, v := range h.allPath {
		if v == path {
			return false
		}
	}
	return true
}

func (man *manager) addResult(result int) {
	if man.result < result {
		man.result = result
	}
}

func (man *manager) updateField(chunk string) {
	//fmt.Println(chunk)
	man.field.cells = man.field.cells + chunk
	man.field.width = len(chunk)
	position := strings.Index(chunk, "x")
	if position >= 0 {
		curPath := position + len(man.field.cells) - man.field.width
		activeHunter := hunter{allPath: []int{curPath}, curPath: curPath, amountGold: 0}
		man.hunters[0] = activeHunter
	}
}

func (man *manager) cloneHunters(h *hunter) []hunter {
	var clonedHunters []hunter = make([]hunter, 0)
	var clonedHunter hunter
	//fmt.Println("--------------------")
	//fmt.Println("before clone", "allPath", h.allPath, "amountGold", h.amountGold)
	diff_xy := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 0; i < 4; i++ {
		clonedHunter = *h
		clonedHunter.allPath = make([]int, len(h.allPath))
		copy(clonedHunter.allPath, h.allPath)
		x, y := man.field.getXY(clonedHunter.curPath)
		nextPath := man.field.getPath(x+diff_xy[i][0], y+diff_xy[i][1])
		if man.field.couldHuntThere(nextPath) && clonedHunter.wasNotHere(nextPath) {
			//fmt.Println("after clone", "allPath", clonedHunter.allPath, "amountGold", clonedHunter.amountGold)
			//fmt.Println(x+diff_xy[i][0], y+diff_xy[i][1])
			//fmt.Println(nextPath)
			clonedHunter.addPath(nextPath, man.field.getAmountGold(nextPath))
			//fmt.Println("after clone", "allPath", clonedHunter.allPath, "amountGold", clonedHunter.amountGold)
			clonedHunters = append(clonedHunters, clonedHunter)
		}
	}
	return clonedHunters
}

func (man *manager) iteration() {
	//clonedHunters := make([]hunter, 0)
	clonedHunters := []hunter{}
	for i := 0; i < len(man.hunters); i++ {
		cloned3Hunters := man.cloneHunters(&man.hunters[i])
		if len(cloned3Hunters) == 0 {
			//fmt.Println("end", "allPath", man.hunters[i].allPath, "amountGold", man.hunters[i].amountGold)
			man.addResult(man.hunters[i].amountGold)
		} else {
			clonedHunters = append(clonedHunters, cloned3Hunters...)
		}
	}
	man.hunters = clonedHunters
}

func (man *manager) proccess() int {
	t := time.Now().Unix()
	interval := t + 2
	maxHunters, sumHunters, iterations := 0, 0, 0
	for len(man.hunters) > 0 {
		man.iteration()
		iterations++
		sumHunters += len(man.hunters)
		if len(man.hunters) > maxHunters {
			maxHunters = len(man.hunters)
		}
		interval = time.Now().Unix()
		_ = interval
	}
	fmt.Println("duration ", time.Now().Unix()-t)
	fmt.Println("iterations ", iterations)
	fmt.Println("maxHunters ", maxHunters)
	avgHunters := sumHunters / iterations
	fmt.Println("avgHunters ", avgHunters)
	return man.result
}

func getDataFile() []string {
	file, err := os.Open("./static/field.txt")
	if err != nil {
		log.Fatalf("failed to open")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	return text
}

func GetResult() int {
	man := manager{hunters: make([]hunter, 1)}
	for _, chunk := range getDataFile() {
		man.updateField(chunk)
	}
	return man.proccess()
}