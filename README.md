# E-Commerce
Fullstack Shopping Cart App built with Go (Gin, GORM, JWT) and React. Features include user registration &amp; login, item browsing, cart management, order checkout, and order history. Backend uses RESTful APIs with SQLite, frontend provides a simple UI for full shopping flow.
cd ecommerce-backend
go mod init ecommerce-backend  
go mod tidy
go run main.go

#Authentication Flow:
User logs in with /users/login
JWT is returned and stored in localStorage
All cart and order operations require the token in

UI Screens:
Login Page
Item List with Add to Cart
Buttons for Checkout, Cart, and Order History
Window alerts show cart content and order status
