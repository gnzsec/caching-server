# Caching Proxy Server

Este proyecto implementa un servidor proxy de caché que reenvía solicitudes a un servidor de origen y almacena en caché las respuestas. Si se realiza la misma solicitud nuevamente, el servidor proxy devuelve la respuesta almacenada en caché en lugar de reenviarla al servidor de origen.

## Requisitos

- Go (versión 1.16 o superior)

## Instalación

1. Clona este repositorio:

   ```bash
   git clone <URL_DEL_REPOSITORIO>
   cd <NOMBRE_DEL_REPOSITORIO>
   ```

2. Asegúrate de tener Go instalado en tu máquina. Puedes verificarlo ejecutando:

   ```bash
   go version
   ```

## Uso

Para iniciar el servidor proxy, ejecuta el siguiente comando:

```bash
go run main.go -port=3000 -origin=http://dummyjson.com
```

# caching-server

https://roadmap.sh/projects/caching-server