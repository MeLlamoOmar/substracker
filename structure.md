project-root/
├── cmd/ # Punto de entrada principal de los binarios
│ └── server/  
│ └── main.go # Archivo principal que inicia el servidor
│
├── internal/ # Código privado de la aplicación (no exportable)
│ ├── server/ # Inicialización y configuración del servidor HTTP
│ ├── handlers/ # Manejadores HTTP (controladores)
│ ├── middleware/ # Middlewares (auth, logging, CORS, etc.)
│ ├── models/ # Modelos de datos (estructuras de la DB)
│ ├── repository/ # Capa de acceso a datos (queries, repositorios)
│ └── service/ # Lógica de negocio de la aplicación
│
├── pkg/ # Código reutilizable por otros proyectos
│ ├── config/ # Configuración global, lectura de .env
│ └── utils/ # Funciones auxiliares o helpers
│
├── web/ # Todo lo relacionado con la interfaz web
│ ├── templates/ # Plantillas Go (HTML)
│ │ ├── layouts/ # Plantillas base (header, footer, layout general)
│ │ └── pages/ # Páginas específicas (home.html, login.html, etc.)
│ └── static/ # Archivos estáticos
│ ├── css/ # Hojas de estilo
│ ├── js/ # Archivos de JavaScript
│ └── img/ # Imágenes
│
├── static/ # Archivos estáticos adicionales (opcional)
│ ├── css/
│ ├── js/
│ └── img/
│
├── tests/ # Pruebas unitarias o de integración
├── .env # Variables de entorno sensibles
├── .gitignore # Archivos/carpetas a ignorar por Git
├── go.mod # Dependencias del proyecto Go
├── README.md # Documentación del proyecto
├── docker-compose.yml # Configuración de contenedores (DB, app, etc.)
└── Makefile # Comandos automatizados (build, up, down, etc.)
