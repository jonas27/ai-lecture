package main

import (
	"log"
	"math/rand"
	"strconv"
)

type Miner struct {
	Sleepy  int
	Thirsty int
	Hungry  int
	Whiskey int
	Gold    int
}

type counts struct {
	sleep   int
	thirst  int
	hunger  int
	whiskey int
	gold    int
}

func main() {
	results := []*counts{}
	goldHighest := 0
	cHighest := &counts{}
	for n := 0; n < 1000; n++ {
		m := &Miner{}
		c := &counts{
			sleep:   rand.Intn(19) + 80,
			thirst:  rand.Intn(19) + 80,
			hunger:  rand.Intn(19) + 80,
			whiskey: rand.Intn(19) + 80,
		}
		for i := 0; i < 1000; i++ {
			if m.areYouDead() {
				log.Println("your are dead")
				log.Println(i)
				log.Println("sleep: " + strconv.Itoa(c.sleep) + " thirst:" + strconv.Itoa(c.thirst) + " hunger: " + strconv.Itoa(c.hunger) + " whiskey: " + strconv.Itoa(c.whiskey))
				log.Println("sleep: " + strconv.Itoa(m.Sleepy) + " thirst:" + strconv.Itoa(m.Thirsty) + " hunger: " + strconv.Itoa(m.Hungry) + " whiskey: " + strconv.Itoa(m.Whiskey) + " gold: " + strconv.Itoa(m.Gold))
				log.Println()
				break
			}
			if m.Sleepy >= c.sleep { // 94
				m.Sleep()
			} else if m.Thirsty >= c.thirst && m.Whiskey == 0 { // 93
				m.buyW()
			} else if m.Thirsty >= c.thirst && m.Whiskey > 0 { // 94
				m.drink()
			} else if m.Hungry >= c.hunger { // 91
				m.eat()
			} else {
				m.Mine()
			}
			if m.areYouDead() {
				log.Println("your are dead")
				log.Println(i)
				log.Println("sleep: " + strconv.Itoa(c.sleep) + " thirst:" + strconv.Itoa(c.thirst) + " hunger: " + strconv.Itoa(c.hunger) + " whiskey: " + strconv.Itoa(c.whiskey))
				log.Println("sleep: " + strconv.Itoa(m.Sleepy) + " thirst:" + strconv.Itoa(m.Thirsty) + " hunger: " + strconv.Itoa(m.Hungry) + " whiskey: " + strconv.Itoa(m.Whiskey) + " gold: " + strconv.Itoa(m.Gold))
				log.Println()
				break
			}
		}
		log.Println(n)
		log.Println("sleep: " + strconv.Itoa(c.sleep) + " thirst:" + strconv.Itoa(c.thirst) + " hunger: " + strconv.Itoa(c.hunger) + " whiskey: " + strconv.Itoa(c.whiskey))
		log.Println("sleep: " + strconv.Itoa(m.Sleepy) + " thirst:" + strconv.Itoa(m.Thirsty) + " hunger: " + strconv.Itoa(m.Hungry) + " whiskey: " + strconv.Itoa(m.Whiskey) + " gold: " + strconv.Itoa(m.Gold))
		log.Println()
		if m.Gold > goldHighest {
			goldHighest = m.Gold
		}
		cHighest = c
		results = append(results, c)
	}

	log.Println()
	log.Println(goldHighest)
	log.Println("sleep: " + strconv.Itoa(cHighest.sleep) + " thirst:" + strconv.Itoa(cHighest.thirst) + " hunger: " + strconv.Itoa(cHighest.hunger) + " whiskey: " + strconv.Itoa(cHighest.whiskey))

}

func (M *Miner) areYouDead() bool {
	if M.Sleepy > 99 || M.Thirsty > 99 || M.Hungry > 99 || M.Whiskey > 10 {
		//log.Println("sleep: " + strconv.Itoa(M.Sleepy) + " thirst:" + strconv.Itoa(M.Thirsty) + " hunger: " + strconv.Itoa(M.Hungry) + " whiskey: " + strconv.Itoa(M.Whiskey) + " gold: " + strconv.Itoa(M.Gold))
		//log.Println()
		return true
	}
	return false
}

func (M *Miner) Sleep() {
	//log.Println("sleep")
	M.Sleepy -= 10
	M.Thirsty++
	M.Hungry++
}
func (M *Miner) buyW() {
	//log.Println("buy")
	M.Sleepy += 5
	M.Thirsty++
	M.Hungry++
	M.Whiskey++
	M.Gold--
}
func (M *Miner) drink() {
	//log.Println("drink")
	M.Sleepy += 5
	M.Thirsty -= 10
	M.Hungry++
	M.Whiskey--
}
func (M *Miner) eat() {
	//log.Println("eat")
	M.Sleepy += 5
	M.Thirsty -= 5
	M.Hungry -= 20
	M.Gold -= 2
}
func (M *Miner) Mine() {
	//log.Println("mine")
	M.Sleepy += 5
	M.Thirsty += 5
	M.Hungry += 5
	M.Gold += 5
}
