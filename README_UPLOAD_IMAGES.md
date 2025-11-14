# 📸 Sistema de Carga de Imágenes para Vehículos

## 🎯 Descripción

Este sistema permite subir y gestionar imágenes de vehículos en la aplicación VehículoMarket. Las imágenes se almacenan físicamente en el servidor y sus rutas se guardan en la base de datos.

## 🚀 Características

- ✅ Subida de múltiples imágenes por vehículo
- ✅ Validación de formatos de imagen (jpg, jpeg, png, gif, webp)
- ✅ Preview de imágenes antes de subir
- ✅ Nombres únicos con timestamp para evitar colisiones
- ✅ Actualización automática de la base de datos
- ✅ Interfaz amigable y responsiva

## 📋 Endpoints Disponibles

### 1. Subir Imágenes a un Vehículo
```
POST /api/vehicles/:id/upload-images
```

**Parámetros:**
- `id` (URL param): UUID del vehículo
- `images` (FormData): Array de archivos de imagen

**Respuesta exitosa:**
```json
{
  "message": "Imágenes subidas exitosamente",
  "vehicle_id": "123e4567-e89b-12d3-a456-426614174000",
  "images": [
    "/images/123e4567-e89b-12d3-a456-426614174000_1234567890.jpg",
    "/images/123e4567-e89b-12d3-a456-426614174000_1234567891.jpg"
  ],
  "total": 2
}
```

**Ejemplo con cURL:**
```bash
curl -X POST \
  http://localhost:8080/api/vehicles/123e4567-e89b-12d3-a456-426614174000/upload-images \
  -F "images=@/path/to/image1.jpg" \
  -F "images=@/path/to/image2.jpg"
```

### 2. Actualizar/Reemplazar Imágenes de un Vehículo
```
PUT /api/vehicles/:id/update-images
```

**Parámetros:**
- `id` (URL param): UUID del vehículo
- Body (JSON):
```json
{
  "images": [
    "/images/image1.jpg",
    "/images/image2.jpg"
  ]
}
```

**Respuesta exitosa:**
```json
{
  "message": "Imágenes actualizadas exitosamente",
  "vehicle_id": "123e4567-e89b-12d3-a456-426614174000",
  "images": [
    "/images/image1.jpg",
    "/images/image2.jpg"
  ]
}
```

## 🌐 Páginas Web

### 1. Lista de Vehículos e IDs
**URL:** `http://localhost:8080/static/list-vehicles-ids.html`

**Características:**
- 📋 Lista todos los vehículos con sus IDs
- 🔍 Búsqueda en tiempo real
- 📊 Estadísticas (total, con imágenes, sin imágenes)
- 📋 Copiar ID con un click
- 📤 Acceso directo a subir imágenes

**Uso:**
1. Abre la página en tu navegador
2. Busca el vehículo que deseas
3. Copia el ID o haz click en "📤 Subir"

### 2. Subir Imágenes
**URL:** `http://localhost:8080/static/upload-images.html`

**Características:**
- 📤 Subida múltiple de imágenes
- 👁️ Preview de imágenes antes de subir
- ✅ Validación de formato UUID
- 🔄 Redirección automática después de subir

**Uso:**
1. Ingresa el ID del vehículo (UUID)
2. Selecciona una o más imágenes
3. Haz click en "📤 Subir Imágenes"
4. Espera la confirmación

**Acceso directo con ID en URL:**
```
http://localhost:8080/static/upload-images.html?id=123e4567-e89b-12d3-a456-426614174000
```

### 3. Ver Vehículos (Catálogo)
**URL:** `http://localhost:8080/static/RequestVehicle.html`

Aquí podrás ver todos los vehículos con sus imágenes cargadas.

## 📁 Estructura de Archivos

```
Go-Vehicle-Management-API-feature0001/
├── static/
│   ├── images/                          # Carpeta donde se guardan las imágenes
│   │   ├── 208.jpg
│   │   ├── boraRs.jpg
│   │   ├── chevy.jpg
│   │   ├── honda.jpg
│   │   ├── default-car.jpg             # Imagen por defecto
│   │   └── [UUID]_[timestamp].jpg      # Imágenes subidas
│   ├── upload-images.html              # Página para subir imágenes
│   ├── list-vehicles-ids.html          # Página para listar IDs
│   └── RequestVehicle.html             # Catálogo de vehículos
├── handlers/
│   └── vehicles.go                      # Handler con funciones de upload
└── app/
    └── app.go                           # Configuración de rutas
```

## 🔧 Flujo de Trabajo Recomendado

### Para agregar imágenes a vehículos existentes:

1. **Listar vehículos y obtener IDs:**
   ```
   http://localhost:8080/static/list-vehicles-ids.html
   ```

2. **Subir imágenes:**
   - Opción A: Click en "📤 Subir" desde la lista
   - Opción B: Ir a `upload-images.html` y pegar el ID

3. **Verificar en el catálogo:**
   ```
   http://localhost:8080/static/RequestVehicle.html
   ```

### Para crear un nuevo vehículo con imágenes:

1. **Crear vehículo via API:**
   ```bash
   curl -X POST http://localhost:8080/api/vehicles \
     -H "Content-Type: application/json" \
     -d '{
       "vehicle_title": "Toyota Yaris 2020",
       "vehicle_brand": "Toyota",
       "vehicle_model": "Yaris",
       "vehicle_price": 15000,
       "vehicle_year": 2020,
       "vehicle_version": "XLS",
       "vehicle_state": "Nuevo",
       "vehicle_image": []
     }'
   ```

2. **Copiar el ID del vehículo creado**

3. **Subir imágenes usando la interfaz web o API**

## 🔍 Verificar Imágenes en la Base de Datos

Las imágenes se almacenan en el campo `vehicle_image` como un array JSON:

```sql
SELECT id, vehicle_brand, vehicle_model, vehicle_image 
FROM vehicles 
WHERE id = '123e4567-e89b-12d3-a456-426614174000';
```

Resultado esperado:
```
id                                   | vehicle_brand | vehicle_model | vehicle_image
-------------------------------------|---------------|---------------|----------------------------------
123e4567-e89b-12d3-a456-426614174000| Toyota        | Yaris         | ["/images/yaris.jpg"]
```

## ⚠️ Validaciones

- **Formato UUID:** El ID del vehículo debe ser un UUID válido
- **Tipos de archivo:** Solo se permiten: jpg, jpeg, png, gif, webp
- **Vehículo existente:** El vehículo debe existir en la base de datos
- **Tamaño máximo:** Depende de la configuración del servidor (por defecto Gin: 32MB)

## 🐛 Solución de Problemas

### Las imágenes no se muestran en el catálogo

**Causa:** El campo `vehicle_image` en la BD no es un array JSON válido.

**Solución:**
1. Verifica el contenido en la BD:
   ```sql
   SELECT vehicle_image FROM vehicles WHERE id = 'tu-id';
   ```

2. Si está vacío o mal formateado, actualiza:
   ```sql
   UPDATE vehicles 
   SET vehicle_image = '[]'::jsonb 
   WHERE vehicle_image IS NULL;
   ```

3. Vuelve a subir las imágenes usando la interfaz web

### Error "Vehículo no encontrado"

**Causa:** El ID no existe o es inválido.

**Solución:**
1. Verifica el ID en `list-vehicles-ids.html`
2. Asegúrate de usar el formato UUID correcto

### Error al subir archivos grandes

**Causa:** Límite de tamaño de archivo excedido.

**Solución:**
Aumenta el límite en Gin (agregar en `app.go`):
```go
router.MaxMultipartMemory = 50 << 20 // 50 MB
```

## 📝 Notas Importantes

1. **Backup:** Las imágenes se almacenan en `static/images/`. Asegúrate de hacer backup regular.

2. **Formato de nombres:** Los archivos subidos se renombran automáticamente a:
   ```
   [vehicleID]_[timestamp][extension]
   ```
   Ejemplo: `123e4567-e89b-12d3-a456-426614174000_1704123456.jpg`

3. **Rutas relativas:** En la BD se guardan rutas relativas que empiezan con `/images/`

4. **Array vacío vs NULL:** Asegúrate que `vehicle_image` sea `[]` (array vacío) y no `null`

## 🎨 Personalización

### Cambiar la imagen por defecto

Reemplaza el archivo: `static/images/default-car.jpg`

### Cambiar el tamaño máximo de archivos

En `app.go`:
```go
router.MaxMultipartMemory = 10 << 20 // 10 MB
```

### Agregar más formatos de imagen

En `handlers/vehicles.go`, línea ~172:
```go
if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" && ext != ".webp" && ext != ".bmp" {
```

## 📞 Soporte

Si encuentras problemas:
1. Verifica los logs del servidor Go
2. Revisa la consola del navegador (F12)
3. Asegúrate que el servidor esté corriendo en `http://localhost:8080`

## ✅ Checklist de Implementación

- [x] Endpoint de subida de imágenes
- [x] Endpoint de actualización de imágenes
- [x] Página web para listar IDs
- [x] Página web para subir imágenes
- [x] Validación de formatos
- [x] Almacenamiento en servidor
- [x] Actualización de BD
- [x] Imagen por defecto
- [x] Interfaz responsiva
- [x] Preview de imágenes

¡Todo listo para usar! 🎉