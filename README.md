### Упражнение 2.4

Напишите версию PopCount, которая подсчитывает биты с помощью сдвига аргумента по всем 64 позициям,
проверяя при каждом сдвиге крайний справа бит. Сравните производительность этой версии с выборкой из таблицы.


### Сравниваеи выборку из таблицы (PopCountOld) 
```
$ go test -bench=Old
goos: linux
goarch: amd64
pkg: github.com/unixlinuxgeek/popcount
cpu: 12th Gen Intel(R) Core(TM) i5-12450H
BenchmarkPopCountOld-12         1000000000               0.0000004 ns/op
PASS
ok      github.com/unixlinuxgeek/popcount       0.005s
```
Результат 0.005s

### Сравниваеи с версией с проверкой крайне правого бита (PopCountNew): 
```
$ go test -bench=New
goos: linux
goarch: amd64
pkg: github.com/unixlinuxgeek/popcount
cpu: 12th Gen Intel(R) Core(TM) i5-12450H
BenchmarkPopCountNew-12         1000000000               0.0000001 ns/op
PASS
ok      github.com/unixlinuxgeek/popcount       0.004s
```
Результат 0.004s


### Модульный тест
```
$ go test -v
=== RUN   TestPopCount
    popcount_test.go:15: Десятичное число 2050, количество установленных битов: 2 (100000000010)
--- PASS: TestPopCount (0.00s)
PASS
ok      github.com/unixlinuxgeek/popcount       0.002s
```