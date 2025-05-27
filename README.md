# go-kick-sdk

> An unofficial Go SDK to interact with Kick.com's public API.

[![Go Reference](https://pkg.go.dev/badge/github.com/Teyz/go-kick-sdk.svg)](https://pkg.go.dev/github.com/Teyz/go-kick-sdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/Teyz/go-kick-sdk)](https://goreportcard.com/report/github.com/Teyz/go-kick-sdk)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

---

## 📦 Installation

```bash
go get github.com/Teyz/go-kick-sdk
```

---

## 🚀 Usage Example

```go
package main

import (
    "fmt"
    "log"

    "github.com/teyz/go-kick-sdk/kick"
)

func main() {
    ctx := context.Background()
    client, err := kick.NewKickClient(context.Background(), "clientID", "clientSecret", "redirectURI")
    if err != nil {
        return
    }

    err = kick.ExchangeCode(context.Background(), "code", "codeVerifier")
    if err != nil {
        return
    }
}
```

---

## 🧰 Features

- ✅ OAuth
- ✅ Users

---

## 📚 Documentation

Full documentation is available on [pkg.go.dev](https://pkg.go.dev/github.com/Teyz/go-kick-sdk).

---

## ✅ TODO

- 🔄 Channels
- 💬 Real-time chat support
- 📺 Subscriptions, followers, etc.
- 🧪 Unit tests

---

## 🤝 Contributing

Contributions are welcome! Feel free to open an **issue** or a **pull request**.

---

## 📄 License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).
