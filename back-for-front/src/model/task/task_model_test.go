package task

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	FIRSTPOSITION  = 0
	SECONDPOSITION = 1
	THIRDPOSITION  = 2
	FOURTHPOSITION = 3
)

func TestStructProperties(t *testing.T) {
	task := Task{}
	elem := reflect.ValueOf(&task).Elem()
	elemType := elem.Type()

	t.Run("Should have an ID property and JSON tag", func(t *testing.T) {
		field := elemType.Field(FIRSTPOSITION)
		assert.Equal(t, "ID", field.Name)
		assert.Equal(t, reflect.StructTag(`json:"id"`), field.Tag)
	})

	t.Run("Should have a Title property and JSON tag", func(t *testing.T) {
		field := elemType.Field(SECONDPOSITION)
		assert.Equal(t, "Title", field.Name)
		assert.Equal(t, reflect.StructTag(`json:"title"`), field.Tag)
	})

	t.Run("Should have a Description property and JSON tag", func(t *testing.T) {
		field := elemType.Field(THIRDPOSITION)
		assert.Equal(t, "Description", field.Name)
		assert.Equal(t, reflect.StructTag(`json:"description"`), field.Tag)
	})

	t.Run("Should have a Status property and JSON tag", func(t *testing.T) {
		field := elemType.Field(FOURTHPOSITION)
		assert.Equal(t, "Status", field.Name)
		assert.Equal(t, reflect.StructTag(`json:"status"`), field.Tag)
	})
}
