## Dithering

This is a simple project to demonstrate [Floyd–Steinberg dithering] as code.


## Requirements
* [Go 1.18]

## Bulid
```console
make build
```

## Run
```console
./build/dither -h
./build/dither --in foo.png --out bar.png
make examples
```

## Test

```console
make test
make lint # requires golangci-lint installed
```

## Examples

| Original                                            | Dithered                                                  |
|-----------------------------------------------------|-----------------------------------------------------------|
| ![american_gothic](examples/american_gothic/in.jpg) | ![american_gothic_dith](examples/american_gothic/out.jpg) |
| ![baby](examples/baby/in.png)                       | ![baby_dith](examples/baby/out.png)                       |
| ![milkyway](examples/milkyway/in.png)               | ![milkyway_dith](examples/milkyway/out.png)               |
| ![gopher](examples/gopher/in.png)                   | ![gopher_dith](examples/gopher/out.png)                   |
| ![michelangelo](examples/michelangelo/in.png)       | ![michelangelo_dith](examples/michelangelo/out.png)       |

| Factor   | Image                                   |
|----------|-----------------------------------------|
| Original | ![factor](examples/factors/in.png)      |
| 1        | ![factor_1](examples/factors/out_1.png) |
| 2        | ![factor_2](examples/factors/out_2.png) |
| 3        | ![factor_3](examples/factors/out_3.png) |
| 4        | ![factor_4](examples/factors/out_4.png) |

[Go 1.18]: https://go.dev/
[Floyd–Steinberg dithering]: https://en.wikipedia.org/wiki/Floyd%E2%80%93Steinberg_dithering
