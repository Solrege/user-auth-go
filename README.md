Este proyecto es una API REST desarrollada en Golang que implementa un sistema de autenticación de usuarios para una aplicación similar a una red social. Los usuarios pueden crear publicaciones, comentar en las publicaciones de otros, dar likes, y seguir a otros usuarios.

# Tecnologías Utilizadas
- Lenguaje: Go (Golang)
- Framework: Gin para manejo de rutas HTTP y middleware
- Base de Datos: MySQL (mediante gorm)
- Autenticación: JWT (JSON Web Tokens)

# Funcionalidades Principales
- Autenticación de usuarios: Registro, inicio de sesión, y gestión de tokens JWT.
- Publicaciones: CRUD para publicaciones de texto.
- Interacción de usuarios: Comentar, dar "likes", seguir y dejar de seguir a otros usuarios.
- Gestión de usuarios: Actualización y eliminación de cuentas, gestión de seguidores y seguidos.
- Manejo de errores: Middleware para la validación y manejo de errores.
