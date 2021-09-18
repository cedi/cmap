package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
	"sort"

	kout "github.com/cedi/kkpctl/pkg/output"
	"github.com/lensesio/tableprinter"
	"gopkg.in/yaml.v2"
)

// HostShort
type HostShort struct {
	Name string `header:"name"`
	IP   string `header:"IP"`
	SSH  bool   `header:"SSH Port open"`
}

func (r HostShort) ParseCollection(inputObj interface{}, output string, sortBy string) (string, error) {
	var err error
	var parsedOutput []byte

	objects, ok := inputObj.([]HostShort)
	if !ok {
		return "", fmt.Errorf("inputObj is not a []HostShort")
	}

	switch output {
	case kout.JSON:
		parsedOutput, err = json.MarshalIndent(objects, "", "  ")

	case kout.YAML:
		parsedOutput, err = yaml.Marshal(objects)

	case kout.Text:
		sort.Slice(objects, func(i, j int) bool {
			return objects[j].Name > objects[i].Name
		})

		var bodyBuf io.ReadWriter
		bodyBuf = new(bytes.Buffer)

		tableprinter.Print(bodyBuf, objects)
		parsedOutput, err = ioutil.ReadAll(bodyBuf)
	}

	return string(parsedOutput), err
}

func init() {
	parser := kout.GetParserFactory()
	parser.AddCollectionParser(reflect.TypeOf([]HostShort{}), HostShort{})
}
