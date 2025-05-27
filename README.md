# go-kick-sdk

> Un SDK non-officiel pour interagir avec l'API publique de Kick.com, écrit en Go.

[![Go Reference](https://pkg.go.dev/badge/github.com/Teyz/go-kick-sdk.svg)](https://pkg.go.dev/github.com/Teyz/go-kick-sdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/Teyz/go-kick-sdk)](https://goreportcard.com/report/github.com/Teyz/go-kick-sdk)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

---

## 📦 Installation

```bash
go get github.com/Teyz/go-kick-sdk
```

---

## 🚀 Exemple d'utilisation

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

## 🧰 Fonctionnalités

- ✅ Récupération de l'utilisateur connecté
- ✅ Récupération d'utilisateur par nom d'utilisateur

---

## 📚 Documentation

La documentation complète est disponible sur [pkg.go.dev](https://pkg.go.dev/github.com/Teyz/go-kick-sdk).

---

## ✅ TODO

- 📺 Channels
- 💬 Gestion du chat en temps réel
- 📺 Abonnements, followers, etc.
- 🧪 Tests unitaires

---

## 🤝 Contribuer

Les contributions sont les bienvenues ! N'hésitez pas à ouvrir une **issue** ou une **pull request**.

---

## 📄 Licence

Ce projet est sous licence [MIT](https://opensource.org/licenses/MIT).
