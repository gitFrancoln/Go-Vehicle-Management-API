# 📁 Estructura del Proyecto - VehículoMarket

## 🌳 Árbol de Directorios

```
Go-Vehicle-Management-API-feature0001/
│
├── 📂 app/
│   └── app.go                          # Configuración principal de la aplicación
│
├── 📂 cmd/
│   └── [archivos de comandos]          # Comandos CLI (si existen)
│
├── 📂 database/
│   └── database.go                     # Configuración de conexión a BD
│
├── 📂 dto/
│   └── [Data Transfer Objects]         # Objetos de transferencia de datos
│
├── 📂 handlers/
│   ├── user_handler.go                 # Handlers de usuarios
│   ├── vehicle_handler.go              # Handlers de vehículos
│   └── contact_handler.go              # Handlers de contacto ⭐ NUEVO
│
├── 📂 internal/
│   └── [paquetes internos]             # Código interno de la aplicación
│
├── 📂 middleware/
│   └── [middlewares]                   # Middlewares (autenticación, CORS, etc.)
│
├── 📂 models/
│   ├── user.go                         # Modelo de Usuario
│   ├── vehicles.go                     # Modelo de Vehículo
│   └── contact.go                      # Modelo de Contacto ⭐ NUEVO
│
├── 📂 repositories/
│   ├── user_repository.go              # Repositorio de usuarios
│   ├── vehicle_repository.go           # Repositorio de vehículos
│   └── contact_repository.go           # Repositorio de contacto ⭐ NUEVO
│
├── 📂 services/
│   ├── user_service.go                 # Servicio de usuarios
│   ├── vehicle_service.go              # Servicio de vehículos
│   └── contact_service.go              # Servicio de contacto ⭐ NUEVO
│
├── 📂 static/                          # 🎨 FRONTEND
│   │
│   ├── 📄 aboutUs.html                 # Página principal ✅
│   ├── 📄 aboutUs.css                  # Estilos de la página principal ✅
│   ├── 📄 aboutUs.js                   # Scripts de la página principal ✅
│   │
│   ├── 📄 login.html                   # Página de login/registro ✅
│   ├── 📄 login.css                    # Estilos de login ✅
│   ├── 📄 login.js                     # Scripts de login ✅ MODIFICADO
│   │
│   ├── 📄 Home.html                    # Dashboard de vehículos ✅ MODIFICADO
│   ├── 📄 Home.js                      # Scripts del dashboard ⭐ NUEVO
│   │
│   ├── 📄 contact.html                 # Formulario de contacto ✅ MODIFICADO
│   ├── 📄 contact.css                  # Estilos del contacto ✅
│   ├── 📄 contact.js                   # Scripts del contacto ✅ MODIFICADO
│   │
│   └── 📂 imagenes/                    # Imágenes y recursos
│       ├── logo.png
│       ├── logo2.png
│       └── golfgt.jpg
│
├── 📄 main.go                          # Punto de entrada de la aplicación
├── 📄 go.mod                           # Dependencias de Go
├── 📄 go.sum                           # Checksums de dependencias
│
├── 📄 sample_data.sql                  # Script SQL con datos de ejemplo ⭐ NUEVO
├── 📄 README_FRONTEND_BACKEND_INTEGRATION.md  # Documentación técnica ⭐ NUEVO
├── 📄 QUICK_START.md                   # Guía de inicio rápido ⭐ NUEVO
└── 📄 PROJECT_STRUCTURE.md             # Este archivo ⭐ NUEVO
```

---

## 🏗️ Arquitectura del Backend

### Capas de la Aplicación

```
┌─────────────────────────────────────┐
│         HTTP REQUESTS               │
│      (Gin Router - app.go)          │
└───────────────┬─────────────────────┘
                │
                ▼
┌─────────────────────────────────────┐
│          HANDLERS                   │
│  (Reciben requests, validan datos)  │
│  - user_handler.go                  │
│  - vehicle_handler.go               │
│  - contact_handler.go               │
└───────────────┬─────────────────────┘
                │
                ▼
┌─────────────────────────────────────┐
│          SERVICES                   │
│    (Lógica de negocio)              │
│  - user_service.go                  │
│  - vehicle_service.go               │
│  - contact_service.go               │
└───────────────┬─────────────────────┘
                │
                ▼
┌─────────────────────────────────────┐
│        REPOSITORIES                 │
│   (Acceso a base de datos)          │
│  - user_repository.go               │
│  - vehicle_repository.go            │
│  - contact_repository.go            │
└───────────────┬─────────────────────┘
                │
                ▼
┌─────────────────────────────────────┐
│         DATABASE (GORM)             │
│        MySQL / PostgreSQL           │
└─────────────────────────────────────┘
```

---

## 🎨 Arquitectura del Frontend

### Flujo de Páginas

```
┌─────────────────────┐
│   aboutUs.html      │  ← Página de inicio
│   (Landing Page)    │
└──────────┬──────────┘
           │
           │ Clic "Iniciar Sesión"
           ▼
┌─────────────────────┐
│    login.html       │  ← Login / Registro
│                     │
└──────────┬──────────┘
           │
           │ Login exitoso
           ▼
┌─────────────────────┐
│    Home.html        │  ← Dashboard de vehículos
│  (Dashboard)        │
│  - Ver vehículos    │
│  - Filtrar          │
│  - Ver detalles     │
└──────────┬──────────┘
           │
           │ Clic "Contacto"
           ▼
┌─────────────────────┐
│   contact.html      │  ← Formulario de contacto
│                     │
└─────────────────────┘
```

### Conexiones Frontend-Backend

```javascript
// login.js
POST /api/login           → Iniciar sesión
POST /api/users           → Registrar usuario

// Home.js
GET /api/vehicles/getAll  → Obtener todos los vehículos
GET /api/vehicles/:id     → Obtener detalles de vehículo

// contact.js
POST /api/contact         → Enviar mensaje de contacto
```

---

## 🗄️ Modelos de Base de Datos

### User (usuarios)
```go
type User struct {
    gorm.Model
    Name     string
    Email    string (unique)
    Password string (hashed)
}
```

### Vehicle (vehículos)
```go
type Vehicle struct {
    gorm.Model
    Brand        string
    Model        string
    Year         int
    Price        float64
    Mileage      int
    Transmission string
    FuelType     string
    Color        string
    BodyType     string
    Status       string
    Description  string
    Images       []string (JSON)
}
```

### Contact (contactos) ⭐ NUEVO
```go
type Contact struct {
    gorm.Model
    Name    string
    Email   string
    Phone   string
    Message string
}
```

---

## 🔌 Endpoints de API

### 👤 Usuarios
| Método | Endpoint | Descripción |
|--------|----------|-------------|
| POST | `/api/users` | Crear usuario |
| GET | `/api/users/:id` | Obtener usuario |
| PUT | `/api/users/:id` | Actualizar usuario |
| DELETE | `/api/users/:id` | Eliminar usuario |
| POST | `/api/login` | Iniciar sesión |
| POST | `/api/forgot-password` | Recuperar contraseña |

### 🚗 Vehículos
| Método | Endpoint | Descripción |
|--------|----------|-------------|
| POST | `/api/vehicles` | Crear vehículo |
| GET | `/api/vehicles/:id` | Obtener vehículo |
| PUT | `/api/vehicles/:id` | Actualizar vehículo |
| GET | `/api/vehicles/getAll` | Listar todos |

### 📧 Contacto ⭐ NUEVO
| Método | Endpoint | Descripción |
|--------|----------|-------------|
| POST | `/api/contact` | Enviar mensaje |
| GET | `/api/contacts` | Listar mensajes |
| GET | `/api/contacts/:id` | Obtener mensaje |

### 🏥 Health Check
| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/health` | Estado del servidor |

---

## 📦 Dependencias Principales

### Backend (Go)
```
- gin-gonic/gin         → Framework web
- gorm.io/gorm          → ORM para base de datos
- gorm.io/driver/mysql  → Driver MySQL
- go-playground/validator → Validación de datos
```

### Frontend (JavaScript Vanilla)
```
- Ionicons              → Iconos
- Fetch API             → Peticiones HTTP
- LocalStorage          → Almacenamiento local
```

---

## 🎯 Características Clave

### ✅ Implementadas
- [x] Sistema de autenticación (login/registro)
- [x] CRUD completo de usuarios
- [x] CRUD completo de vehículos
- [x] Dashboard con filtros y paginación
- [x] Formulario de contacto funcional
- [x] Diseño responsivo
- [x] Validación de formularios
- [x] Manejo de errores
- [x] CORS habilitado
- [x] Rutas protegidas (frontend)

### 🔄 En Desarrollo
- [ ] Carrito de compras
- [ ] Sistema de favoritos
- [ ] Perfil de usuario completo
- [ ] Galería de imágenes por vehículo
- [ ] Sistema de notificaciones
- [ ] Búsqueda avanzada

---

## 🔒 Seguridad

### Backend
- Contraseñas hasheadas (bcrypt)
- Validación de datos con validator
- CORS configurado
- Sanitización de inputs

### Frontend
- Validación de formularios
- Verificación de autenticación
- Tokens en localStorage
- Validación de emails

---

## 🚀 Flujo de Datos

### Ejemplo: Ver Vehículos

```
1. Usuario abre Home.html
   ↓
2. Home.js verifica autenticación
   ↓
3. Si autenticado → fetch('/api/vehicles/getAll')
   ↓
4. Backend (vehicle_handler) recibe request
   ↓
5. vehicle_service procesa la solicitud
   ↓
6. vehicle_repository consulta la BD
   ↓
7. Respuesta JSON → Frontend
   ↓
8. Home.js renderiza las tarjetas de vehículos
   ↓
9. Usuario ve los vehículos en pantalla
```

---

## 📝 Archivos Modificados vs Nuevos

### ⭐ Archivos NUEVOS
```
backend/
  ├── models/contact.go
  ├── handlers/contact_handler.go
  ├── repositories/contact_repository.go
  └── services/contact_service.go

frontend/
  └── static/Home.js

docs/
  ├── README_FRONTEND_BACKEND_INTEGRATION.md
  ├── QUICK_START.md
  ├── PROJECT_STRUCTURE.md
  └── sample_data.sql
```

### ✏️ Archivos MODIFICADOS
```
backend/
  └── app/app.go (rutas y handlers)

frontend/
  ├── static/login.js (redirección)
  ├── static/Home.html (script y CSS)
  ├── static/aboutUs.html (enlaces)
  ├── static/contact.html (rutas)
  └── static/contact.js (backend integration)
```

---

## 🎓 Patrones de Diseño Utilizados

1. **Repository Pattern** - Separación de lógica de datos
2. **Service Layer Pattern** - Lógica de negocio independiente
3. **MVC** (Modelo-Vista-Controlador) - Estructura general
4. **Dependency Injection** - Inyección de dependencias
5. **RESTful API** - Endpoints siguiendo REST

---

## 🌐 URLs del Proyecto

**Base URL:** `http://localhost:8080`

| Tipo | URL Completa |
|------|--------------|
| Home | `http://localhost:8080/` |
| Login | `http://localhost:8080/static/login.html` |
| Dashboard | `http://localhost:8080/static/Home.html` |
| Contacto | `http://localhost:8080/static/contact.html` |
| API Docs | `http://localhost:8080/api/*` |

---

## 📊 Base de Datos

### Tablas Creadas Automáticamente (GORM Auto-Migrate)
1. `users` - Usuarios del sistema
2. `vehicles` - Vehículos disponibles
3. `contacts` - Mensajes de contacto

### Columnas Estándar (GORM)
Todas las tablas incluyen:
- `id` (Primary Key)
- `created_at`
- `updated_at`
- `deleted_at` (Soft delete)

---

**📌 Nota:** Este proyecto sigue una arquitectura limpia y escalable, facilitando el mantenimiento y la adición de nuevas funcionalidades.

**✨ Versión:** 1.0.0  
**📅 Última actualización:** 2025