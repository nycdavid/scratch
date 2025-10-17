package main

import (
	"strconv"
)

type (
	Event struct {
		Kind     string // "acquire" or "release"
		Thread   string // T0, T1,...,Tn
		Resource string // gpu0, gpu1, ctx, etc.
	}

	Detector struct {
		graph       [][]int
		resourceIdx map[string]string
	}

	DetectorMod func(*Detector)
)

// returns true when deadlocked
func (d *Detector) Process(e *Event) ([]string, bool) {
	tnum := threadInt(e.Thread)
	res := e.Resource
	resUser := d.resourceIdx[res]

	if e.Kind == "acquire" {
		if resUser == "" {
			// resource is available
			d.resourceIdx[res] = e.Thread
		}

		if resUser != "" {
			resUserNum := threadInt(resUser)
			// resource is unavailable, wait on user
			d.graph[tnum] = append(d.graph[tnum], resUserNum)
		}
	} else if e.Kind == "release" {
		if resUser == "" {
			// resource wasn't busy, do nothing
		}

		if resUser != "" {
			// free resource
			d.resourceIdx[res] = ""
		}
	}

	return nil, false
}

func WithGraph(graph [][]int) DetectorMod {
	return func(d *Detector) {
		d.graph = graph
	}
}

func WithResourceIdx(resourceIdx map[string]string) DetectorMod {
	return func(d *Detector) {
		d.resourceIdx = resourceIdx
	}
}

func NewDetector(opts ...DetectorMod) *Detector {
	d := new(Detector)

	for _, opt := range opts {
		opt(d)
	}

	if d.graph == nil {
		d.graph = make([][]int, 0)
	}

	if d.resourceIdx == nil {
		d.resourceIdx = make(map[string]string)
	}

	return d
}

func threadInt(thread string) int {
	strInt := thread[1:len(thread)]
	i, err := strconv.Atoi(strInt)
	if err != nil {
		panic(err)
	}

	return i
}

func main() {
}
