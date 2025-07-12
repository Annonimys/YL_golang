package main

import "fmt"

type Animal interface {
	MakeSound() string //возвращает строку со звуком животного
	GetName() string   //возвращает строку с именем животного
	GetInfo() string   //возвращает строку с информацией о животном в формате: Имя: *name*, Вид: *species*, Возраст: *age*
}

type animal struct {
	name    string
	species string
	age     int
	sound   string
}

func NewAnimal(name, species string, age int, sound string) Animal {
	return animal{
		name:    name,
		species: species,
		age:     age,
		sound:   sound,
	}
}

func (a animal) MakeSound() string {
	return a.sound
}

func (a animal) GetName() string {
	return a.name
}

func (a animal) GetInfo() string {
	return fmt.Sprintf("Имя: %s, Вид: %s, Возраст: %d", a.name, a.species, a.age)
}

func ZooShow(animals []Animal) {
	for _, animal := range animals {
		fmt.Println(animal.GetInfo())
		fmt.Println(animal.MakeSound())
	}
}

type ZooKeeper struct{}

func (z ZooKeeper) Feed(animal Animal) {
	fmt.Printf("Смотритель зоопарка кормит %s. %s!", animal.GetName(), animal.MakeSound())
}
