Backend - echo/v4.
Frontend - Templ template rendering - SSR; HTMX(?) - SPA-alike behavior; React(?)
Database - PostgreSQL

Todo:
1. Session "github.com/labstack/echo-contrib/session"
    "github.com/gorilla/sessions". 
  RequireLogin function will be a must. Example here: https://github.com/ArturSS7/TukTuk/blob/master/backend/backend.go
Create adequate login/register logic.
Alternative: JWT

2. Hashing in registration 

3. Live reloading => go install github.com/cosmtrek/air@latest. Create an .air.toml.

3. Funny design. Tailwind.css. And/or React (bruh; long live HTMX).

4. Fix Logout button. Do kinda HTMX that logouts to /login

5. Do I need to query database in AccountHandler? I think it is not effective. Session cashing is better. Redis, e.g.