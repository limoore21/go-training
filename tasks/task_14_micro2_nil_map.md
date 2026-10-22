# Микро-задача 2: «Блуждающий счётчик» (Maps & Nil)

## Тема
Nil map, `make`, паника при записи в неинициализированную мапу.

## Условие
Что произойдёт при запуске?

```go
type Stats struct {
    counters map[string]int
}

func (s *Stats) Increment(metric string) {
    s.counters[metric]++
}

func main() {
    var st Stats
    st.Increment("cpu_spike")
    fmt.Println(st.counters["cpu_spike"])
}
```

Варианты:
1. Выведет `1`
2. Выведет `0`
3. Упадёт в панику

## Ответ
3 — **паника**: `assignment to entry in nil map`

## Объяснение
- `var st Stats` создаёт структуру с нулевыми значениями. Поле `counters` = `nil`.
- Читать из nil map **можно** (вернёт дефолтное значение `0`).
- Писать в nil map **нельзя** — рантайм бросает панику.
- Исправление: инициализировать через `make` в конструкторе:
  ```go
  func NewStats() *Stats {
      return &Stats{counters: make(map[string]int)}
  }
  ```
