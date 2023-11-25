package main

import (
	"fmt"
	"math"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var total int = 0
var beams = []int{3, 3, 3, 2, 4, 4, 4, 4, 4, -1}

var perfectMatches map[string]string
var noMatches map[string]string
var matchups [9]map[string]string

var pairs map[string]map[string]int

func initPairs() {

	pairs = make(map[string]map[string]int)

	asia := make(map[string]int)
	asia["andrew"] = 0
	asia["brett"] = 0
	asia["cam"] = 0
	asia["daniel"] = 0
	asia["kwasi"] = 0
	asia["lewis"] = 0
	asia["moe"] = 0
	asia["shamoy"] = 0
	asia["tevin"] = 0
	asia["tomas"] = 0
	asia["zak"] = 0
	pairs["asia"] = asia

	bria := make(map[string]int)
	bria["andrew"] = 0
	bria["brett"] = 0
	bria["cam"] = 0
	bria["daniel"] = 0
	bria["kwasi"] = 0
	bria["lewis"] = 0
	bria["moe"] = 0
	bria["shamoy"] = 0
	bria["tevin"] = 0
	bria["tomas"] = 0
	bria["zak"] = 0
	pairs["bria"] = bria

	cali := make(map[string]int)
	cali["andrew"] = 0
	cali["brett"] = 0
	cali["cam"] = 0
	cali["daniel"] = 0
	cali["kwasi"] = 0
	cali["lewis"] = 0
	cali["moe"] = 0
	cali["shamoy"] = 0
	cali["tevin"] = 0
	cali["tomas"] = 0
	cali["zak"] = 0
	pairs["cali"] = cali

	jasmine := make(map[string]int)
	jasmine["andrew"] = 0
	jasmine["brett"] = 0
	jasmine["cam"] = 0
	jasmine["daniel"] = 0
	jasmine["kwasi"] = 0
	jasmine["lewis"] = 0
	jasmine["moe"] = 0
	jasmine["shamoy"] = 0
	jasmine["tevin"] = 0
	jasmine["tomas"] = 0
	jasmine["zak"] = 0
	pairs["jasmine"] = jasmine

	kayla := make(map[string]int)
	kayla["andrew"] = 0
	kayla["brett"] = 0
	kayla["cam"] = 0
	kayla["daniel"] = 0
	kayla["kwasi"] = 0
	kayla["lewis"] = 0
	kayla["moe"] = 0
	kayla["shamoy"] = 0
	kayla["tevin"] = 0
	kayla["tomas"] = 0
	kayla["zak"] = 0
	pairs["kayla"] = kayla

	kenya := make(map[string]int)
	kenya["andrew"] = 0
	kenya["brett"] = 0
	kenya["cam"] = 0
	kenya["daniel"] = 0
	kenya["kwasi"] = 0
	kenya["lewis"] = 0
	kenya["moe"] = 0
	kenya["shamoy"] = 0
	kenya["tevin"] = 0
	kenya["tomas"] = 0
	kenya["zak"] = 0
	pairs["kenya"] = kenya

	lauren := make(map[string]int)
	lauren["andrew"] = 0
	lauren["brett"] = 0
	lauren["cam"] = 0
	lauren["daniel"] = 0
	lauren["kwasi"] = 0
	lauren["lewis"] = 0
	lauren["moe"] = 0
	lauren["shamoy"] = 0
	lauren["tevin"] = 0
	lauren["tomas"] = 0
	lauren["zak"] = 0
	pairs["lauren"] = lauren

	maria := make(map[string]int)
	maria["andrew"] = 0
	maria["brett"] = 0
	maria["cam"] = 0
	maria["daniel"] = 0
	maria["kwasi"] = 0
	maria["lewis"] = 0
	maria["moe"] = 0
	maria["shamoy"] = 0
	maria["tevin"] = 0
	maria["tomas"] = 0
	maria["zak"] = 0
	pairs["maria"] = maria

	morgan := make(map[string]int)
	morgan["andrew"] = 0
	morgan["brett"] = 0
	morgan["cam"] = 0
	morgan["daniel"] = 0
	morgan["kwasi"] = 0
	morgan["lewis"] = 0
	morgan["moe"] = 0
	morgan["shamoy"] = 0
	morgan["tevin"] = 0
	morgan["tomas"] = 0
	morgan["zak"] = 0
	pairs["morgan"] = morgan

	nutsa := make(map[string]int)
	nutsa["andrew"] = 0
	nutsa["brett"] = 0
	nutsa["cam"] = 0
	nutsa["daniel"] = 0
	nutsa["kwasi"] = 0
	nutsa["lewis"] = 0
	nutsa["moe"] = 0
	nutsa["shamoy"] = 0
	nutsa["tevin"] = 0
	nutsa["tomas"] = 0
	nutsa["zak"] = 0
	pairs["nutsa"] = nutsa

	samantha := make(map[string]int)
	samantha["andrew"] = 0
	samantha["brett"] = 0
	samantha["cam"] = 0
	samantha["daniel"] = 0
	samantha["kwasi"] = 0
	samantha["lewis"] = 0
	samantha["moe"] = 0
	samantha["shamoy"] = 0
	samantha["tevin"] = 0
	samantha["tomas"] = 0
	samantha["zak"] = 0
	pairs["samantha"] = samantha

}

func MatchupCeremonies() {

	// truth booths
	perfectMatches = make(map[string]string)
	noMatches = make(map[string]string)
	//noMatches["asia"] = "andrew"
	//noMatches["kenya"] = "brett"
	//noMatches["bria"] = "zak"
	//noMatches["cali"] = "tomas"
	//noMatches["nutsa"] = "zak"

	// week 1
	noMatches["maria"] = "tomas"     // truth booth
	week1 := make(map[string]string) //matchup
	week1["asia"] = "kwasi"
	week1["bria"] = "zak"
	week1["cali"] = "brett"
	week1["jasmine"] = "moe"
	week1["kayla"] = "cam"
	week1["kenya"] = "tevin"
	week1["lauren"] = "andrew"
	week1["maria"] = "shamoy"
	week1["morgan"] = "tomas"
	week1["nutsa"] = "daniel"
	week1["samantha"] = "lewis"
	matchups[0] = week1

	// week 2
	noMatches["asia"] = "andrew"     // truth booth
	week2 := make(map[string]string) //matchup
	week2["asia"] = "brett"
	week2["bria"] = "moe"
	week2["cali"] = "tomas"
	week2["jasmine"] = "lewis"
	week2["kayla"] = "cam"
	week2["kenya"] = "tevin"
	week2["lauren"] = "kwasi"
	week2["maria"] = "shamoy"
	week2["morgan"] = "andrew"
	week2["nutsa"] = "daniel"
	week2["samantha"] = "zak"
	matchups[1] = week2

	// week 3
	perfectMatches["maria"] = "shamoy" // truth booth
	week3 := make(map[string]string)   //matchup
	week3["asia"] = "lewis"
	week3["bria"] = "tomas"
	week3["cali"] = "brett"
	week3["jasmine"] = "kwasi"
	week3["kayla"] = "cam"
	week3["kenya"] = "tevin"
	week3["lauren"] = "andrew"
	week3["maria"] = "shamoy"
	week3["morgan"] = "zak"
	week3["nutsa"] = "moe"
	week3["samantha"] = "daniel"
	matchups[2] = week3

	// week 4
	noMatches["kenya"] = "brett"     // truth booth
	week4 := make(map[string]string) //matchup
	week4["asia"] = "cam"
	week4["bria"] = "kwasi"
	week4["cali"] = "tomas"
	week4["jasmine"] = "tevin"
	week4["kayla"] = "brett"
	week4["kenya"] = "lewis"
	week4["lauren"] = "daniel"
	week4["maria"] = "shamoy"
	week4["morgan"] = "zak"
	week4["nutsa"] = "andrew"
	week4["samantha"] = "moe"
	matchups[3] = week4

	// week 5
	noMatches["bria"] = "zak"        // truth booth
	week5 := make(map[string]string) //matchup
	week5["asia"] = "moe"
	week5["bria"] = "daniel"
	week5["cali"] = "tomas"
	week5["jasmine"] = "kwasi"
	week5["kayla"] = "cam"
	week5["kenya"] = "tevin"
	week5["lauren"] = "lewis"
	week5["maria"] = "shamoy"
	week5["morgan"] = "zak"
	week5["nutsa"] = "brett"
	week5["samantha"] = "andrew"
	matchups[4] = week5

	// week 6
	noMatches["cali"] = "brett"      // truth booth
	week6 := make(map[string]string) //matchup
	week6["asia"] = "kwasi"
	week6["bria"] = "lewis"
	week6["cali"] = "tomas"
	week6["jasmine"] = "moe"
	week6["kayla"] = "cam"
	week6["kenya"] = "tevin"
	week6["lauren"] = "andrew"
	week6["maria"] = "shamoy"
	week6["morgan"] = "zak"
	week6["nutsa"] = "brett"
	week6["samantha"] = "daniel"
	matchups[5] = week6

	// week 7
	noMatches["nutsa"] = "zak"       // truth booth
	week7 := make(map[string]string) //matchup
	week7["asia"] = "kwasi"
	week7["bria"] = "lewis"
	week7["cali"] = "cam"
	week7["jasmine"] = "moe"
	week7["kayla"] = "tomas"
	week7["kenya"] = "tevin"
	week7["lauren"] = "andrew"
	week7["maria"] = "shamoy"
	week7["morgan"] = "zak"
	week7["nutsa"] = "brett"
	week7["samantha"] = "daniel"
	matchups[6] = week7

	// week 8
	perfectMatches["kenya"] = "tevin" // truth booth
	week8 := make(map[string]string)  //matchup
	week8["asia"] = "daniel"
	week8["bria"] = "lewis"
	week8["cali"] = "zak"
	week8["jasmine"] = "kwasi"
	week8["kayla"] = "moe"
	week8["kenya"] = "tevin"
	week8["lauren"] = "cam"
	week8["maria"] = "shamoy"
	week8["morgan"] = "tomas"
	week8["nutsa"] = "brett"
	week8["samantha"] = "andrew"
	matchups[7] = week8

	// week 9
	noMatches["samantha"] = "cam"    // truth booth
	week9 := make(map[string]string) //matchup
	week9["asia"] = "lewis"
	week9["bria"] = "brett"
	week9["cali"] = "andrew"
	week9["jasmine"] = "tomas"
	week9["kayla"] = "moe"
	week9["kenya"] = "tevin"
	week9["lauren"] = "daniel"
	week9["maria"] = "shamoy"
	week9["morgan"] = "cam"
	week9["nutsa"] = "kwasi"
	week9["samantha"] = "zak"
	matchups[8] = week9

	// week 10
	perfectMatches["nutsa"] = "brett" // truth booth

}

func main() {

	// init contestant list
	guys := []string{"andrew", "brett", "cam", "daniel", "kwasi", "lewis", "moe", "shamoy", "tevin", "tomas", "zak"}
	girls := []string{"asia", "bria", "cali", "jasmine", "kayla", "kenya", "lauren", "maria", "morgan", "nutsa", "samantha"}
	initPairs()

	MatchupCeremonies()

	confirmed(girls, guys, make(map[string]string))
	for girl, guy := range pairs {
		fmt.Printf("%s:\n", girl)
		for dude, num := range guy {
			fmt.Printf("%s-%.1f%%\n", dude, math.Ceil((float64(num)/float64(total))*100*10)/10)
		}
		fmt.Println()
	}

	p := message.NewPrinter(language.English)
	p.Printf("Total possible combinations: %d", total)
}

func confirmed(girls []string, guys []string, matches map[string]string) {

	var girlIndex int
	var guyIndex int
	for girl, guy := range perfectMatches {
		matches[girl] = guy
		girlIndex = FindIndex(girls, girl)
		girls = RemoveIndex(girls, girlIndex)
		guyIndex = FindIndex(guys, guy)
		guys = RemoveIndex(guys, guyIndex)
	}

	match(girls, guys, matches)

}

func match(girls []string, guys []string, matches map[string]string) {

	if len(girls) == 1 && len(guys) == 1 {
		matches[girls[0]] = guys[0]
		if checkBeams(matches) {
			updateMatches(matches)
			total++
		}
		return
	}
	for i, guy := range guys {
		matches[girls[0]] = guy
		if noMatches[girls[0]] != guy {
			match(girls[1:], RemoveIndex(guys, i), matches)
		}
	}

}

func RemoveIndex(g []string, index int) []string {
	tmp := make([]string, len(g))
	copy(tmp, g)
	return append(tmp[:index], tmp[index+1:]...)
}

func FindIndex(arr []string, elem string) int {
	for index, value := range arr {
		if value == elem {
			return index
		}
	}
	return -1
}

func checkBeams(matches map[string]string) bool {

	var weeklyBeams int
	for index, numBeams := range beams {
		if numBeams < 0 {
			break
		}
		ceremony := matchups[index]
		weeklyBeams = 0
		for girl, guy := range matches {
			if guy == ceremony[girl] {
				weeklyBeams++
			}
		}
		if weeklyBeams != numBeams {
			return false
		}
	}
	return true
}

func updateMatches(matches map[string]string) {

	for girl, guy := range matches {
		pairs[girl][guy]++
	}
}
