### API-MASTER-GOLANG-PRODUCTS
Golang API to register products in SQL Server with the maximum design systems, redundancy and fault tolerance, all for the sole purpose of study


## Sequence paths
- <span style="color:blue">[GET]</span> /api/getjwt - get jwt from key the name Access in Headers
- <span style="color:blue">[GET]</span> /api/validjwt - get Token from exclusive key 
- <span style="color:green">[POST]</span> /api/products - Register new products
- <span style="color:blue">[GET]</span> /api/products/{id} - Get products from specified id 

1. getjwt with Header key Access - 1234 in /api/getjwt
2. With token jwt you can use another API about products

## 📜 Features:

1. API Restfull 
2. Gorm ORM to connect SQL Server
3. Authentication with JWT