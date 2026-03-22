# 🏋️ Clarity Gym

O aplicație web completă pentru managementul unei săli de fitness, dezvoltată ca proiect academic pentru materia **Proiectarea și Dezvoltarea Aplicațiilor Web**.

---

## 🚀 Tehnologii Folosite

### Frontend
- **Vue 3** + Vue Router + Pinia
- **Tailwind CSS v3** — stilizare responsive
- **Axios** — comunicare cu API-ul
- **Cropperjs** — decupare poze de profil

### Backend
- **Go (Golang)** + Chi Router
- **JWT** — autentificare și autorizare
- **bcrypt** — criptare parole
- **CORS** — comunicare cross-origin

### Baza de Date
- **PostgreSQL** — stocare date
- Migrații SQL manuale

---

## 👥 Roluri și Funcționalități

### 🙋 Client
- Vizualizare antrenori reali din baza de date
- Trimitere cereri de colaborare către antrenori
- Notificări persistente când cererea e acceptată/respinsă
- Vizualizare abonamente cu prețuri actualizate în timp real
- Formular de suport tehnic
- Profil editabil cu upload poză (crop inclus)

### 💪 Antrenor
- Vizualizare și gestionare cereri de colaborare
- Acceptare/respingere cereri
- Listă clienți activi
- Profil editabil cu specialitate, experiență, bio, poză

### 🔧 Admin
- Dashboard complet cu toți utilizatorii
- Suspendare/ștergere utilizatori
- Gestionare abonamente (CRUD complet) — salvate în DB
- Vizualizare mesaje suport — marcare ca rezolvate
- Notificări cu numărul de mesaje deschise

---

## 🔐 Autentificare

- Înregistrare cu validare pe frontend:
  - Email valid (Gmail sau Yahoo)
  - Parolă cu minim 8 caractere, literă mare, număr și caracter special
  - Vârstă minimă 16 ani
- Parolele sunt criptate cu **bcrypt** în baza de date
- Autentificare cu **JWT Token** salvat în localStorage
- Routing protejat pe roluri — fiecare utilizator vede doar pagina sa

---

## 🗄️ Schema Baza de Date

```sql
users               — utilizatori (client, trainer, admin)
subscription_plans  — planuri abonamente
user_subscriptions  — abonamente active
collaboration_requests — cereri colaborare client-antrenor
support_requests    — mesaje suport tehnic
```

---

## 📱 Design & UX

- **Responsive** — mobil, tabletă, desktop
- **Navbar hamburger** pe mobil
- **Animații fade-in** la scroll
- **Hover** pe carduri
- **Loading screen** la pornire
- **Tranziții** între pagini
- Upload poză profil cu **crop/decupare**
- Paleta de culori: crem, bej, verde sage, maro

---

## ⚙️ Instalare și Rulare

### Cerințe
- Go 1.21+
- Node.js 18+
- PostgreSQL 14+

### Pași

```bash
# 1. Clonează proiectul
git clone https://github.com/FlaviusMarian27/Clarity-Gym.git
cd Clarity-Gym

# 2. Configurează baza de date
sudo -u postgres psql -c "CREATE DATABASE clarity_gym;"
sudo -u postgres psql -d clarity_gym -f backend/db/migrations/001_users.sql
sudo -u postgres psql -d clarity_gym -f backend/db/migrations/002_subscriptions.sql
sudo -u postgres psql -d clarity_gym -f backend/db/migrations/003_collaborations.sql

# 3. Pornire rapidă
chmod +x start.sh
./start.sh
```

### Sau manual

```bash
# Backend
cd backend
go run cmd/main.go

# Frontend (alt terminal)
cd frontend
npm install
npm run dev
```

### Accesare
- Frontend: http://localhost:5173
- Backend API: http://localhost:8080

---

## 🔑 Conturi de Test

| Email | Parolă | Rol |
|-------|--------|-----|
| client@test.com | 1234 | Client |
| trainer@test.com | 1234 | Antrenor |
| admin@test.com | 1234 | Admin |

---

## 📁 Structură Proiect

```
Clarity-Gym/
├── frontend/
│   ├── src/
│   │   ├── components/     # Navbar-uri
│   │   ├── views/          # Pagini principale
│   │   ├── stores/         # Pinia (auth)
│   │   ├── router/         # Vue Router
│   │   └── composables/    # useScrollAnimation
│   └── public/             # Assets statice
├── backend/
│   ├── cmd/main.go         # Entry point
│   ├── config/             # Configurare DB
│   ├── internal/
│   │   ├── auth/           # Login, Register, JWT
│   │   ├── user/           # Profil, Admin
│   │   ├── trainer/        # Antrenori
│   │   ├── collaboration/  # Cereri colaborare
│   │   ├── subscription/   # Abonamente
│   │   └── support/        # Suport tehnic
│   ├── db/migrations/      # SQL migrații
│   └── uploads/avatars/    # Poze profil
└── start.sh                # Script pornire
```

---

## 👨‍💻 Autor

**Flavius Marian**  
Proiect academic — Proiectarea și Dezvoltarea Aplicațiilor Web