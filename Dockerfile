# Etapa de construcción
FROM golang:1.21.3-bookworm as build

WORKDIR /app

# Copia los archivos de descripción de dependencias
COPY go.mod go.sum ./
RUN go mod download

# Instala Git si es necesario
RUN apt-get update && apt-get install --no-install-recommends -y git

# Copia todo el código fuente
COPY . .

# Copia el archivo de base de datos SQLite
COPY task.db ./

# Compila la aplicación
RUN go build -mod=vendor -ldflags '-s -w' -o build/bin/task main.go

# Etapa de producción
FROM gcr.io/distroless/base-debian12:nonroot

ENV GIN_MODE=release
WORKDIR /app

# Copia directorio de configuración
COPY --from=build /app/cmd/configs /app/cmd/configs

# Copia el binario
COPY --from=build /app/build/bin/task /app

# Copia el archivo de base de datos SQLite
COPY --from=build /app/task.db /app

# Exponer puertos
EXPOSE 8086

# Establecer el usuario no root
USER nonroot:nonroot

# Define el comando de entrada
ENTRYPOINT ["./task"]