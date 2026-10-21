# Задача 12: Security Manager (Blacklist)

## Тема
Структуры, методы с указателем-получателем (*receiver), мапы, множественный возврат значений (multiple return), обработка ошибок (errors.New).

## Условие
1. Создать структуру `SecurityManager` с полем `blacklist` (мапа: ключ `string` — IP, значение `string` — причина блокировки).
2. Написать конструктор `NewSecurity() *SecurityManager`.
3. Написать метод `(s *SecurityManager) BlockIP(ip string, reason string) error`:
   - Если `ip` пустой, вернуть ошибку: `"IP-адрес не может быть пустым"`.
   - Если `reason` пустой, вернуть ошибку: `"причина блокировки обязательна"`.
   - Иначе сохранить запись в мапу `s.blacklist[ip] = reason` и вернуть `nil`.
4. Написать метод `(s *SecurityManager) CheckAccess(ip string) (bool, string)`:
   - Проверить наличие `ip` в мапе через синтаксис `val, exists := s.blacklist[ip]`.
   - Если найден: вернуть `false, val` (доступ запрещен, отдаем причину).
   - Если не найден: вернуть `true, ""` (доступ разрешен).
5. В `main` протестировать добавление, валидацию и проверку доступа.
