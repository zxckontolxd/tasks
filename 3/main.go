package main

// Я не дочитал все условия и принялся выполнять по порядку, так что некоторые моменты не по условию, но они лучше, чем в условии. Если есть в этом проблема, то могу исправить.
// Я сделал много упрощений, которые не позволительны в реальном коде, но подчеркнул их. Так что опять же, не думаю, что в этом есть проблема.

import (
	"fmt"
	"errors"
)

type Pair struct {
	key string
	value int
}

type StringIntMap struct {
	// по хорошему, еще нужно реализовать интерфейс, который бы позволял пробегаться через foreach
	// ну и String() тоже не помешал бы
	storage []Pair
}

func (mmap *StringIntMap) Add(key string, value int) {
	// по идее нужно еще добавить поиск "коллизий" и создавать бакеты
	mmap.storage = append(mmap.storage, Pair{key: key, value: value})
}

// чет я хз зачем это сделал, но пусть будет
// вместо Get() будет
func (mmap *StringIntMap) Find(value int) (string, error) {
    for i := range mmap.storage {
        if mmap.storage[i].value == value {
            return mmap.storage[i].key, nil
        }
    }
    return "", errors.New("Cannot find value")
}

func (mmap *StringIntMap) Remove(key string) error {
	for i := range mmap.storage {
        if mmap.storage[i].key == key {
            mmap.storage = append(mmap.storage[:i], mmap.storage[i+1:]...)
            return nil
        }
    }
    return errors.New("Cannot find key")
}

// я сейчас заметил, что достаточно было в качестве хранилища использовать map[string]int, что было бы гораздо проще
// не думаю, что это проблема, поэтому продолжу уже доделывать
// если необходимо, могу потом исправить
func (mmap *StringIntMap) Copy() StringIntMap {
	ret := StringIntMap{
		storage: make([]Pair, len(mmap.storage)),
	}
	copy(ret.storage, mmap.storage)
	return ret
}

func (mmap *StringIntMap) Exist(key string) bool {
	for i := range mmap.storage {
        if mmap.storage[i].key == key {
            return true
        }
    }
    return false
}

func main() {
	var mmap StringIntMap
	mmap.Add("str", 42)
	mmap.Add("te4st", 228)
	fmt.Println(mmap.storage)

	key, _ := mmap.Find(42)
	fmt.Println(key)

	mmap.Remove("str")
	fmt.Println(mmap.storage)

	smap := mmap.Copy()
	fmt.Println(smap.storage)

	//тут нужно было использовать printf, но мне лень гуглить, как вывести bool(( 
	fmt.Println(mmap.Exist("str"), ", ", mmap.Exist("te4st"))
}