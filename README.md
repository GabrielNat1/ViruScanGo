# 🛡️ VirusScannerGo

**VirusScannerGo** is a lightweight microservice written in Go that scans files uploaded via HTTP for potentially malicious content, such as embedded scripts, disguised executables, and other suspicious patterns.

This project focuses on protecting file uploads in systems used by schools, NGOs, and libraries, offering a fast, efficient, and easily integrable solution.

---

## 🛠️ How to Run

```bash
git clone https://github.com/your-username/virus-scanner-go.git
```

```bash
cd virus-scanner-go
```

```bash
go mod tidy
```

```bash
go run main.go
```

The API will be available at: `http://localhost:8080`

---

## 🚀 Features

- 📦 File upload via HTTP API  
- 🔍 Static analysis to detect common threats  
- ⚡ Concurrent processing with Go workers  
- 🧠 Smart caching for already scanned files  
- 📈 Prometheus metrics endpoint  

---
