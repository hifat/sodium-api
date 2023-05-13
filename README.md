# Hifat Blog API

## Command
Use [Makefile](https://makefiletutorial.com/)
- Migrate database
```bash
make migrate
```
- Run app
```bash
make run
```

- Swag init
```bash
make swag
```

# Feature

- [x] Auth
- [x] Pretty validate form message
- [ ] Rate limit
- [ ] RBCA

# Issues

- [ ] Error validate message not show index when use `dive`
- [ ] Change key name to json tag when validation error
- [x] Swagger not support some type such as `net.IP`, `utype.IP`