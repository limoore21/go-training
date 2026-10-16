package main

import ("fmt"
	"errors")

type Cache struct{
	storage map[string]string
}

func NewCache() *Cache{
	a := Cache{
		make(map[string]string),
	}
	return &a
}

func Set(key, value string) error{
	if key == (""){
		return error.New("Ошибка:Пустой ключ недопустим")
	} else if value := c.storage[key]{
		return nil
	}
}

func Get(key string) (string, error){
	val, isFound := c.storage[key]
	if isFound ==false{
		return ""
		return errors.New ("ошибка: ключ не найден")
	} else {
		return val, nil
	}
}
            
		return nil
	}	
}

func main() {
	err := cache.Ser("", "123")
	if err != nil{
	fmt.Println(err)
	} else {
		"token" -> "super-secret"
	}
	if "token" 
}