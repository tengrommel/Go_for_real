package main
import (
	"fmt"
	"time"
	"math/rand"
)

var greetings = [][]string{
	{"Hello, World!", "English"},
	{"Salut Monde", "French"},
	{"世界您好", "Simplified Chinese"},
	{"qo' vIvan", "Klingon"},
	{"हैलो वर्ल्ड", "Hindi"},
	{"안녕하세요", "Korean"},
	{"привет мир", "Russian"},
	{"Wapendwa Dunia", "Swahili"},
	{"Hola Mundo", "Spanish"},
	{"Merhaba Dünya", "Turkish"},
}

func greeting() []string {
	// 生成加密种子
	seed := time.Now().UnixNano()
	// 生成随机数
	rnd := rand.New(rand.NewSource(seed))
	return greetings[rnd.Intn(len(greetings))]
}

func main() {
	g := greeting()
	fmt.Printf("%s (%s)\n", g[0], g[1])
}
