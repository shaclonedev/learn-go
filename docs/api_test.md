# API Testing Commands (curl)

Use these commands to test your Category CRUD API implementation.

## 🔗 Endpoints

### 1. Get All Categories

Retrieve a list of all categories.

```bash
curl -X GET https://learn-go-production.up.railway.app/categories
```

### 2. Create a New Category

Add a new category to the list.

```bash
curl -X POST https://learn-go-production.up.railway.app/categories \
     -H "Content-Type: application/json" \
     -d '{"name": "Hardware", "description": "PC parts and tools"}'
```

### 3. Get Detail of One Category

Retrieve details for a specific category by ID.

```bash
curl -X GET https://learn-go-production.up.railway.app/categories/1
```

### 4. Update a Category

Modify an existing category.

```bash
curl -X PUT https://learn-go-production.up.railway.app/categories/1 \
     -H "Content-Type: application/json" \
     -d '{"name": "Updated Hardware", "description": "Revised description"}'
```

### 5. Delete a Category

Remove a category by ID.

```bash
curl -X DELETE https://learn-go-production.up.railway.app/categories/1
```
