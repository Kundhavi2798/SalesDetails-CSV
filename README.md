# SalesDetails-CSV
This project involves designing a backend system to efficiently ingest, normalize, and analyze a large-scale CSV dataset containing historical sales data with potentially millions of records. The goal is to structure the raw data into a well-optimized relational database and expose it via a RESTful API for further analysis and reporting.
SetUP :
1. Installed go verion - cmd : "go version"
2. Check postgresSQL connection
3. create database and created collection
-- Create Customers Table

CREATE TABLE customers (
    customer_id VARCHAR(10) PRIMARY KEY,
    customer_name VARCHAR(100),
    customer_email VARCHAR(100),
    customer_address TEXT
);

-- Create Products Table

CREATE TABLE products (
    product_id VARCHAR(10) PRIMARY KEY,
    product_name VARCHAR(100),
    category VARCHAR(50),
    unit_price DECIMAL(10,2)
);

-- Create Orders Table

CREATE TABLE orders (
    order_id INT PRIMARY KEY,
    product_id VARCHAR(10) REFERENCES products(product_id),
    customer_id VARCHAR(10) REFERENCES customers(customer_id),
    region VARCHAR(50),
    date_of_sale DATE,
    quantity_sold INT,
    discount DECIMAL(5,2),
    shipping_cost DECIMAL(10,2),
    payment_method VARCHAR(50)
);

4. Testing handler Functions
     1. Load the details ---> ![csv-Load-success](https://github.com/user-attachments/assets/226afcae-1f39-44a8-9e36-b2a00024412e)
     2. Get the revenue ---> ![Revenue-customers](https://github.com/user-attachments/assets/1c16ed3a-f82f-4d94-bf78-29372441f3be)
     3. Get the top-products --->
         ![top-products1](https://github.com/user-attachments/assets/d6b772dc-1505-4059-85b9-eb217e78b620)
         ![top-products2](https://github.com/user-attachments/assets/a2a0ad62-f10f-4c6b-b261-d34eeb2e54ae)
    4. Get Customers count ---> ![customers](https://github.com/user-attachments/assets/8c050ed0-6045-4622-960c-a71e1d98ce99)
  

   Thanks If any queries
   Contact : kundhavi2798@gmail,com





