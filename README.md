# Hasura authentication webhook

The webhook enables integration between Hasura GraphQL engine and Ory Cloud

## Prerequisites

For successful integration you need to register the following services

1. Ory Cloud with CNAME support (Start Up plan)
2. Hasura free tier
3. Ngrok for local environment

## Domains setup

1. graphql.example.com - CNAME for Hasura
2. auth.example.com - CNAME for Ory Cloud. Cookie domain should be `example.com`
3. You can deploy webhook anywhere because Hasura sends data via post request.

## Configuring for local enviroment

```
git clone git@github.com:gen1us2k/hasura-ory-webhook
cd hasura-ory-webhook
go run cmd/main.go
ngrok http 8090
```

Copy public URL and set environment variables for Hasura Cloud project

```
HASURA_GRAPHQL_AUTH_HOOK=https://...ngrok.io
HASURA_GRAPHQL_AUTH_HOOK_MODE=post
```

## Deploying it to production
