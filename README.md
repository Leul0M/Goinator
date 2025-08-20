

# 🧞‍♂️ Goinator
A tiny [Akinator](https://en.akinator.com) clone in [Go](https://go.dev)
Ask questions, guess characters, learn new ones.

---
## 📸 Screenshot

![demo](https://github.com/Leul0M/Goinator/blob/main/Screenshot/image1.png)

![demo](https://github.com/Leul0M/Goinator/blob/main/Screenshot/image2.png)

---

# 🚀 Quick Start

# 1. Clone & enter
git clone https://github.com/Leul0M/Goinator.git

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
### 🧠 How the Magic Works

1️⃣ **Load the brain**  
   On start-up we read `data/entities.json` (23 yes/no traits per character).

2️⃣ **Build a smart tree**  
   Using **information-gain (ID3)** we build a decision tree that always asks the *most useful* question next—so you usually finish in **4–7 questions** instead of 23.

3️⃣ **Walk the tree**  
   Every answer (`y` / `n`) moves you down a branch until we hit a **leaf**.

4️⃣ **Bayesian tie-breaker**  
   If traits are missing and multiple characters are still possible, we rank them with **Naïve Bayes** and pick the highest-probability one.

5️⃣ **Learn on the fly**  
   Run `goinator learn` at any time to append or edit records—no recompile needed!

🔄 **Cycle**: play → miss → fix → play again. The more you teach it, the smarter it gets!
---


### ✅ Up next – polish & gotchas 🚧

🐛 **Data quality**  
Most traits were scraped from the internet; expect **occasional wrong labels**.  
Use `goinator learn --edit-id <id>` to fix them on the fly.

🎯 **Short TODO list**
- 🧮 **Smarter priors** – replace the hard-coded 0.9 / 0.1 with real probabilities.  
- 🖼️ **Richer UI** – add progress bar, colors, and a final “was I right?” screen.  
- 🔍 **Search & stats** – `/stats` command to list the most-confusing traits.  
- 🧹 **Auto-prune** – automatically drop questions that never split entities.  
- 🪄 **Persist session** – remember the last 10 games for quick replay.

Happy guessing! 🎲
