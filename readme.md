# Maratona FullCycle 4.0

### Aula: Autenticação com OpenID Connect e Keycloak

- [OAuth 2 e OpenID Connect com Keycloak](https://youtu.be/K9SLsPUzApY)
- [Keycloak on Docker](https://www.keycloak.org/getting-started/getting-started-docker)

#### Configurar Keycloack:
- Rodar o keycloack via docker: `docker run -p 8080:8080 -e KEYCLOAK_USER=admin -e KEYCLOAK_PASSWORD=admin quay.io/keycloak/keycloak:11.0.1`
- Configurar o Keycloak: realm, clients, users via [Admim Console](http://localhost:8080)

#### Rodar aplicação:
```
> go run client/main.go
> http://localhost:8081
```



 



