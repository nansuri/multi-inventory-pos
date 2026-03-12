# Invent Story

> Every Story start with small story.

Invent Story is a modern, open-source multi-tenant inventory and production management system designed specifically for restaurants. It helps you track raw ingredients, manage complex recipes, and monitor cooked stock levels with a sleek, intuitive interface.

## 🌟 Key Features

- **Multi-Tenancy:** Robust data isolation for multiple restaurant outlets.
- **Recipe Management:** Define "Blueprints" for your dishes with precise ingredient requirements.
- **Production Hall:** Record batch preparations (cooking) that automatically deduct raw materials and increase saleable stock.
- **POS & Table Management:** A receipt-style interface for waiters to place orders and manage floor plans.
- **Multi-Language Support:** Localized in English, Bahasa Indonesia, and Japanese.
- **Multi-Currency:** Support for USD, IDR, and JPY with dynamic formatting.
- **Dashboard Analytics:** Real-time metrics on revenue, orders, and low-stock alerts.

## 🚀 Tech Stack

- **Backend:** Go (Gin, GORM, Postgres, Redis) - Domain-Driven Design (DDD).
- **Frontend:** Vue 3 (TypeScript, Vite, Tailwind CSS 4, Pinia).
- **Infrastructure:** Docker & Docker Compose.

## 🛠️ Quick Start (Docker)

1. **Clone the repository.**
2. **Setup Configuration:** The project uses a root `.env` file for all services.
3. **Launch:**
   ```bash
   docker-compose up --build
   ```
4. **Access:**
   - Web UI: `http://localhost`
   - API: `http://localhost:8080`

## 👨‍💻 Development

### Backend
```bash
cd backend
go run cmd/api/main.go
```

### Frontend
```bash
cd frontend
npm install
npm run dev
```

## 📄 License

This project is licensed under the **MIT License** - see the [LICENSE](LICENSE) file for details.

---
Built with ❤️ for the restaurant community.
