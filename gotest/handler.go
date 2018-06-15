package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprintf("Hello, Go, from Eric at the kuberneetes meetup. You said: %s", string(req))
}
