# Description
Naive go bindings towards the C-API of CPython-3.6, gopy provides minimum runnable function set contains `Initialize`, `InsertPackagePath`, `CallFunc`, `Finalize` and so on.

Inspired by [go-python](https://github.com/sbinet/go-python).

# Environment

1. python3.6
2. pkg-config

```
add-apt-repository -y ppa:deadsnakes/ppa && apt install -y python3.6 python3.6-dev
apt install -y pkg-config
```

# Usage

See python_test.go file.
