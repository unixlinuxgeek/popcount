### Упражнение 2.4

Напишите версию PopCount, которая подсчитывает биты с помощью сдвига аргумента по всем 64 позициям,
проверяя при каждом сдвиге крайний справа бит. Сравните производительность этой версии с выборкой из таблицы.

Решение:

```
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/unixlinuxgeek/popcount
cpu: 12th Gen Intel(R) Core(TM) i5-12450H
BenchmarkPopCount-12            1000000000               0.0000002 ns/op
BenchmarkPopCount2-12           1000000000               0.0000003 ns/op
BenchmarkPopCount3-12           1000000000               0.0000002 ns/op
PASS
ok      github.com/unixlinuxgeek/popcount       0.007s
```

PopCount:  0.0000002 ns/op
PopCount2  0.0000003 ns/op
PopCount3  0.0000002 ns/op
