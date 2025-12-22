# CRUD de Jugadores – Go + Gin

Proyecto backend desarrollado en **Go** usando el framework **Gin**.  
Implementa un CRUD completo de jugadores aplicando una **arquitectura por capas**
para separar correctamente responsabilidades.

---

## 🚀 Tecnologías utilizadas

- Go
- Gin
- MariaDB
- database/sql
- godotenv

---

## 🧱 Arquitectura del proyecto

El proyecto sigue una estructura por capas:

- **handlers/**  
  Se encargan únicamente del manejo HTTP (requests, responses y códigos de estado).

- **services/**  
  Contienen la lógica de negocio y validaciones del dominio.

- **repository/**  
  Manejan el acceso a datos y consultas SQL.

- **models/**  
  Definen las entidades principales del sistema.

- **dto/**  
  Data Transfer Objects usados para operaciones parciales (PATCH).

- **db/**  
  Configuración y conexión a la base de datos.

- **main.go**  
  Punto de entrada del proyecto. Inicializa dependencias y levanta el servidor.

---

## 📌 Endpoints disponibles

### ➕ Crear jugador
**POST** `/jugadores`

```json
{
  "nombre": "Carlos",
  "puntaje": 100
}
📄 Obtener todos los jugadores
GET /jugadores

🔍 Obtener jugador por ID
GET /jugadores/{id}

✏️ Actualizar jugador (PATCH parcial)
PATCH /jugadores/{id}

Ejemplos:

json
{
  "nombre": "Nuevo nombre"
}

json
{
  "puntaje": 50
}
Se utilizan DTOs con punteros para permitir la actualización parcial de campos.

🗑️ Eliminar jugador
DELETE /jugadores/{id}



⚙️ Configuración del entorno

Crear un archivo .env en la raíz del proyecto:

DB_USER=usuario
DB_PASSWORD=password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=nombre_db


### Ejecucion del Programa
- go run main.go