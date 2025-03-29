# API de La Liga

API REST para gestionar partidos de La Liga, desarrollada en Go con Gin y GORM.

## Características

- Gestión completa de partidos (CRUD)
- Registro de goles
- Registro de tarjetas amarillas y rojas
- Gestión de tiempo extra
- Documentación con Swagger
- Interfaz web para gestión de partidos
- Soporte CORS
- Base de datos PostgreSQL

## Requisitos

- Go 1.16 o superior
- PostgreSQL
- Docker (opcional)

## Instalación

1. Clonar el repositorio:
```bash
git clone https://github.com/tu-usuario/laliga-api.git
cd laliga-api
```

2. Instalar dependencias:
```bash
go mod download
```

3. Configurar la base de datos:
   - Crear una base de datos PostgreSQL
   - Actualizar las credenciales en `internal/config/database.go`

4. Ejecutar con Docker:
```bash
docker-compose up --build
```

O ejecutar localmente:
```bash
go run main.go
```

## Uso

La API estará disponible en `http://localhost:8080/api`

### Documentación

- Swagger UI: `http://localhost:8080/swagger/index.html`
- Documentación detallada: Ver archivo `llms.txt`
- Colección de Hoppscotch: [Importar Colección](https://hoppscotch.io/es?method=GET&url=http%3A%2F%2Flocalhost%3A8080%2Fapi%2Fmatches&headers=%5B%7B%22key%22%3A%22Content-Type%22%2C%22value%22%3A%22application%2Fjson%22%7D%5D&body=%7B%22homeTeam%22%3A%22Real%20Madrid%22%2C%22awayTeam%22%3A%22Barcelona%22%2C%22matchDate%22%3A%222024-03-20%22%7D)

### Interfaz Web

La interfaz web está disponible en `http://localhost:8080/LaLigaTracker.html`

## Endpoints

### GET /api/matches
Obtiene todos los partidos

### GET /api/matches/:id
Obtiene un partido específico

### POST /api/matches
Crea un nuevo partido

### PUT /api/matches/:id
Actualiza un partido existente

### DELETE /api/matches/:id
Elimina un partido

### PATCH /api/matches/:id/goals
Registra un gol en un partido

### PATCH /api/matches/:id/yellowcards
Registra una tarjeta amarilla

### PATCH /api/matches/:id/redcards
Registra una tarjeta roja

### PATCH /api/matches/:id/extratime
Establece el tiempo extra del partido

## Ejemplos de Peticiones con Hoppscotch

### Crear un Partido
```http
POST http://localhost:8080/api/matches
Content-Type: application/json

{
  "homeTeam": "Real Madrid",
  "awayTeam": "Barcelona",
  "matchDate": "2024-03-20"
}
```

### Registrar un Gol
```http
PATCH http://localhost:8080/api/matches/1/goals
Content-Type: application/json

{
  "team": "home"
}
```

### Registrar Tarjeta Amarilla
```http
PATCH http://localhost:8080/api/matches/1/yellowcards
```

### Establecer Tiempo Extra
```http
PATCH http://localhost:8080/api/matches/1/extratime
Content-Type: application/json

{
  "minutes": 15
}
```

## Estructura del Proyecto

```
.
├── cmd/
│   └── main.go
├── internal/
│   ├── config/
│   │   └── database.go
│   ├── controllers/
│   │   └── match_controller.go
│   └── models/
│       └── match.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── LaLigaTracker.html
├── llms.txt
├── README.md
└── swagger.yaml
```

## Contribuir

1. Fork el repositorio
2. Crear una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abrir un Pull Request

## Licencia

Este proyecto está licenciado bajo la Licencia MIT - ver el archivo [LICENSE](LICENSE) para más detalles.

## Contacto

Javier Benitez - [@ja](https://twitter.com/tutwitter) - email@example.com

Link del Proyecto: [https://github.com/tu-usuario/laliga-api](https://github.com/tu-usuario/laliga-api) 