package todo

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)

const (
	PriorityNone Priority = iota
	PriorityLow
	PriorityMedium
	PriorityHigh
)

type (
	Priority uint8

	DueDate struct {
		Time time.Time
	}

	Item struct {
		Id       string   `json:"id"`
		Content  string   `json:"content"`
		DueDate  *DueDate `json:"due_date,omitempty"`
		Priority Priority `json:"priority"`
		Done     bool     `json:"done"`
	}

	List map[string]*Item
)

func (p Priority) String() string {
	switch p {
	case PriorityNone:
		return "None"
	case PriorityLow:
		return "Low"
	case PriorityMedium:
		return "Medium"
	case PriorityHigh:
		return "High"
	}
	return "<UNKNOWN>"
}

func (d *DueDate) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = s[1 : len(s)-1]
	if t, err := time.Parse(time.DateOnly, s); err != nil {
		return err
	} else {
		d.Time = t
	}
	return nil
}

func (d *DueDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format(time.DateOnly))
}

func NewItem(content string) *Item {
	return &Item{Id: uuid.New().String(), Content: content}
}

func (t *Item) WithDueDate(dueDate *DueDate) *Item {
	t.DueDate = dueDate
	return t
}

func (t *Item) WithPriority(priority Priority) *Item {
	t.Priority = priority
	return t
}

func (t *Item) String() string {
	var sb strings.Builder
	var doneRune rune

	if t.Done == true {
		doneRune = 'x'
	} else {
		doneRune = ' '
	}

	sb.WriteString(fmt.Sprintf("[%c] %s", doneRune, t.Content))
	if t.Priority != PriorityNone {
		sb.WriteString(" [" + t.Priority.String() + "]")
	}
	if t.DueDate != nil {
		sb.WriteString(t.DueDate.Time.Format(time.DateOnly))
	}

	return sb.String()
}

func NewListFromJson(jsonData []byte) (*List, error) {
	var items []*Item
	list := make(List)

	if err := json.Unmarshal(jsonData, &items); err != nil {
		return nil, err
	}
	for _, item := range items {
		list[item.Id] = item
	}
	return &list, nil
}

func (l *List) ToJson() ([]byte, error) {
	items := make([]*Item, 0, len(*l))

	for _, v := range *l {
		items = append(items, v)
	}

	return json.Marshal(items)
}

func (l *List) Print() {
	for _, item := range *l {
		fmt.Println(item.String())
	}
}
