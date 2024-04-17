package meshGen

import (
	"math"
)

func XYShift(radius float64, angle float64) (x, y float64) {
	counterAngle := counterAngle(angle)
	segment := segmentLength(radius, angle)

	xShift := math.Sin(degreeToRadian(counterAngle)) * segment
	yShift := math.Cos(degreeToRadian(counterAngle)) * segment

	return xShift, yShift
}

func segmentLength(radius float64, angle float64) float64 {
	x := math.Sin(degreeToRadian(angle / 2))
	halfLen := x * radius
	return halfLen * 2
}

func degreeToRadian(angle float64) float64 {
	return angle * (math.Pi / 180)
}

func counterAngle(angle float64) float64 {
	return 180 - 90 - angle/2
}

func connectCirles(points []Vertex, index1 int, index2 int) []Triangle {
	var ts []Triangle
	for i := 1; i < index2; i += 1 {
		if i == index2-1 { // special case for closing cirle
			ts = append(ts, Triangle{i, 1, i + index2})
			ts = append(ts, Triangle{i + index2, 1, index2 + 1})
		} else {
			ts = append(ts, Triangle{i, i + 1, i + index2})
			ts = append(ts, Triangle{i + index2, i + 1, i + index2 + 1})
		}
	}

	return ts
}

func getCircle(radius float64, segments int, centerPoint Vertex, reverse bool) ([]Vertex, []Triangle) {
	points := PointsForCircle(radius, segments, centerPoint)
	triangles := CirleMeshFromPoints(points, reverse)

	return points, triangles
}

func CirleMeshFromPoints(points []Vertex, reverse bool) []Triangle {
	ts := []Triangle{}

	center := points[0]
	for i := 1; i < len(points); i++ {
		point1 := points[i]
		var point2 Vertex
		if i == len(points)-1 { //if it's last, that cirle back
			point2 = points[1]
		} else {
			point2 = points[i+1]
		}

		if !reverse {
			ts = append(ts, Triangle{center.id, point2.id, point1.id})
		} else {
			ts = append(ts, Triangle{center.id, point1.id, point2.id})
		}
	}

	//ts = append(ts, Triangle{center.id, points[0].id, lastPoint.id})

	return ts
}

func PointsForCircle(radius float64, segments int, center Vertex) []Vertex {
	var startX float64 = 0
	var startY float64 = radius
	var segment float64 = float64(360) / float64(segments)
	var angle float64 = 0

	var ys []Vertex = make([]Vertex, segments+1)
	ys[0] = center

	indexStart := center.id + 1

	// TODO: can be optimized, can be run for 1 quarter and mapped for rest from that
	for i := 1; i <= segments; i++ {
		x, y := XYShift(radius, angle)
		ys[i] = Vertex{X: startX + x, Y: startY - y, Z: center.Z, id: indexStart}
		angle += segment
		indexStart++
	}

	return ys
}
