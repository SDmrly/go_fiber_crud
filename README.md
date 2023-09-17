
# Golang RestAPI

Golang dili kullanılarak RestAPI crud operasyonları yapıldı

## Kullanılan Teknolojiler

**Server:** Golang, Go Fiber, Go Viper, Go Validator

**Database:** PosgreSQL




## API Kullanımı

#### Tüm kullanıcıları getir

```http
  GET /api/users
```

#### Kullanıcıyı getir

```http
  GET /api/users/:userId
```

#### Kullanıcı ekle

```http
  POST /api/users
```

#### Kullanıcıyı sil

```http
  DELETE /api/users/:userId
```

#### Kullanıcıyı güncelle

```http
  PATCH /api/users/:userId
```

#### Parola güncelle

```http
  PATCH /api/users/change_password/:userId
```




## Ortam Değişkenleri

Bu projeyi çalıştırmak için postgresql bilgilerinizi app.env dosyanızda eklemeniz gerekecek

`POSTGRES_HOST`

`POSTGRES_USER`

`POSTGRES_PASSWORD`

`POSTGRES_DB`

`POSTGRES_PORT`


## Bilgisayarınızda Çalıştırın

Projeyi klonlayın

```bash
  git clone https://github.com/SDmrly/go_fiber_crud.git
```

Proje dizinine gidin

```bash
  cd go_fiber_crud
```

Gerekli paketleri yükleyin

```bash
  go mod tidy
```

Sunucuyu çalıştırın

```bash
  go build . && ./go_fiber_crud
```
