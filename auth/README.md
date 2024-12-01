# Auth сервис

Основная цель сервиса - обеспечить пользователя jwt токеном, для аунтентификации и авторизации.
_В будующем можно накрутить сюда систему ролей_

Используется _jwt_, шифрования по схеме RS256
jwt : header, payload, signature
private key: header, payload --> signature
public key: signature --> header, payload

таким образом только наш сервис умеет создавать ключи, остальные лиш проверять что они были выданны им

## Tech

### Сущности их отношеня и ручки

#### User
* id _соответствует с id мед работна_
* email 
* password _в захешированным виде_
* rtw _refresh token word_

/register
    - -> email, password
    - <- id

/login
    - -> email, password
    - <- access_token, refresh_token

/refresh
    - -> refresh_token
    - <- access_token