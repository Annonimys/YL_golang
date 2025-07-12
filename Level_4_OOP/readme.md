# Виртуальный зоопарк

**Ограничение времени:** 60 секунд  
**Ограничение памяти:** 1524 Мб  
**Ввод:** стандартный ввод или `main.go`  
**Вывод:** стандартный вывод  

## Условие

Реализуйте виртуальный зоопарк на Go с использованием принципов ООП.

### Требования:

1. **Интерфейс `Animal`**:
   ```go
   type Animal interface {
       MakeSound() string
       GetName() string
       GetInfo() string
   }
   ```
   - `MakeSound()` — возвращает звук животного
   - `GetName()` — возвращает имя животного
   - `GetInfo()` — возвращает информацию в формате:  
     `Имя: <name>, Вид: <species>, Возраст: <age>`

2. **Структура `animal`** (приватные поля):
   ```go
   type animal struct {
       name    string
       species string
       age     int
       sound   string
   }
   ```
   - Конструктор:  
     ```go
     func NewAnimal(name, species string, age int, sound string) Animal
     ```

3. **Функция `ZooShow`**:
   ```go
   func ZooShow(animals []Animal)
   ```
   - Для каждого животного выводит:  
     ```
     <результат GetInfo()>
     <результат MakeSound()>
     ```

4. **Структура `ZooKeeper`**:
   ```go
   type ZooKeeper struct{}
   ```
   - Метод:  
     ```go
     func (z ZooKeeper) Feed(animal Animal)
     ```
     Выводит:  
     `Смотритель зоопарка кормит <name>. <sound>!`

### Пример использования:

```go
func main() {
    animals := []Animal{
        NewAnimal("Барсик", "Тигр", 5, "Ррр"),
        NewAnimal("Кеша", "Попугай", 2, "Чирик"),
    }

    ZooShow(animals)

    keeper := ZooKeeper{}
    keeper.Feed(animals[0])
}
```

**Вывод:**
```
Имя: Барсик, Вид: Тигр, Возраст: 5
Ррр
Имя: Кеша, Вид: Попугай, Возраст: 2
Чирик
Смотритель зоопарка кормит Барсик. Ррр!
```

### Примечания:
- Все поля структур должны быть приватными (`animal`, а не `Animal`)
- Доступ к полям — только через методы интерфейса
- Для вывода информации используйте методы интерфейса
- Реализуйте полиморфизм через интерфейс `Animal`