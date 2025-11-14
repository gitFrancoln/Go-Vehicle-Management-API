# Integración Frontend-Backend - VehículoMarket

## 📋 Resumen de Cambios

Este documento describe la integración completa entre el frontend y el backend de la aplicación VehículoMarket.

## 🔧 Cambios Realizados

### 1. **Sistema de Autenticación y Redirección**

#### `login.js`
- ✅ Agregada redirección automática a `Home.html` después de login exitoso
- ✅ Almacenamiento de datos del usuario en `localStorage`
- ✅ Manejo de errores mejorado

**Endpoint utilizado:** `POST /api/login`

```javascript
// Después de login exitoso
localStorage.setItem("user", JSON.stringify(data.user));
window.location.href = "/static/Home.html";
```

---

### 2. **Dashboard de Vehículos (Home.html)**

#### `Home.js` (NUEVO)
Archivo JavaScript creado para manejar toda la funcionalidad del dashboard:

**Funcionalidades principales:**
- ✅ Verificación de autenticación al cargar la página
- ✅ Carga de vehículos desde el backend
- ✅ Sistema de filtrado por marca, modelo, año y precio
- ✅ Paginación (6 vehículos por página)
- ✅ Vista de detalles de vehículos
- ✅ Función de cerrar sesión

**Endpoints utilizados:**
- `GET /api/vehicles/getAll` - Obtener todos los vehículos
- `GET /api/vehicles/:id` - Obtener detalles de un vehículo específico

**Ejemplo de uso:**
```javascript
// Cargar vehículos
const response = await fetch('http://localhost:8080/api/vehicles/getAll');
const data = await response.json();
allVehicles = data.vehicles || [];
```

#### `Home.html`
- ✅ Actualizado para usar `/static/aboutUs.css`
- ✅ Incluido script `/static/Home.js`
- ✅ Enlaces de navegación actualizados

---

### 3. **Sistema de Contacto**

#### Backend - Nuevos Archivos

**`models/contact.go`** (NUEVO)
```go
type Contact struct {
    gorm.Model
    Name    string `json:"name" binding:"required"`
    Email   string `json:"email" binding:"required,email"`
    Phone   string `json:"phone"`
    Message string `json:"message" binding:"required"`
}
```

**`handlers/contact_handler.go`** (NUEVO)
- Handler para crear, obtener y listar mensajes de contacto

**`repositories/contact_repository.go`** (NUEVO)
- Repositorio con métodos CRUD para contactos

**`services/contact_service.go`** (NUEVO)
- Servicio con lógica de negocio para contactos

#### Frontend

**`contact.js`** (ACTUALIZADO)
- ✅ Conexión con backend para enviar formularios
- ✅ Validación de campos (nombre, email, mensaje)
- ✅ Validación de formato de email
- ✅ Indicador de carga durante el envío
- ✅ Manejo de errores y respuestas del servidor

**Endpoint utilizado:** `POST /api/contact`

**`contact.html`** (ACTUALIZADO)
- ✅ Rutas corregidas para CSS y JS

---

### 4. **Actualización del Backend Principal**

#### `app/app.go`
**Cambios realizados:**

1. **Nuevas rutas HTML:**
```go
router.GET("/", func(c *gin.Context) {
    c.File(filepath.Join(staticPath, "aboutUs.html"))
})
router.GET("/login", func(c *gin.Context) {
    c.File(filepath.Join(staticPath, "login.html"))
})
router.GET("/home", func(c *gin.Context) {
    c.File(filepath.Join(staticPath, "Home.html"))
})
router.GET("/contact", func(c *gin.Context) {
    c.File(filepath.Join(staticPath, "contact.html"))
})
```

2. **Nuevas rutas de API:**
```go
// Contacto
router.POST("/api/contact", contactHandler.CreateContact)
router.GET("/api/contacts", contactHandler.GetAllContacts)
router.GET("/api/contacts/:id", contactHandler.GetContactByID)
```

3. **Integración de ContactHandler:**
- Repositorio de contacto inicializado
- Servicio de contacto creado
- Handler de contacto agregado
- Migración de tabla `contacts` en la base de datos

---

### 5. **Página Inicial (aboutUs.html)**
- ✅ Enlaces de navegación actualizados
- ✅ Rutas corregidas para CSS y JS
- ✅ Botón "Iniciar Sesión" redirige a `/static/login.html`

---

## 🚀 Cómo Ejecutar la Aplicación

### 1. Iniciar el Backend

```bash
cd Go-Vehicle-Management-API-feature0001
go run main.go
```

El servidor se iniciará en `http://localhost:8080`

### 2. Acceder a la Aplicación

Abre tu navegador y ve a:
```
http://localhost:8080
```

---

## 📍 Flujo de Usuario

### Flujo Completo:

1. **Página Principal** → `http://localhost:8080` (aboutUs.html)
2. **Clic en "Iniciar Sesión"** → `/static/login.html`
3. **Login Exitoso** → Redirección automática a `/static/Home.html`
4. **Dashboard** → Ver, filtrar y explorar vehículos
5. **Contacto** → `/static/contact.html` para enviar mensajes

---

## 🗄️ Estructura de la Base de Datos

### Tablas Creadas Automáticamente:

1. **users** - Usuarios del sistema
2. **vehicles** - Vehículos disponibles
3. **contacts** - Mensajes de contacto (NUEVA)

---

## 🔑 Endpoints de API Disponibles

### Usuarios
- `POST /api/users` - Crear usuario
- `GET /api/users/:id` - Obtener usuario
- `PUT /api/users/:id` - Actualizar usuario
- `DELETE /api/users/:id` - Eliminar usuario
- `POST /api/login` - Iniciar sesión
- `POST /api/forgot-password` - Recuperar contraseña

### Vehículos
- `POST /api/vehicles` - Crear vehículo
- `GET /api/vehicles/:id` - Obtener vehículo por ID
- `PUT /api/vehicles/:id` - Actualizar vehículo
- `GET /api/vehicles/getAll` - Obtener todos los vehículos

### Contacto (NUEVO)
- `POST /api/contact` - Enviar mensaje de contacto
- `GET /api/contacts` - Obtener todos los contactos
- `GET /api/contacts/:id` - Obtener contacto por ID

---

## 🎨 Archivos de Estilos

- `login.css` - Estilos del login/registro
- `aboutUs.css` - Estilos de la página principal y dashboard
- `contact.css` - Estilos del formulario de contacto

---

## 📱 Funcionalidades del Dashboard (Home.html)

### Filtros Disponibles:
- **Marca**: Mercedes, Ford, Honda, Toyota, BMW, Volkswagen, Tesla, Audi
- **Modelo**: C-Class, Mustang, CR-V, Camry, X5, Golf, Model 3, A4
- **Año**: 2018-2023
- **Precio**: Rango mínimo y máximo

### Acciones:
- **Ver Detalles**: Muestra información completa del vehículo
- **Contactar Vendedor**: Funcionalidad en desarrollo
- **Aplicar Filtros**: Filtra vehículos según criterios
- **Restablecer**: Limpia todos los filtros

### Paginación:
- 6 vehículos por página
- Botones de navegación (Anterior/Siguiente)
- Botones de página numerados

---

## 🔒 Seguridad

### Autenticación:
- Los datos del usuario se almacenan en `localStorage` después del login
- `Home.js` verifica la autenticación al cargar la página
- Si no hay usuario autenticado, redirige automáticamente a login

```javascript
const user = JSON.parse(localStorage.getItem('user'));
if (!user) {
    alert('⚠️ Debes iniciar sesión para acceder');
    window.location.href = '/static/login.html';
    return;
}
```

---

## 🐛 Manejo de Errores

Todos los archivos JavaScript incluyen manejo de errores completo:

```javascript
try {
    const response = await fetch('...');
    // ... código
} catch (err) {
    console.error('Error de conexión:', err);
    alert('⚠️ Error de conexión con el servidor');
}
```

---

## 📝 Notas Importantes

1. **CORS está habilitado** en el backend para permitir peticiones desde el frontend
2. **Todas las rutas estáticas** usan el prefijo `/static/`
3. **El servidor backend** debe estar ejecutándose en `http://localhost:8080`
4. **Las imágenes de vehículos** usan placeholders si no hay imagen disponible

---

## 🔄 Próximas Mejoras Sugeridas

- [ ] Implementar funcionalidad "Contactar Vendedor"
- [ ] Agregar sistema de favoritos
- [ ] Implementar carrito de compras funcional
- [ ] Agregar página de perfil de usuario
- [ ] Implementar sistema de notificaciones
- [ ] Agregar más filtros (tipo de combustible, transmisión, etc.)
- [ ] Implementar búsqueda por texto
- [ ] Agregar galería de imágenes para cada vehículo

---

## 🆘 Solución de Problemas

### El dashboard no carga vehículos:
1. Verifica que el backend esté ejecutándose
2. Asegúrate de tener vehículos en la base de datos
3. Revisa la consola del navegador (F12) para ver errores

### No puedo acceder al dashboard después del login:
1. Verifica que el login sea exitoso
2. Revisa que `localStorage` tenga los datos del usuario
3. Verifica que la ruta `/static/Home.html` esté accesible

### El formulario de contacto no envía datos:
1. Verifica que todos los campos obligatorios estén completos
2. Asegúrate de que el email tenga formato válido
3. Revisa la consola para ver errores de conexión

---

## 👨‍💻 Archivos Modificados/Creados

### Nuevos Archivos:
- `static/Home.js`
- `models/contact.go`
- `handlers/contact_handler.go`
- `repositories/contact_repository.go`
- `services/contact_service.go`
- `README_FRONTEND_BACKEND_INTEGRATION.md`

### Archivos Modificados:
- `static/login.js`
- `static/Home.html`
- `static/aboutUs.html`
- `static/contact.html`
- `static/contact.js`
- `app/app.go`

---

**✅ Integración Completa: Frontend ↔️ Backend**

Todas las páginas ahora están conectadas y funcionando con el backend en `http://localhost:8080`
