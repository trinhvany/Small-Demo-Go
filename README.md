# MyProject

## Mô tả

MyProject là một ứng dụng Go được thiết kế để:
- Kiểm tra kho xem có đáp ứng đơn hàng không, xuất lỗi và gửi mail nếu không đủ hàng
- Áp dụng để tính toán khuyến mãi
- Gửi email thông báo cho khách hàng là đơn thành công
- Cập nhật trạng thái đơn hàng trong database

---

## Cài đặt

### Yêu cầu

- **Go**: Phiên bản 1.18 hoặc cao hơn.
- **MySQL**: Để lưu trữ dữ liệu sản phẩm và đơn hàng.
- **Gomail**: Thư viện Go để gửi email qua SMTP.

---

### Các bước cài đặt

1. Clone repository:
	```bash
	git clone https://github.com/trinhvany/small-demo-go.git
	cd small-demo-go

2. Create table
	```bash
	-- Tạo bảng products
	CREATE TABLE products (
		id INT AUTO_INCREMENT PRIMARY KEY,
		product_name VARCHAR(255) NOT NULL,
		quantity INT NOT NULL
	);

	-- Tạo bảng customers
	CREATE TABLE customers (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		address TEXT NOT NULL
	);

	-- Tạo bảng orders
	CREATE TABLE orders (
		id INT AUTO_INCREMENT PRIMARY KEY,
		customer_id INT,
		quantity INT NOT NULL,
		product_id INT,
		total_amount DECIMAL(10, 2) NOT NULL,
		discount DECIMAL(10, 2) DEFAULT 0,
		tax DECIMAL(10, 2) DEFAULT 0,
		status INT DEFAULT 0,
		create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (customer_id) REFERENCES customers(id),
		FOREIGN KEY (product_id) REFERENCES products(id)
	);

3. Insert data
	```bash
	-- Insert dữ liệu vào bảng products
	INSERT INTO products (product_name, quantity)
	VALUES 
		('Laptop', 50),
		('Smartphone', 100),
		('Tablet', 30);

	-- Insert dữ liệu vào bảng customers
	INSERT INTO customers (name, email, address)
	VALUES 
		('John Doe', 'john@example.com', '123 Main St, City, Country'),
		('Jane Smith', 'jane@example.com', '456 Oak St, City, Country'),
		('Alice Brown', 'alice@example.com', '789 Pine St, City, Country');

	-- Insert dữ liệu vào bảng orders
	INSERT INTO orders (customer_id, quantity, product_id, total_amount, discount, tax, status)
	VALUES 
		(1, 2, 1, 2000.00, 100.00, 150.00, 1),  -- Đơn hàng của John Doe, 2 laptop
		(2, 1, 2, 500.00, 50.00, 30.00, 1),    -- Đơn hàng của Jane Smith, 1 smartphone
		(3, 3, 3, 900.00, 75.00, 60.00, 0);    -- Đơn hàng của Alice Brown, 3 tablet

5. Change infor db and mail in db.go and sendMail.go

4. Run repo:
	```bash
	go run main.go

