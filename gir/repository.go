package gir

import (
	"encoding/xml"
	"os"

	"github.com/pkg/errors"
)

type Repository struct {
	Includes   []Include   `xml:"http://www.gtk.org/introspection/core/1.0 include"`
	CIncludes  []CInclude  `xml:"http://www.gtk.org/introspection/c/1.0 include"`
	Namespaces []Namespace `xml:"http://www.gtk.org/introspection/core/1.0 namespace"`
}

func ParseRepositoryFile(path string) (*Repository, error) {
	f, err := os.Open("./Handy-1.gir")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open file")
	}
	defer f.Close()

	var r Repository

	if err := xml.NewDecoder(f).Decode(&r); err != nil {
		return nil, errors.Wrap(err, "Failed to decode gir XML")
	}

	return &r, nil
}
