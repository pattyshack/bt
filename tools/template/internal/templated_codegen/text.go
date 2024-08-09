// Auto-generated from source: text.template

package templated_codegen

import (
	_fmt "fmt"
	_io "io"

	"github.com/pattyshack/gt/tools/template/internal"
)

type Text struct {
	ind  string
	stmt *template.Atom
}

func (Text) Name() string { return "Text" }

func (template *Text) writeValue(
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

func (_template *Text) WriteTo(
	_output _io.Writer) (
	int64,
	error) {

	_numWritten := int64(0)

	ind := _template.ind
	stmt := _template.stmt

	// text.template:14:0
	if stmt.Value != "" {
		// text.template:15:0
		{
			_n, _err := _template.writeValue(
				_output,
				(ind),
				"text.template:15:0")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// text.template:15:4
		{
			_n, _err := _output.Write([]byte(`// `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// text.template:15:7
		{
			_n, _err := _template.writeValue(
				_output,
				(stmt.Loc()),
				"text.template:15:7")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// text.template:15:20
		{
			_n, _err := _output.Write([]byte(`
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// text.template:16:0
		{
			_n, _err := _template.writeValue(
				_output,
				(ind),
				"text.template:16:0")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// text.template:16:4
		{
			_n, _err := _output.Write([]byte(`{
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// text.template:17:0
		{
			_n, _err := _template.writeValue(
				_output,
				(ind),
				"text.template:17:0")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// text.template:17:4
		{
			_n, _err := _output.Write([]byte(`	_n, _err := _output.Write([]byte(` + "`" + ``))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// text.template:17:39
		{
			_n, _err := _template.writeValue(
				_output,
				(stmt.Value),
				"text.template:17:39")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// text.template:17:52
		{
			_n, _err := _output.Write([]byte(`` + "`" + `))
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// text.template:18:0
		{
			_n, _err := _template.writeValue(
				_output,
				(ind),
				"text.template:18:0")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// text.template:18:4
		{
			_n, _err := _output.Write([]byte(`	_numWritten += int64(_n)
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// text.template:19:0
		{
			_n, _err := _template.writeValue(
				_output,
				(ind),
				"text.template:19:0")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// text.template:19:4
		{
			_n, _err := _output.Write([]byte(`	if _err != nil {
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// text.template:20:0
		{
			_n, _err := _template.writeValue(
				_output,
				(ind),
				"text.template:20:0")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// text.template:20:4
		{
			_n, _err := _output.Write([]byte(`		return _numWritten, _err
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// text.template:21:0
		{
			_n, _err := _template.writeValue(
				_output,
				(ind),
				"text.template:21:0")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// text.template:21:4
		{
			_n, _err := _output.Write([]byte(`	}
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// text.template:22:0
		{
			_n, _err := _template.writeValue(
				_output,
				(ind),
				"text.template:22:0")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// text.template:22:4
		{
			_n, _err := _output.Write([]byte(`}
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}

	return _numWritten, nil
}
