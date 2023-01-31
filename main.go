package main

import (
	"math/rand"
	"path"
	"runtime"
	"time"

	"github.com/yuuna-stack/go_snake/wrapper"

	"github.com/telroshan/go-sfml/v2/graphics"
	"github.com/telroshan/go-sfml/v2/window"
)

const resourcesDir = "images"

var s1 rand.Source = rand.NewSource(time.Now().UnixNano())
var r1 *rand.Rand = rand.New(s1)

const N = 30
const M = 20
const size = 16
const w = size * N
const h = size * M

var dir int = 0
var num int = 4

type Snake struct {
	x int
	y int
}

var s [100]Snake

type Fruit struct {
	x int
	y int
}

var f Fruit

func tick() {
	for i := num; i > 0; i-- {
		s[i].x = s[i-1].x
		s[i].y = s[i-1].y
	}

	if dir == 0 {
		s[0].y += 1
	}
	if dir == 1 {
		s[0].x -= 1
	}
	if dir == 2 {
		s[0].x += 1
	}
	if dir == 3 {
		s[0].y -= 1
	}

	if s[0].x == f.x && s[0].y == f.y {
		num++
		f.x = r1.Int() % N
		f.y = r1.Int() % M
	}

	if s[0].x > N {
		s[0].x = 0
	}
	if s[0].x < 0 {
		s[0].x = N
	}
	if s[0].y > M {
		s[0].y = 0
	}
	if s[0].y < 0 {
		s[0].y = M
	}

	for i := 1; i < num; i++ {
		if s[0].x == s[i].x && s[0].y == s[i].y {
			num = i
		}
	}
}

func init() { runtime.LockOSThread() }

func fullname(filename string) string {
	return path.Join(resourcesDir, filename)
}

func main() {
	resources := wrapper.Resources{}

	option := uint(window.SfResize | window.SfClose)
	wnd := wrapper.CreateWindow(w, h, "Snake Game!", option, 60)

	sprite1, err := wrapper.FileToSprite(fullname("white.png"), &resources)
	if err != nil {
		panic("Couldn't load white.png")
	}
	sprite2, err := wrapper.FileToSprite(fullname("red.png"), &resources)
	if err != nil {
		panic("Couldn't load red.png")
	}
	sprite3, err := wrapper.FileToSprite(fullname("green.png"), &resources)
	if err != nil {
		panic("Couldn't load green.png")
	}

	timer := 0.0
	delay := 0.1

	f.x = 10
	f.y = 10

	timeStamp := time.Now()

	for wnd.IsOpen() {
		for wnd.Poll_Event() {
			if wnd.Close_Window() {
				return
			}
			if wnd.Key_Pressed() {
				if wnd.Key_Is(window.SfKeyLeft) {
					if dir != 2 {
						dir = 1
					}
				} else if wnd.Key_Is(window.SfKeyRight) {
					if dir != 1 {
						dir = 2
					}
				} else if wnd.Key_Is(window.SfKeyUp) {
					if dir != 0 {
						dir = 3
					}
				} else if wnd.Key_Is(window.SfKeyDown) {
					if dir != 3 {
						dir = 0
					}
				}
			}
		}

		deltaTime := float64(time.Now().Sub(timeStamp).Seconds())
		timeStamp = time.Now()
		timer += deltaTime

		if timer > delay {
			timer = 0
			tick()
		}

		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				sprite1.SetPosition(float32(i*size), float32(j*size))
				sprite1.Draw(wnd.Get_Window())
			}
		}

		for i := 0; i < num; i++ {
			sprite2.SetPosition(float32(s[i].x*size), float32(s[i].y*size))
			sprite2.Draw(wnd.Get_Window())
		}

		sprite3.SetPosition(float32(f.x*size), float32(f.y*size))
		sprite3.Draw(wnd.Get_Window())

		graphics.SfRenderWindow_display(wnd.Get_Window())
	}

	resources.Clear()
	wnd.Clear()
}
