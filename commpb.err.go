package commpb

import "fmt"

func (v *Status) Error() string {
	return fmt.Sprintf("%+v", *v)
}
