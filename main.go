package main

import (
	"archive/zip"
	"compress/flate"
	"io"
	"log"
	"main/meshGen"
	"os"
)

func main() {
	vs, ts := meshGen.GenUzi()

	//vs, ts := meshGen.GenMesh()
	create3mf((meshGen.Body(vs, ts)))
	//create3mf(meshGen.TriangleBody())
	//create3mf(meshGen.CubeBody())

	//f3d -e --resolution 2000,1000 archive.3mf
}

func create3mf(modelFile string) {
	archive, err := os.Create("archive.3mf")
	if err != nil {
		panic(err)
	}
	defer archive.Close()

	w := zip.NewWriter(archive)

	w.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(out, flate.BestCompression)
	})

	var files = []struct {
		Name, Body string
	}{
		{"3D/3dmodel.model", modelFile},
		{"_rels/.rels", relsBody()},
	}

	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	err = w.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func relsBody() string {
	return `<?xml version="1.0" encoding="utf-8"?>
	<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">
		<Relationship Type="http://schemas.microsoft.com/3dmanufacturing/2013/01/3dmodel" Target="/3D/3dmodel.model" Id="rel0"/>
	</Relationships>
	`
}
