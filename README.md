### Упражнение 2.5

Выражение ```x&(x-1)``` сбрасывает крайний справа ненулевой бит ```x```.
Напишите версию ```PopCount```, которая подсчитывает биты с использованием этого факта,
и оцените производительность.

Запускаеи модульные тесты: 

```
go test -v
```
Вывод:
```
=== RUN   TestPopCountNew1
popcount_test.go:13: TestPopCountNew1: Десятичное число 2050, количество установленных битов: 2 (100000000010)
--- PASS: TestPopCountNew1 (0.00s)
=== RUN   TestPopCountNew2
popcount_test.go:25: TestPopCountNew2: Десятичное число 3, количество установленных битов: 2 (11)
--- PASS: TestPopCountNew2 (0.00s)
=== RUN   TestPopCountNew3
popcount_test.go:37: TestPopCountNew3: Десятичное число 55, количество установленных битов: 5 (110111)
--- PASS: TestPopCountNew3 (0.00s)
PASS
ok  	github.com/unixlinuxgeek/popcount	0.002s
```


Перенаправляем вывод в файл(если необходимо):
```
go test -v > test_out.txt
```
