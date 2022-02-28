package breadcrumb

import (
	"os"
	"strings"
)

type Segment struct {
	Name string
	Path string
}

type Breadcrumb struct {
	Segments []Segment
}

func NewBreadcrumb(dir string) Breadcrumb {
	segments := make([]Segment, 0)
	s := strings.Split(dir, string(os.PathSeparator))

	segments = append(segments, Segment{
		Name: "/",
		Path: "/",
	})

	var way string
	for _, segment := range s {
		if segment == "" {
			continue
		}

		way += "/" + segment
		segments = append(segments, Segment{
			Name: segment,
			Path: way,
		})
	}

	return Breadcrumb{
		Segments: segments,
	}
}
