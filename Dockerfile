# Usar una imagen base oficial de Go
FROM golang:lastest

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos de la aplicación al contenedor
COPY . .

# Descargar las dependencias de la aplicación
RUN go mod download

#
ENV PORT 8000

# Construir la aplicación
RUN go build

# Comando para ejecutar la aplicación
CMD ["./api_tracking"]