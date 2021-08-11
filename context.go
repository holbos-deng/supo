package main

type context struct {
	values map[string]interface{}
}

func (c context) Get(k string) interface{} {
	return c.values[k]
}

func (c context) Set(k string, v interface{}) {
	c.values[k] = v
}

func (c context) Remove(k string) interface{} {
	var v interface{}
	var ok bool
	if v, ok = c.values[k]; ok {
		delete(c.values, k)
	}
	return v
}

func (c context) Has(k string) bool {
	_, b := c.values[k]
	return b
}
