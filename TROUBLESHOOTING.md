# 🔧 Guía de Solución de Problemas - Redirección después del Login

## ❌ Problema: No redirige a Home.html después del login exitoso

### 📋 Checklist de Diagnóstico

Sigue estos pasos en orden para identificar el problema:

---

## 1️⃣ Verificar que el Backend está funcionando

### Prueba 1: Health Check
```bash
curl http://localhost:8080/health
```
**Resultado esperado:** `OK`

### Prueba 2: Verificar que la ruta /static está funcionando
Abre en el navegador:
```
http://localhost:8080/static/Home.html
```
**Resultado esperado:** Deberías ver la página Home.html (aunque te redirija al login si no estás autenticado)

---

## 2️⃣ Probar el Login desde la Consola del Navegador

### Paso A: Abre la consola del navegador
1. Presiona `F12` en tu navegador
2. Ve a la pestaña "Console"

### Paso B: Ejecuta este código en la consola
```javascript
// Test 1: Verificar si fetch funciona
fetch('http://localhost:8080/health')
  .then(r => r.text())
  .then(data => console.log('Health check:', data))
  .catch(err => console.error('Error:', err));

// Test 2: Probar login manualmente
fetch('http://localhost:8080/api/login', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({
    user_email: 'TU_EMAIL_AQUI',
    user_password: 'TU_PASSWORD_AQUI'
  })
})
.then(r => r.json())
.then(data => {
  console.log('Login response:', data);
  if (data.user) {
    localStorage.setItem('user', JSON.stringify(data.user));
    console.log('Usuario guardado:', data.user);
    window.location.href = '/static/Home.html';
  }
})
.catch(err => console.error('Error:', err));
```

**Reemplaza `TU_EMAIL_AQUI` y `TU_PASSWORD_AQUI` con tus credenciales reales**

---

## 3️⃣ Verificar el archivo login.js

### Opción A: Usar la página de test
1. Ve a: `http://localhost:8080/static/test-redirect.html`
2. Haz clic en "Simular Login"
3. Ingresa tus credenciales
4. Observa los logs

### Opción B: Verificar que el archivo se está cargando
1. Abre `http://localhost:8080/static/login.html`
2. Presiona `F12` → pestaña "Network"
3. Recarga la página (`Ctrl+Shift+R` o `Cmd+Shift+R` en Mac)
4. Busca `login.js` en la lista
5. Verifica que el Status sea `200` (no `304`)

**Si ves `304`:** El navegador está usando caché. Haz un **hard reload**:
- Chrome/Edge: `Ctrl+Shift+R` (Windows) o `Cmd+Shift+R` (Mac)
- Firefox: `Ctrl+F5` (Windows) o `Cmd+Shift+R` (Mac)

---

## 4️⃣ Limpiar Caché del Navegador

A veces el navegador guarda la versión antigua del JavaScript.

### Chrome/Edge
1. Presiona `F12`
2. Click derecho en el botón de recarga
3. Selecciona "Vaciar caché y volver a cargar de forma forzada"

### Firefox
1. Presiona `Ctrl+Shift+Delete`
2. Selecciona "Caché"
3. Click en "Limpiar ahora"

### Safari
1. Menú Safari → Preferencias
2. Avanzado → Marcar "Mostrar menú Desarrollo"
3. Menú Desarrollo → Vaciar cachés

---

## 5️⃣ Verificar CORS

Si ves errores de CORS en la consola:

```
Access to fetch at 'http://localhost:8080/api/login' from origin 'http://localhost:8080' has been blocked by CORS policy
```

**Solución:** El backend ya tiene CORS configurado en `app.go`. Reinicia el servidor:
```bash
# Detén el servidor (Ctrl+C)
# Vuelve a iniciarlo
go run main.go
```

---

## 6️⃣ Verificar la Consola del Navegador

Después de hacer login, **deberías ver estos logs en la consola**:

```
🔄 Intentando login con: tu@email.com
📡 Response status: 200
📦 Response data: {message: "Login exitoso", user: {...}}
✅ Login exitoso!
💾 Usuario guardado en localStorage
🔄 Redirigiendo a Home.html...
```

### Si NO ves estos logs:
- El archivo `login.js` no se está cargando correctamente
- Haz un **hard reload** (Ctrl+Shift+R)
- Verifica que la ruta del script en `login.html` sea `/static/login.js`

### Si ves los logs pero no redirige:
- Puede haber un error de JavaScript después
- Busca errores en rojo en la consola
- Verifica que `/static/Home.html` sea accesible

---

## 7️⃣ Probar Redirección Manual

En la consola del navegador, después de hacer login exitoso, ejecuta:

```javascript
// Verificar que el usuario esté guardado
console.log('Usuario en localStorage:', localStorage.getItem('user'));

// Probar diferentes tipos de redirección
window.location.href = '/static/Home.html';  // Opción 1
// O
window.location.replace('/static/Home.html'); // Opción 2
// O
window.location = '/static/Home.html';        // Opción 3
```

---

## 8️⃣ Verificar el archivo login.html

Asegúrate de que el script esté correctamente referenciado:

```html
<!-- Debe estar al final del body -->
<script src="/static/login.js"></script>
```

**NO debe ser:**
- `<script src="login.js"></script>` (sin /static/)
- `<script src="./login.js"></script>` (con ./)
- `<script src="../login.js"></script>` (con ../)

---

## 9️⃣ Crear un Usuario de Prueba

Si no estás seguro de tus credenciales:

### Opción A: Desde la interfaz
1. Ve a `http://localhost:8080/static/login.html`
2. Haz clic en "Registrarse"
3. Completa el formulario con:
   - Nombre: Test User
   - Email: test@test.com
   - Password: 12345678
4. Regístrate
5. Luego haz login con esas credenciales

### Opción B: Desde la API con curl
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "user_name": "Test User",
    "user_email": "test@test.com",
    "user_password": "12345678"
  }'
```

---

## 🔟 Solución Alternativa: Archivo de Login Simplificado

Si nada funciona, prueba este código simplificado en `login.js`:

```javascript
document.addEventListener("DOMContentLoaded", () => {
  const loginForm = document.getElementById("loginForm");
  
  if (!loginForm) {
    console.error("❌ No se encontró el formulario de login");
    return;
  }

  loginForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    
    const email = document.getElementById("email")?.value;
    const password = document.getElementById("password")?.value;
    
    if (!email || !password) {
      alert("Por favor completa todos los campos");
      return;
    }

    try {
      const response = await fetch("http://localhost:8080/api/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ 
          user_email: email, 
          user_password: password 
        }),
      });

      const data = await response.json();

      if (response.ok && data.user) {
        localStorage.setItem("user", JSON.stringify(data.user));
        window.location.replace("/static/Home.html");
      } else {
        alert("Error: " + (data.error || "Credenciales inválidas"));
      }
    } catch (err) {
      alert("Error de conexión: " + err.message);
    }
  });
});
```

---

## 📊 Tabla de Diagnóstico Rápido

| Síntoma | Causa Probable | Solución |
|---------|---------------|----------|
| No pasa nada al hacer clic | Script no cargado | Hard reload (Ctrl+Shift+R) |
| Error de CORS | Backend no responde | Reiniciar servidor |
| "Credenciales inválidas" | Usuario no existe | Crear usuario primero |
| Redirige a login otra vez | Home.js rechaza sin auth | Verificar localStorage |
| Página en blanco | Error 404 en Home.html | Verificar ruta /static/Home.html |

---

## 🚨 Errores Comunes

### Error 1: "Cannot read property 'addEventListener' of null"
**Causa:** El elemento con ID no existe en el HTML
**Solución:** Verifica los IDs en `login.html`:
- `id="loginForm"` en el form de login
- `id="email"` en el input de email
- `id="password"` en el input de password

### Error 2: "NetworkError when attempting to fetch resource"
**Causa:** Backend no está corriendo
**Solución:** 
```bash
cd Go-Vehicle-Management-API-feature0001
go run main.go
```

### Error 3: Redirige pero vuelve al login
**Causa:** `Home.js` verifica autenticación y no encuentra usuario
**Solución:** Verifica que el usuario se guardó:
```javascript
console.log(localStorage.getItem('user'));
```

---

## ✅ Test Final

Ejecuta este script completo en la consola del navegador:

```javascript
console.clear();
console.log("=== TEST COMPLETO DE LOGIN Y REDIRECCIÓN ===\n");

// 1. Verificar servidor
fetch('http://localhost:8080/health')
  .then(r => r.text())
  .then(data => console.log("✅ Servidor funcionando:", data))
  .catch(() => console.error("❌ Servidor no responde"));

// 2. Verificar Home.html accesible
fetch('/static/Home.html')
  .then(r => console.log("✅ Home.html accesible, status:", r.status))
  .catch(() => console.error("❌ Home.html no accesible"));

// 3. Verificar localStorage
const user = localStorage.getItem('user');
console.log(user ? "✅ Usuario en localStorage" : "⚠️ No hay usuario en localStorage");

// 4. Test de redirección
console.log("\n🔄 Probando redirección en 3 segundos...");
setTimeout(() => {
  console.log("🚀 Redirigiendo...");
  window.location.href = '/static/Home.html';
}, 3000);
```

---

## 📞 Ayuda Adicional

Si después de todos estos pasos aún no funciona:

1. **Verifica la versión de Go:**
   ```bash
   go version
   ```
   Debería ser 1.19 o superior

2. **Reinstala dependencias:**
   ```bash
   go mod tidy
   go mod download
   ```

3. **Comparte estos datos:**
   - Navegador y versión
   - Sistema operativo
   - Mensajes de la consola del navegador
   - Logs del servidor Go

---

## 🎯 Solución Más Rápida

**Si tienes prisa y solo quieres que funcione:**

1. Cierra completamente tu navegador
2. Abre una ventana de incógnito/privada
3. Ve a `http://localhost:8080/static/login.html`
4. Haz login
5. Debería redirigir correctamente

**Esto funciona porque el modo incógnito no usa caché.**

---

**Última actualización:** 2025
**Versión:** 1.0.0