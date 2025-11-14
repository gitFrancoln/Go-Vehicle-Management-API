# 🚀 Guía de Inicio Rápido - VehículoMarket

## ⚡ Inicio Rápido (3 pasos)

### 1. Iniciar el Servidor Backend
```bash
cd Go-Vehicle-Management-API-feature0001
go run main.go
```

### 2. Abrir el Navegador
```
http://localhost:8080
```

### 3. ¡Listo! Navega por la aplicación

---

## 📱 Flujo de la Aplicación

```
Página Inicial (aboutUs.html)
    ↓
Clic en "Iniciar Sesión"
    ↓
Login / Registro
    ↓
Dashboard (Home.html) - Ver vehículos
    ↓
Formulario de Contacto
```

---

## 🔑 Credenciales de Prueba

Si ya tienes usuarios creados, usa tus credenciales.
Si no, regístrate primero usando el formulario de registro.

### Para crear un usuario:
1. Ve a http://localhost:8080/static/login.html
2. Clic en "Registrarse"
3. Completa el formulario
4. Inicia sesión

---

## 📊 Datos de Ejemplo

Si tu base de datos está vacía y quieres datos de prueba:

### Opción 1: Usar la API
Abre Postman o cualquier cliente HTTP y envía:

```http
POST http://localhost:8080/api/vehicles
Content-Type: application/json

{
  "brand": "Tesla",
  "model": "Model 3",
  "year": 2023,
  "price": 45000,
  "mileage": 8000,
  "transmission": "Automática",
  "fuel_type": "Eléctrico",
  "color": "Blanco",
  "body_type": "Sedán",
  "status": "Disponible",
  "description": "Tesla Model 3 2023, vehículo eléctrico de alto rendimiento."
}
```

### Opción 2: Usar el archivo SQL
```bash
# Conecta a tu base de datos y ejecuta:
mysql -u tu_usuario -p tu_base_de_datos < sample_data.sql
```

---

## 🌐 URLs Importantes

| Página | URL |
|--------|-----|
| **Inicio** | http://localhost:8080 |
| **Login** | http://localhost:8080/static/login.html |
| **Dashboard** | http://localhost:8080/static/Home.html |
| **Contacto** | http://localhost:8080/static/contact.html |
| **Health Check** | http://localhost:8080/health |

---

## 🔧 Endpoints de API

### Usuarios
- `POST /api/users` - Registrar usuario
- `POST /api/login` - Iniciar sesión
- `GET /api/users/:id` - Obtener usuario

### Vehículos
- `POST /api/vehicles` - Crear vehículo
- `GET /api/vehicles/getAll` - Listar todos
- `GET /api/vehicles/:id` - Ver detalles

### Contacto
- `POST /api/contact` - Enviar mensaje
- `GET /api/contacts` - Ver todos los mensajes

---

## ✅ Verificación de Funcionamiento

### 1. Backend funcionando:
```bash
curl http://localhost:8080/health
# Debería responder: OK
```

### 2. Ver vehículos:
```bash
curl http://localhost:8080/api/vehicles/getAll
# Debería devolver JSON con vehículos
```

---

## 🎯 Funcionalidades Principales

### ✨ Dashboard de Vehículos
- Ver todos los vehículos disponibles
- Filtrar por marca, modelo, año y precio
- Ver detalles de cada vehículo
- Paginación (6 vehículos por página)

### 🔐 Autenticación
- Registro de nuevos usuarios
- Inicio de sesión seguro
- Protección de rutas privadas
- Cierre de sesión

### 📧 Formulario de Contacto
- Enviar mensajes al sistema
- Validación de campos
- Guardado en base de datos

---

## 🐛 Solución Rápida de Problemas

### ❌ Error: "cannot find package"
```bash
go mod tidy
go mod download
```

### ❌ Error: "database connection failed"
Verifica tu configuración de base de datos en `database/database.go`

### ❌ No se cargan los vehículos
1. Verifica que el backend esté corriendo
2. Verifica que tengas vehículos en la BD
3. Revisa la consola del navegador (F12)

### ❌ No puedo iniciar sesión
1. Asegúrate de estar registrado primero
2. Verifica tus credenciales
3. Revisa la consola del navegador

---

## 📦 Requisitos

- **Go** 1.19 o superior
- **MySQL** o base de datos compatible
- **Navegador web** moderno (Chrome, Firefox, Edge, Safari)

---

## 📚 Documentación Completa

Para más detalles, consulta:
- `README_FRONTEND_BACKEND_INTEGRATION.md` - Documentación técnica completa

---

## 💡 Tips

1. **Abre la consola del navegador** (F12) para ver logs y errores
2. **Usa el LocalStorage** para ver datos del usuario guardados
3. **Prueba todos los filtros** en el dashboard
4. **El formulario de contacto** guarda los mensajes en la BD

---

## 🎉 ¡Eso es todo!

Ya puedes empezar a usar VehículoMarket. Explora, filtra vehículos y disfruta de la aplicación.

**¿Necesitas ayuda?** Revisa los logs del servidor y la consola del navegador.

---

**Hecho con ❤️ para VehículoMarket**