# Internal Game Server Status API

This API is designed for internal use to check the status of various game servers.

## Supported Games

1. Minecraft
2. Ragnarok Online

## Endpoints

### GET /status

Query Parameters:
- `games`: Integer value representing the game server to check
  - `0`: Minecraft
  - `1`: Ragnarok Online

### Response :

Offline: 
```json
{
   "online": false,
   "server_status": "Offline!"
}
```

Online:
```json
{
   "online": true,
   "players": {
       "online": 6
   },
   "server_status": "Online!"
}
```