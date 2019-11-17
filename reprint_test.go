package reprint

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var addressSearch = regexp.MustCompile(`0x[0-9a-f]+`).Find

func getAddressOf(obj interface{}) string {
	s := fmt.Sprintf("%p", obj)
	b := addressSearch([]byte(s))
	if b == nil {
		return getAddressOf(&obj)
	}
	return string(b)
}

func assertAddressesAreDifferent(t *testing.T, obj1, obj2 interface{}) {
	ptr1 := getAddressOf(obj1)
	ptr2 := getAddressOf(obj2)
	if ptr1 == ptr2 {
		t.Errorf("pointers of %#v and %#v are equal: %s and %s", obj1, obj2, ptr1, ptr2)
	}
}

// Pretty much an integration test
func Test_deepCopy(t *testing.T) {
	t.Parallel()
	type customStruct struct {
		a int
		B *int
		C **int
		D []int
		E map[int]*int
		F *struct {
			Z *[]*int
		}
	}
	one := 1
	onePtr := &one
	onePtrPtr := &onePtr
	original := customStruct{
		a: one,
		B: onePtr,
		C: onePtrPtr,
		D: []int{one, one},
		E: map[int]*int{
			0: onePtr,
		},
		F: &struct{ Z *[]*int }{
			&[]*int{onePtr},
		},
	}
	originalValue := reflect.ValueOf(original)

	// Copies original
	copyValue := deepCopy(originalValue)
	require.True(t, copyValue.CanInterface())
	copy, ok := copyValue.Interface().(customStruct)
	require.True(t, ok)
	require.Equal(t, original, copy)

	// Change copy's value entirely
	copy.a = 2
	*copy.B = 2
	**copy.C = 2
	for i := range copy.D {
		copy.D[i] = 2
	}
	for k := range copy.E {
		*copy.E[k] = 2
	}
	*copy.F.Z = nil

	// Verify original is not equal to copy at all
	assert.NotEqual(t, copy.a, original.a)
	assert.NotEqual(t, copy.B, original.B)
	assert.NotEqual(t, copy.C, original.C)
	assert.NotEqual(t, copy.D, original.D)
	assert.NotEqual(t, copy.E, original.E)
	assert.NotEqual(t, copy.F, original.F)
}

func Test_This(t *testing.T) {
	t.Parallel()
	one := 1
	type customType struct{ A *int }
	original := customType{&one}
	copy := This(original)
	copyTyped, ok := copy.(customType)
	require.True(t, ok)
	require.Equal(t, original, copyTyped)
	*copyTyped.A = 2
	assert.NotEqual(t, copyTyped, original)
}

func Test_FromTo(t *testing.T) {
	t.Parallel()
	one := 1
	type customType struct{ A *int }
	t.Run("pointer to pointer", func(t *testing.T) {
		t.Parallel()
		original := &customType{&one}
		copy := new(customType)
		err := FromTo(original, copy)
		require.NoError(t, err)
		require.NotEmpty(t, copy)
		require.Equal(t, original, copy)
		*copy.A = 2
		assert.NotEqual(t, copy, original)
	})
	t.Run("value to pointer", func(t *testing.T) {
		t.Parallel()
		original := customType{&one}
		copy := new(customType)
		err := FromTo(original, copy)
		require.NoError(t, err)
		require.NotEmpty(t, copy)
		require.Equal(t, &original, copy)
		*copy.A = 2
		assert.NotEqual(t, copy, &original)
	})
	t.Run("nil to pointer", func(t *testing.T) {
		t.Parallel()
		copy := new(customType)
		FromTo(nil, copy)
		require.Empty(t, copy)
	})
	t.Run("pointer to nil", func(t *testing.T) {
		t.Parallel()
		original := customType{&one}
		err := FromTo(original, nil)
		require.Error(t, err)
		require.Equal(t, err.Error(), "FromTo: copy target is nil, it should be a valid pointer")
	})
	// t.Run("nil pointer to pointer", func(t *testing.T) {
	// 	t.Parallel()
	// 	var original *customType
	// 	copy := new(customType)
	// 	FromTo(original, copy)
	// 	require.Empty(t, copy)
	// 	// TODO require.Equal(t, original, copy)
	// })
	// t.Run("pointer to nil pointer", func(t *testing.T) {
	// 	t.Parallel()
	// 	original := &customType{&one}
	// 	var copy *customType
	// 	FromTo(original, copy)
	// })
}

func Test_forceCopyValue(t *testing.T) {
	t.Parallel()
	one := 1
	tests := map[string]interface{}{
		"integer":            one,
		"pointer to integer": &one,
		"string":             "a",
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc := tc
			t.Parallel()
			original := reflect.ValueOf(tc)
			copy := forceCopyValue(original)
			require.True(t, copy.CanInterface())
			copyInterface := copy.Interface()
			if &original == &copy {
				t.Errorf("pointers %s and %s are the same", &original, &copy)
			}
			assert.Equal(t, tc, copyInterface)
		})
	}
}

func Test_deepCopySlice(t *testing.T) {
	t.Parallel()
	// empty slice pointer does not change
	// but that is ok as appending would create
	// another slice
	tests := map[string]interface{}{
		"slice with nil":    []*int{nil},
		"slice of integers": []int{1, 2, 3},
		"slice of pointers": []int{1, 2, 3},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc := tc
			t.Parallel()
			original := reflect.ValueOf(tc)
			copy := deepCopySlice(original)
			require.True(t, copy.CanInterface())
			copyInterface := copy.Interface()
			assert.Equal(t, tc, copyInterface)
			assertAddressesAreDifferent(t, tc, copyInterface)
		})
	}
}

func Test_deepCopyMap(t *testing.T) {
	t.Parallel()
	tests := map[string]interface{}{
		"empty map":       map[int]int{},
		"map of integers": map[int]int{1: 1, 2: 2, 3: 3},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc := tc
			t.Parallel()
			original := reflect.ValueOf(tc)
			copy := deepCopyMap(original)
			require.True(t, copy.CanInterface())
			copyInterface := copy.Interface()
			assertAddressesAreDifferent(t, tc, copyInterface)
			assert.Equal(t, tc, copyInterface)
		})
	}
}

func Test_deepCopyPointer(t *testing.T) {
	t.Parallel()
	var nilPtr *int = nil
	one := 1
	onePtr := &one
	tests := map[string]interface{}{
		"nil int pointer":    nilPtr,
		"int pointer":        onePtr,
		"int nested pointer": &onePtr,
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc := tc
			t.Parallel()
			original := reflect.ValueOf(tc)
			copy := deepCopyPointer(original)
			require.True(t, copy.CanInterface())
			copyInterface := copy.Interface()
			if tc != nilPtr {
				assertAddressesAreDifferent(t, tc, copyInterface)
			}
			assert.Equal(t, tc, copyInterface)
		})
	}
}

func Test_deepCopyStruct(t *testing.T) {
	t.Parallel()
	tests := map[string]interface{}{
		"empty struct":                   struct{}{},
		"simple struct":                  struct{ A int }{1},
		"simple struct unexported field": struct{ a int }{1},
		"empty struct of pointers of structs": struct {
			A *struct{ A int }
			B *struct{ B int }
		}{},
		"struct of pointers of structs": struct {
			A *struct{ A int }
			B *struct{ B int }
		}{&struct{ A int }{1}, &struct{ B int }{2}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc := tc
			t.Parallel()
			original := reflect.ValueOf(tc)
			copy := deepCopyStruct(original)
			require.True(t, copy.CanInterface())
			copyInterface := copy.Interface()
			assert.Equal(t, tc, copyInterface)
			assertAddressesAreDifferent(t, tc, copyInterface)
		})
	}
}

func Test_deepCopyFunc(t *testing.T) {
	t.Parallel()
	f := func(a int) int { return a + 1 }
	original := reflect.ValueOf(f)
	copy := deepCopy(original)
	require.True(t, copy.CanInterface())
	copyInterface := copy.Interface()
	g, ok := copyInterface.(func(a int) int)
	require.True(t, ok)
	assert.Equal(t, f(0), g(0))
	g = func(a int) int { return a + 2 }
	assert.NotEqual(t, f(0), g(0))
}
