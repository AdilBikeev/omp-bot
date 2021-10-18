package work

import (
	"encoding/json"
	"fmt"
)

const prettyPrintFormat = `{
	id: %d
	name: %s
	description: %s
	size: %d
	created_at: %s
}`

type Team struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Size        uint64 `json:"size"`
	CreatedAt   string `json:"created_at"`
}

func (t *Team) String() string {
	str, _ := json.Marshal(t)
	return fmt.Sprintf(`%s`, string(str))
}

func (t *Team) PrettyPrint() string {
	return fmt.Sprintf(
		prettyPrintFormat,
		t.ID,
		t.Name,
		t.Description,
		t.Size,
		t.CreatedAt,
	)
}
