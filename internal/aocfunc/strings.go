package aocfunc

type Strings []string

type StringFunction func(string) string
type StringConsumer func(string)
type StringPredicate func(string) bool

func (s Strings) ForEach(fn StringConsumer) {
	for _, el := range s {
		fn(el)
	}
}

func (s Strings) FilterString(predicate StringPredicate) (res Strings) {
	for _, el := range s {
		if predicate(el) {
			res = append(res, el)
		}
	}
	return
}

func (s Strings) MapString(fn StringFunction) (res Strings) {
	for _, el := range s {
		el = fn(el)
		res = append(res, el)
	}
	return
}
