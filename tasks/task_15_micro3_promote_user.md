# Микро-задача 3: «Логический детектив» (PromoteUser)

## Тема
Срезы структур, `for … range` и копирование, изменение элемента по индексу.

## Условие
Написать функцию `PromoteUser(users []User, name string) bool`, которая:
1. Ищет в срезе пользователя с именем `name`.
2. Если находит — устанавливает `IsAdmin = true` **в оригинальном элементе среза**.
3. Возвращает `true`, если нашёл и повысил; `false` — если не нашёл.

```go
type User struct {
    Name    string
    IsAdmin bool
}
```

## Типичная ошибка
```go
// ❌ user — копия, изменение не сохранится
for _, user := range users {
    if user.Name == name {
        user.IsAdmin = true
    }
}
```

## Правильное решение
```go
// ✅ обращаемся к элементу по индексу — меняем оригинал
func PromoteUser(users []User, name string) bool {
    for i := range users {
        if users[i].Name == name {
            users[i].IsAdmin = true
            return true
        }
    }
    return false
}
```

## Ключевой принцип
При `for _, v := range slice` переменная `v` — **копия** элемента.
Чтобы изменить оригинал: используй индекс `slice[i]` или храни указатели `[]*User`.
