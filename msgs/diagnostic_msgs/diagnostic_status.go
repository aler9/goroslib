// Autogenerated with msg-import, do not edit.
package diagnostic_msgs

import (
	"github.com/aler9/goroslib/msgs"
)

type DiagnosticStatus struct {
	Level      msgs.Byte
	Name       msgs.String
	Message    msgs.String
	HardwareId msgs.String
	Values     []KeyValue
}