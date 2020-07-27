package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	patterns, paths := readStdinSync()
	getBestMatches(patterns, paths)
}

func readStdinSync() ([]string, []string) {
	scanner := bufio.NewScanner(os.Stdin)

	patterns := []string{}
	paths := []string{}
	lineCounter := 0
	patternCount := 0

	seperatePatternsAndPaths := func(line string) {

		// fmt.Println("lineCounter", lineCounter)

		if lineCounter == 0 {
			// Convert incoming String to type int
			line, err := strconv.Atoi(line)
			if err != nil {
				// handle error
				fmt.Println("ERROR", err)
				os.Exit(2)
			}
			patternCount = line
			lineCounter++
			return
			//fmt.Println("type", reflect.TypeOf(line))

		}

		// Set amount of incoming paths
		if lineCounter-1 == patternCount {
			fmt.Println("Here we set incoming paths")
			lineCounter++
			return
		}

		if lineCounter <= patternCount {
			// Push pattern to patterns slice
			patterns = append(patterns, scanner.Text())
		}

		if lineCounter > patternCount {
			paths = append(paths, scanner.Text())
		}

		lineCounter++
	}

	// Sort each incoming line into pattern or path
	for scanner.Scan() {
		seperatePatternsAndPaths(scanner.Text())
		//fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return patterns, paths
}

func getBestMatches(patterns []string, paths []string) {
	fmt.Println("patterns", patterns)
	fmt.Println("paths", paths)

	pTrie := generatePatternTrie(false)
	pTrie.printTrie()

	// fmt.Println("pTrie", pTrie)

	for i := 0; i < len(patterns); i++ {
		//fmt.Println(patterns[i])
		pTrie.addPattern(patterns[i])
	}

	// n := newNode("test")
	// m := newNode("other")
	// fmt.Println("N", n)
	// fmt.Println("m", m)
	// fmt.Println("N", n)

	// fmt.Println("Test func running")
}

type node struct {
	next   map[string]*node
	data   string
	isLast struct{}
}

type patternTrie struct {
	root       node
	tieBreaker bool
}

func generatePatternTrie(tieBreaker bool) *patternTrie {
	p := new(patternTrie)
	p.tieBreaker = tieBreaker
	p.root = *newNode("root")
	fmt.Println("p.root", p.root)
	return p
}

func newNode(data string) *node {
	n := new(node)
	n.data = data
	return n
}

func (p patternTrie) printTrie() {
	fmt.Println("tieBreaker", p.tieBreaker)
}

func (p patternTrie) addPattern(pattern string) {

	tokenized := strings.Split(pattern, ",")

	findInjectionPoint := func(n node, fieldList []string) {

		// Stop running and set isLast once we have stored all the fields
		if len(fieldList) == 0 {
			// node.setLast(input)
			fmt.Println("set last field here")
			return
		}

		//see if current node has children that match the field we're looking at. ex "A" or "*" or "foo"
		if _, ok := n.next[fieldList[0]]; !ok {
			// fmt.Println("field was not found in the map", val)
			fmt.Println("Current Node", n)
			next := map[string]*node{}
			next[fieldList[0]] = 
			n.next = next
			// n.next["test"] = &node{
			// 	data: fieldList[0],
			// }
			fmt.Println("NODE NEXT ==", n.next)
			// node.next[fieldList[0]] = *newNode("sup")
			//===================
			//HERE
			//===================
		}

		fmt.Println("add this fieldList", fieldList)
	}

	findInjectionPoint(p.root, tokenized)

	fmt.Println("p ::", p)
}
