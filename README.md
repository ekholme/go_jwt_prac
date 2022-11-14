# Practicing JWT with GO

Idea is to build a site with the following endpoints:

- `/register` to register as a new user
- `/login` to login after registration
- `/0/welcome` redirect after login; requires auth
- `/0/secrets` an endpoint that requires authentication and will show secret values for the logged-in user