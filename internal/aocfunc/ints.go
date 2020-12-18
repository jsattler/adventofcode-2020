package aocfunc

type Ints []int

type IntFunction func(int)int
type IntConsumer func(int)
type IntPredicate func(int)bool

func (i Ints) ForEach(fn IntConsumer){
	for _, el := range i {
		fn(el)
	}
}


func (i Ints) FilterInt(predicate IntPredicate) (res Ints) {
	for _, el := range i{
		if predicate(el){
			res = append(res, el)
		}
	}
	return
}

func (i Ints) MapInt(fn IntFunction) (res Ints){
	for _, el := range i {
		el = fn(el)
		res = append(res, el)
	}
	return
}
