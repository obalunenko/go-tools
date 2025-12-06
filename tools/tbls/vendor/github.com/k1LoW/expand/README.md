# expand [![CI](https://github.com/k1LoW/expand/actions/workflows/ci.yml/badge.svg)](https://github.com/k1LoW/expand/actions/workflows/ci.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/k1LoW/expand.svg)](https://pkg.go.dev/github.com/k1LoW/expand) ![Coverage](https://raw.githubusercontent.com/k1LoW/octocovs/main/badges/k1LoW/expand/coverage.svg) ![Code to Test Ratio](https://raw.githubusercontent.com/k1LoW/octocovs/main/badges/k1LoW/expand/ratio.svg) ![Test Execution Time](https://raw.githubusercontent.com/k1LoW/octocovs/main/badges/k1LoW/expand/time.svg)

`expand` package provides convenient functions to apply [`func os.Expand`](https://pkg.go.dev/os#Expand) efficiently.

## Import

``` golang
import "github.com/k1LoW/expand"
```

## Usage

``` golang
c := &Config{}
p := "config.yml"
buf, err := os.ReadFile(p)
if err != nil {
    return err
}
if err := yaml.Unmarshal(expand.ExpandenvYAMLBytes(buf), c); err != nil {
    return err
}
```
