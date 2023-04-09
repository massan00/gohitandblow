package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

// HitAndBlow本体を作成
type HitAndBlow struct {
	answerSource []int
	answer       []int
	tryCount     int
	mode         int
}

// 解答の型
type CorrectCheck struct {
	hits  int
	blows int
}

func NewHitAndBlow() HitAndBlow {
	count := setMode()
	hitandblow := HitAndBlow{
		answerSource: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		tryCount:     0,
		mode:         count,
	}
	var answer []int
	for len(answer) < count {
		n := hitandblow.answerSource[rand.Intn(10)]
		if !slices.Contains(answer, n) {
			answer = append(answer, n)
		}
	}
	hitandblow.answer = answer

	return hitandblow
}

// 入力からモードを決定する
func setMode() int {
	var mode string
	fmt.Println("Hit&Blowは数字当てゲームです。")
	fmt.Println("詳しい説明はネットで調べてください。")
	fmt.Println("ノーマルモードは3桁の数字、ハードモードは4桁の数字を当てます。")
	fmt.Printf("モードを選択してください\nノーマル:1\nハード:2\n> ")
	fmt.Scan(&mode)
	num, err := strconv.Atoi(mode)
	if err != nil {
		fmt.Println("1 or 2を入力してください。")
		return setMode()
	}
	if num == 1 {
		return 3
	} else if num == 2 {
		return 4
	} else {
		fmt.Println("1 or 2を入力してください。")
		return setMode()
	}
}

// 入力からh.answerにわたすスライスを生成
func setAnswer(num int) []int {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%d個カンマ区切りで数値を入力してください\n> ", num)
	scanner.Scan()
	input := scanner.Text()
	parts := strings.Split(input, ",")

	if len(parts) != int(num) {
		fmt.Println("入力が正しくありません")
		return setAnswer(num)
	}
	answers := make([]int, num)
	for i, v := range parts {
		n, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			fmt.Println("入力が正しくありません")
			return setAnswer(num)
		}
		//nが0-9の範囲にあるかチェック
		if n < 0 || n > 9 {
			fmt.Println("0から9の範囲で入力してください")
			return setAnswer(num)
		} else {
			answers[i] = n
		}
	}
	return answers
}

func (h *HitAndBlow) Start() {
	for {
		input := setAnswer(len(h.answer))
		h.tryCount++
		c := CorrectCheck{}
		for i := 0; i < h.mode; i++ {
			ans := h.answer[i]
			if ans == input[i] {
				c.hits++
			} else if slices.Contains(input, ans) {
				c.blows++
			}
		}
		if c.hits == h.mode {
			fmt.Printf("正解です!\n試行回数は%d回です\n", h.tryCount)
			break
		} else {
			fmt.Printf("%dヒット%dブローです\n", c.hits, c.blows)
		}
	}
}

func main() {
	h := NewHitAndBlow()
	h.Start()

}
