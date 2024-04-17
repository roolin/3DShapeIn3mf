package meshGen

import (
	"fmt"
	"testing"
)

func TestXml(test *testing.T) {
	vs := []Vertex{{0, 1, 2, 0}, {1, 2, 3, 1}}
	out := getXML4Vertices(vs)
	fmt.Println(out)

	ts := []Triangle{{2, 3, 4}, {4, 5, 6}}
	out = getXML4Triangles(ts)
	fmt.Println(out)

	out = Body(vs, ts)
	fmt.Println(out)

	test.Error("aaa")
}
