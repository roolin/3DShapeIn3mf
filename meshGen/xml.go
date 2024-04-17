package meshGen

import (
	"encoding/xml"
	"log"
)

type Vertex struct {
	X  float64 `xml:"x,attr"`
	Y  float64 `xml:"y,attr"`
	Z  float64 `xml:"z,attr"`
	id int
}

type Vertices struct {
	XMLName xml.Name `xml:"vertices"`
	Vertex  []Vertex `xml:"vertex"`
}

type Triangle struct {
	V1 int `xml:"v1,attr"`
	V2 int `xml:"v2,attr"`
	V3 int `xml:"v3,attr"`
}

type Triangles struct {
	XMLName  xml.Name   `xml:"triangles"`
	Triangle []Triangle `xml:"triangle"`
}

func getXML4Vertices(vs []Vertex) string {
	v := Vertices{Vertex: vs}
	out, err := xml.MarshalIndent(v, " ", "  ")
	if err != nil {
		log.Printf("error %v\n", err)
	}
	return string(out)
}

func getXML4Triangles(ts []Triangle) string {
	v := Triangles{Triangle: ts}
	out, err := xml.MarshalIndent(v, " ", "  ")
	if err != nil {
		log.Printf("error %v\n", err)
	}
	return string(out)
}

func Body(vs []Vertex, ts []Triangle) string {
	start := `<?xml version="1.0" encoding="UTF-8"?>
	<model unit="mm"
		   xml:lang="en-US"
		   xmlns="http://schemas.microsoft.com/3dmanufacturing/core/2015/02">
		<metadata name="Copyright">
			Copyright (c) 2015 3MF Consortium. All rights reserved.
		</metadata>
		<resources>
			<object id="1" type="model">
				<mesh>
`

	end := `
				</mesh>
			</object>
		</resources>
		<build>
			<item objectid="1" />
		</build>
	</model>
	`
	return start + getXML4Vertices(vs) + getXML4Triangles(ts) + end

}

func TriangleBody() string {
	return `<?xml version="1.0" encoding="UTF-8"?>
	<model unit="inch"
		   xml:lang="en-US"
		   xmlns="http://schemas.microsoft.com/3dmanufacturing/core/2015/02">
		<metadata name="Copyright">
			Copyright (c) 2015 3MF Consortium. All rights reserved.
		</metadata>
		<resources>
			<object id="1" type="model">
				<mesh>
					<vertices>
						  <vertex x="0" y="0" z="0" />
						  <vertex x="1" y="0" z="0" />
						  <vertex x="0" y="1" z="0" />
						  <vertex x="0" y="0" z="1" />
					</vertices>
					<triangles>
						  <triangle v1="0" v2="2" v3="1" />
						  <triangle v1="0" v2="1" v3="3" />
						  <triangle v1="0" v2="3" v3="2" />
						  <triangle v1="1" v2="2" v3="3" />
					</triangles>
				</mesh>
			</object>
		</resources>
		<build>
			<item objectid="1" />
		</build>
	</model>
	`
}

func CubeBody() string {
	return `<?xml version="1.0" encoding="UTF-8"?>
	<model unit="inch"
		   xml:lang="en-US"
		   xmlns="http://schemas.microsoft.com/3dmanufacturing/core/2015/02">
		<metadata name="Copyright">
			Copyright (c) 2015 3MF Consortium. All rights reserved.
		</metadata>
		<resources>
			<object id="1" type="model">
				<mesh>
					<vertices>
						  <vertex x="0" y="0" z="0" />
						  <vertex x="1" y="0" z="0" />
						  <vertex x="1" y="2" z="0" />
						  <vertex x="0" y="2" z="0" />
						  <vertex x="0" y="0" z="3" />
						  <vertex x="1" y="0" z="3" />
						  <vertex x="1" y="2" z="3" />
						  <vertex x="0" y="2" z="3" />
					</vertices>
					<triangles>
						  <triangle v1="3" v2="2" v3="1" />
						  <triangle v1="1" v2="0" v3="3" />
						  <triangle v1="4" v2="5" v3="6" />
						  <triangle v1="6" v2="7" v3="4" />
						  <triangle v1="0" v2="1" v3="5" />
						  <triangle v1="5" v2="4" v3="0" />
						  <triangle v1="1" v2="2" v3="6" />
						  <triangle v1="6" v2="5" v3="1" />
						  <triangle v1="2" v2="3" v3="7" />
						  <triangle v1="7" v2="6" v3="2" />
						  <triangle v1="3" v2="0" v3="4" />
						  <triangle v1="4" v2="7" v3="3" />
					</triangles>
				</mesh>
			</object>
		</resources>
		<build>
			<item objectid="1" />
		</build>
	</model>
	`
}
