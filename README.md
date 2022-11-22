# Practicing JWT with GO

Idea is to build a site with the following endpoints:

- `/register` to register as a new user
- `/login` to login after registration
- `/o/welcome` redirect after login; requires auth

*11/22/22 update:*
I sort of got this working, but it's super hacky and probably has too much going on with splitting up files into sub-packages and whatnot. I'll likely recreate this using just the standard library (not gin) and a simpler file structure in the near future.

see [Ben Johnson wtf](https://github.com/benbjohnson) as well as [gorilla securecookie](https://github.com/gorilla/securecookie)