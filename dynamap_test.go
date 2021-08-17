package dynamap

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	var (
		err  error
		node interface{} = make(map[string]interface{})
	)

	node, err = Set(node, time.Second*3, "timeout")

	assert.NoError(t, err)
	assert.Equal(
		t,
		map[string]interface{}{
			"timeout": time.Second * 3,
		},
		node,
	)

	node, err = Set(node, 32, "clients", 1, "max_connections")

	assert.NoError(t, err)
	assert.Equal(
		t,
		map[string]interface{}{
			"timeout": time.Second * 3,
			"clients": []interface{}{
				nil,
				map[string]interface{}{
					"max_connections": 32,
				},
			},
		},
		node,
	)

	node, err = Set(node, "Content-Type", "clients", 0, "headers_blacklist", 0)

	assert.NoError(t, err)
	assert.Equal(
		t,
		map[string]interface{}{
			"timeout": time.Second * 3,
			"clients": []interface{}{
				map[string]interface{}{
					"headers_blacklist": []interface{}{
						"Content-Type",
					},
				},
				map[string]interface{}{
					"max_connections": 32,
				},
			},
		},
		node,
	)

	node, err = Set(node, "hello", "clients", 0, "headers_blacklist", "test")

	assert.Error(t, err)
}

func TestGet(t *testing.T) {
	var (
		err error
		m   = map[string]interface{}{
			"timeout": time.Second * 3,
			"clients": []interface{}{
				map[string]interface{}{
					"headers_blacklist": []interface{}{
						"Content-Type",
					},
				},
				map[string]interface{}{
					"max_connections": 32,
				},
			},
		}
	)

	v, err := Get(m, "clients", 1)

	assert.NoError(t, err)
	assert.Equal(
		t,
		map[string]interface{}{
			"max_connections": 32,
		},
		v,
	)

	v, err = Get(m, "clients", 0, "headers_blacklist", 0)

	assert.NoError(t, err)
	assert.Equal(
		t,
		"Content-Type",
		v,
	)

	v, err = Get(m, "does", "not", "exists")

	assert.NoError(t, err)
	assert.Nil(t, v)

	v, err = Get(m, "clients", 0, "headers_blacklist", 1)

	assert.NoError(t, err)
	assert.Nil(t, v)

	v, err = Get(m, "clients", "key")

	assert.Error(t, err)
}
