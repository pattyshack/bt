// Auto-generated from source: continue.template

package templated_codegen

import (
	_fmt "fmt"
	_io "io"

	"github.com/pattyshack/bt/tools/template/internal"
)

type Continue struct {
	ind  string
	stmt *template.Atom
}

func (Continue) Name() string { return "Continue" }

func (template *Continue) writeValue(
	output _io.Writer,
	value interface{},
	loc string) (
	int,
	error) {

	var valueBytes []byte
	switch val := value.(type) {
	case _fmt.Stringer:
		valueBytes = []byte(val.String())
	case string:
		valueBytes = []byte(val)
	case []byte:
		valueBytes = val
	case bool:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case uint:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case uint8:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case uint16:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case uint32:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case uint64:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case int:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case int8:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case int16:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case int32:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case int64:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case float32:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case float64:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case complex64:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case complex128:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	default:
		return 0, _fmt.Errorf(
			"Unsupported output value type (%s): %v",
			loc,
			value)
	}

	return output.Write(valueBytes)
}

func (_template *Continue) WriteTo(
	_output _io.Writer) (
	int64,
	error) {

	_numWritten := int64(0)

	ind := _template.ind
	stmt := _template.stmt

	// continue.template:13:0
	{
		_n, _err := _template.writeValue(
			_output,
			(ind),
			"continue.template:13:0")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// continue.template:13:4
	{
		_n, _err := _output.Write([]byte(`// `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// continue.template:13:7
	{
		_n, _err := _template.writeValue(
			_output,
			(stmt.Loc()),
			"continue.template:13:7")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// continue.template:13:20
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// continue.template:14:0
	{
		_n, _err := _template.writeValue(
			_output,
			(ind),
			"continue.template:14:0")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// continue.template:14:6
	{
		_n, _err := _output.Write([]byte(`continue
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	return _numWritten, nil
}
