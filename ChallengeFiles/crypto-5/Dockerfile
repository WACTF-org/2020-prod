FROM mcr.microsoft.com/dotnet/core/sdk:3.1 AS builder
WORKDIR /app

COPY docker-files/*.csproj ./
RUN dotnet restore

COPY docker-files/ ./
RUN dotnet publish -c Release -o out

FROM mcr.microsoft.com/dotnet/core/aspnet:3.1
WORKDIR /app
COPY --from=builder /app/out .

USER nobody

ENTRYPOINT ["dotnet", "app.dll"]
