# API Testing Commands (curl)

Gunakan perintah ini untuk mencoba API CRUD Category dan Product.

## 🔗 Base URLs

- **Local:** `http://localhost:8080` (sesuaikan port di `.env`)
- **Production:** `https://learn-go-production.up.railway.app`

---

## 🏥 Health Check

Cek status aplikasi.

```bash
curl -X GET http://localhost:8080/health
```

---

## 📁 Category API

Endpoint: `/categories`

### 1. Get All Categories

```bash
curl -X GET http://localhost:8080/categories
```

### 2. Create a New Category

```bash
curl -X POST http://localhost:8080/categories \
     -H "Content-Type: application/json" \
     -d '{"name": "Hardware", "description": "PC parts and tools"}'
```

### 3. Get Detail of One Category

```bash
curl -X GET http://localhost:8080/categories/1
```

### 4. Update a Category

```bash
curl -X PUT http://localhost:8080/categories/1 \
     -H "Content-Type: application/json" \
     -d '{"name": "Updated Hardware", "description": "Revised description"}'
```

### 5. Delete a Category

```bash
curl -X DELETE http://localhost:8080/categories/1
```

---

## 📦 Product API

Endpoint: `/api/products`

### 1. Get All Products

```bash
curl -X GET http://localhost:8080/api/products
```

### 2. Create a New Product

```bash
curl -X POST http://localhost:8080/api/products \
     -H "Content-Type: application/json" \
     -d '{
          "name": "Laptop Gaming",
          "description": "High performance laptop",
          "price": 15000000,
          "stock": 10
     }'
```

### 3. Get Detail of One Product

```bash
curl -X GET http://localhost:8080/api/products/1
```

### 4. Update a Product

```bash
curl -X PUT http://localhost:8080/api/products/1 \
     -H "Content-Type: application/json" \
     -d '{
          "name": "Laptop Gaming Pro",
          "description": "Powerful gaming machine",
          "price": 17000000,
          "stock": 5
     }'
```

### 5. Delete a Product

```bash
curl -X DELETE http://localhost:8080/api/products/1
```
