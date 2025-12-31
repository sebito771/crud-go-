# 🎮 Jugadores CRUD – Fullstack (Go + Gin & Vanilla JS)

Este proyecto es un sistema de gestión de jugadores que demuestra una integración completa entre un **Backend** robusto desarrollado en Go y un **Frontend** moderno y ligero con Vanilla JavaScript. La aplicación permite realizar todas las operaciones de un CRUD (Crear, Leer, Actualizar y Borrar).

---

## 🚀 Tecnologías Utilizadas

### **Backend**
* **Lenguaje:** [Go (Golang)](https://golang.org/)
* **Web Framework:** [Gin Gonic](https://gin-gonic.com/)
* **Base de Datos:** MariaDB / MySQL
* **Gestión de Variables:** `godotenv`
* **Persistencia:** `database/sql`

### **Frontend**
* **Lenguaje:** JavaScript (ES Modules)
* **Estilos:** CSS3
* **Estructura:** HTML5
* **Comunicación:** Fetch API

---

## 🧱 Arquitectura del Proyecto (Capas)

El backend sigue un diseño de **Arquitectura por Capas** para separar las responsabilidades y facilitar el mantenimiento.



### **Estructura de Directorios**

* `db/`: Configuración y pool de conexión a MariaDB.
* `models/`: Definición de las entidades (Structs de Go).
* `dto/`: Data Transfer Objects (especialmente para actualizaciones parciales con punteros).
* `repository/`: Consultas SQL y comunicación directa con la base de datos.
* `services/`: Lógica de negocio y validaciones.
* `handlers/`: Controladores HTTP, manejo de rutas y respuestas JSON.
* `main.go`: Punto de entrada que inicializa las dependencias y levanta el servidor.

---

## 📌 Endpoints de la API

| Método | Endpoint | Acción |
| :--- | :--- | :--- |
| **GET** | `/jugadores` | Lista todos los jugadores registrados. |
| **GET** | `/jugadores/:id` | Obtiene los detalles de un jugador por ID. |
| **POST** | `/jugadores` | Crea un nuevo registro de jugador. |
| **PATCH** | `/jugadores/:id` | Actualiza campos específicos (Nombre/Puntaje). |
| **DELETE** | `/jugadores/:id` | Elimina un jugador permanentemente. |

---

## 🎨 Interfaz Web (Frontend)

El frontend está diseñado como una **SPA (Single Page Application)** modular.

* **`index.html`**: Contiene la estructura, formularios y la tabla de jugadores.
* **`js/api.js`**: Módulo encargado exclusivamente de las peticiones `fetch`.
* **`js/app.js`**: Controlador de la interfaz; maneja eventos de botones, validaciones y renderizado dinámico.
* **`css/styles.css`**: Estilos para una interfaz limpia y responsiva.

---

## ⚙️ Configuración e Instalación

### 1. Requisitos Previos
* **Go** instalado (v1.18 o superior).
* Servidor **MariaDB** o **MySQL** activo.
* Extensión **Live Server** en Visual Studio Code.

### 2. Base de Datos
Crea una base de datos y la tabla correspondiente:
```sql
CREATE DATABASE nombre_db;
USE nombre_db;

CREATE TABLE jugadores (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    puntaje INT NOT NULL DEFAULT 0
);



3. Variables de Entorno
Crea un archivo .env en la raíz del backend:

Fragmento de código

DB_USER=tu_usuario
DB_PASSWORD=tu_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=nombre_db
4. Ejecución del Proyecto
Paso 1: Iniciar el Backend

Bash

go mod tidy
go run main.go
El servidor se ejecutará en: http://localhost:8080

Paso 2: Iniciar el Frontend

Abre el archivo index.html en VS Code.

Haz clic derecho y selecciona "Open with Live Server". El cliente se ejecutará en: http://127.0.0.1:5050