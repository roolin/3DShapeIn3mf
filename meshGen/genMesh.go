package meshGen

import (
	"fmt"
)

func GenUzi() ([]Vertex, []Triangle) {
	var r float64 = 30
	var segments int = 50
	var height float64 = 60

	//TODO: koło na spodzie
	//TODO: zamknięcie góry
	//TODO: czy jeden krok daje taki sam wynik jak interacja względem z?

	index := 0
	vs, ts := getCircle(r, segments, Vertex{0, 0, 0, index}, false)
	index += len(vs)

	vs1 := PointsForStadium(r*2, r*2, Vertex{0, 0, 0, index}, 50)
	index += len(vs1)

	vs2 := PointsForStadium(r*2-height, r*2, Vertex{0, 0, height, index}, 50)
	index += len(vs2)

	ts = append(ts, connectStadiums(vs1, vs2)...)

	vs = append(vs, vs1...)
	vs = append(vs, vs2...)

	for a, b := range vs {
		fmt.Printf("%d, %f\n", a, b)
	}

	for a, b := range ts {
		fmt.Printf("%d, %v\n", a, b)
	}

	return vs, ts
}

func connectStadiums(bottom []Vertex, top []Vertex) []Triangle {
	var ts []Triangle
	lenght := len(bottom)
	fmt.Printf("LEN: %d\n", lenght)
	for i := 0; i < lenght; i += 1 {
		if i == lenght-1 { // special case for closing cirle
			fmt.Printf("SPECIAL: %d\n", i)
			ts = append(ts, Triangle{bottom[i].id, bottom[0].id, top[i].id})
			ts = append(ts, Triangle{top[i].id, bottom[0].id, top[0].id})
		} else {
			fmt.Printf("NORMAL: %d\n", i)
			ts = append(ts, Triangle{bottom[i].id, bottom[i+1].id, top[i].id})
			ts = append(ts, Triangle{top[i].id, bottom[i+1].id, top[i+1].id})
		}
	}

	return ts
}

func PointsForStadium(width float64, lenght float64, center Vertex, segments int) []Vertex {
	radius := width / 2
	straightLen := lenght - width

	pointsForCicle := PointsForCircle(radius, segments, center)[1:]

	rightPoints := MovePointsHorizontal(pointsForCicle[0:len(pointsForCicle)/2+1], straightLen/2, pointsForCicle[0].id-1)

	leftPoints := MovePointsHorizontal(append(pointsForCicle[len(pointsForCicle)/2:], pointsForCicle[0]), -straightLen/2, pointsForCicle[0].id+len(rightPoints)-1)

	return append(rightPoints, leftPoints...)
}

func MovePointsHorizontal(points []Vertex, shift float64, startId int) []Vertex {
	var newPoints []Vertex
	for _, p := range points {
		newPoints = append(newPoints, Vertex{p.X + shift, p.Y, p.Z, startId})
		startId++
	}
	return newPoints
}

func Cylinder(radius float64, segments int) ([]Vertex, []Triangle) {
	vIndex := 0

	vs, ts := getCircle(radius, segments, Vertex{0, 0, 0, vIndex}, false)
	vIndex += len(vs)

	vs2, ts2 := getCircle(radius, segments, Vertex{0, 0, 60, vIndex}, true)
	vs = append(vs, vs2...)
	ts = append(ts, ts2...)

	walls := connectCirles(vs, 0, vIndex)
	ts = append(ts, walls...)

	return vs, ts
}
