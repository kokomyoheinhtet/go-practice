package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/google/uuid"
	"math/rand"
	"sort"
	"strings"
	"time"
	"unsafe"
)

type LockStatus struct {
	Status     string `redis:"status"`
	UnlockedAt string `redis:"unlockedAt"`
}

func (ls *LockStatus) updateStatus(status string) {
	ls.Status = status
}

func (ls *LockStatus) print() {
	lsMap := structs.Map(ls)
	for k, v := range lsMap {
		fmt.Println("key is ", k, "and value is ", v)
	}
}

func main() {
	//testEqual()
	//testSort()

	//genRandomNumsWithMathRand()

	//getRandomWithCrypoRand()
	//fmt.Println(RandStringBytesMaskImprSrcUnsafe(6))

	//trySplit()
	fmt.Println(uuid.NewUUID())
	fmt.Println(uuid.NewUUID())
	fmt.Println(uuid.NewUUID())
	fmt.Println(uuid.NewUUID())
	fmt.Println(uuid.NewUUID())
	fmt.Println(uuid.NewRandom())
	fmt.Println(uuid.NewRandom())
	fmt.Println(uuid.NewRandom())
	fmt.Println(uuid.NewRandom())
	fmt.Println(uuid.NewRandom())
}

func trySplit() {
	name := "k"
	fmt.Println(strings.Split(name, "@")[0])
}

//const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const letterBytes = "0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrcUnsafe(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		fmt.Println("loop: ", i)
		fmt.Println("remain 1: ", remain)
		fmt.Println("cache 1: ", cache)
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		fmt.Println("remain 2: ", remain)
		fmt.Println("cache 2: ", cache)
		fmt.Println("int(cache & letterIdxMask): ", int(cache&letterIdxMask))
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			fmt.Println("idx: ", idx)

			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		fmt.Println("cache 3: ", cache)

		remain--
		fmt.Println("remain: ", remain)
		fmt.Println("result: ", b)
		fmt.Println()
	}

	return *(*string)(unsafe.Pointer(&b))
}

func getRandomWithCrypoRand() {
	length := 6
	b := make([]byte, length)

	_, err := rand.Read(b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
	for i := range b {
		fmt.Println(i, int(b[i]))
		//b[i] = OTPChars[int(b[i])%OTPCharsLength]
	}
}

func genRandomNumsWithMathRand() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := 0; i < 10; i++ {
		for i := 0; i < 6; i++ {
			fmt.Print(r.Intn(10))
		}
		fmt.Println()
	}

}

func testSort() {
	input := []string{"bird", "apple", "ocean", "fork", "anchor"}
	fmt.Println(input)
	sort.Strings(input)
	fmt.Println(input)
	i := sort.SearchStrings(input, "attach")
	fmt.Println(i)
}

func testEqual() {
	namea := "koko"
	nameb := "kOkO"
	fmt.Println(namea)
	fmt.Println(nameb)

	fmt.Println(strings.EqualFold(namea, nameb))
}

func testStruct() {
	ls := LockStatus{}
	ls.Status = "ok"
	ls.UnlockedAt = "tmr"
	ls.print()
	ls.updateStatus("new status")
	ls.print()
	lsjson, _ := json.Marshal(ls)
	fmt.Println(string(lsjson))
}
