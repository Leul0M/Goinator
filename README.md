

# 🧞‍♂️ Goinator
A tiny Akinator clone in Go
Ask questions, guess characters, learn new ones.

---

# 🚀 Quick Start


# 1. Clone & enter
git clone https://github.com/yourname/Goinator.git
cd Goinator

# 2. Run
go run .




## 🎮 How to Play

| Key | Action |
|-----|--------|
| `y` | answer **yes** |
| `n` | answer **no** |
| `q` / `Ctrl-C` | quit |



## 🧠 Data

All entities live in  
[`data/entities.json`](data/entities.json)  
with 23 boolean traits such as `is_real`, `can_fly`, `is_youtuber` …



## 🔧 Commands

| Command | Purpose |
|---------|---------|
| `go run .` | start the game |
| `go run . learn` | add a new entity interactively |
| `go run . --stats` | see tree statistics |



## 🛠️ Build

```bash
go build -o goinator
./goinator
```

---

## 📸 Screenshot

![demo](https://user-images.githubusercontent.com/yourname/.../goinator.gif)
*(replace with your own GIF)*

---

## 📄 License

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
```
