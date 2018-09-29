package bug

import (
	"testing"
	"time"

	"github.com/go-test/deep"
)

func TestCreate(t *testing.T) {
	snapshot := Snapshot{}

	var rene = Person{
		Name:  "René Descartes",
		Email: "rene@descartes.fr",
	}

	unix := time.Now().Unix()

	create := NewCreateOp(rene, unix, "title", "message", nil)

	create.Apply(&snapshot)

	hash, err := create.Hash()
	if err != nil {
		t.Fatal(err)
	}

	comment := Comment{Author: rene, Message: "message", UnixTime: create.UnixTime}

	expected := Snapshot{
		Title: "title",
		Comments: []Comment{
			comment,
		},
		Author:    rene,
		CreatedAt: create.Time(),
		Timeline: []TimelineItem{
			NewCreateTimelineItem(hash, comment),
		},
	}

	deep.CompareUnexportedFields = true
	if diff := deep.Equal(snapshot, expected); diff != nil {
		t.Fatal(diff)
	}
}
