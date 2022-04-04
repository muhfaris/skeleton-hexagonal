package schema

import (
	"bytes"
)

// $ cd transport/http/graphql
// $ go-bindata -ignore=\.go -pkg=schema -o=schema/bindata.go schema/...

// LoadGraphqlSchemas is Read all graphql type and then to string
func LoadGraphqlSchemas() string {
	buf := bytes.Buffer{}
	for _, name := range AssetNames() {
		b := MustAsset(name)
		buf.Write(b)

		if len(b) > 0 && b[len(b)-1] != '\n' {
			buf.WriteByte('\n')
		}
	}

	return buf.String()
}
