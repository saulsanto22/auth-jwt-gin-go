🧠 Go JWT Auth – User Management API
====================================

A simple yet production-ready **RESTful API** built with **Golang (Gin Framework)** and **GORM**, implementing secure **JWT Authentication** with modular, clean architecture (Handler–Service–Repository pattern).

This project is designed as a **starter kit** for building authentication-based systems using Go — following good practices in structure, error handling, and modularity.

🚀 Features
-----------

✅ **User Authentication**

*   Register new users
    
*   Login with JWT token
    
*   Password hashing using bcrypt
    

✅ **JWT Authorization**

*   Protect routes with middleware
    
*   Access protected data using Bearer Token
    

✅ **User Management**

*   Get logged-in user profile
    
*   Extendable for Admin/User roles
    

✅ **Clean Architecture**

*   Separation between handler, service, and repository
    
*   Reusable and testable components
    

✅ **Ready for Production**

*   Auto DB migration with GORM
    
*   Consistent JSON responses
    
*   Easy to extend (add Swagger, Docker, RBAC, etc.)
    

🗂️ Project Structure
---------------------

Plain textANTLR4BashCC#CSSCoffeeScriptCMakeDartDjangoDockerEJSErlangGitGoGraphQLGroovyHTMLJavaJavaScriptJSONJSXKotlinLaTeXLessLuaMakefileMarkdownMATLABMarkupObjective-CPerlPHPPowerShell.propertiesProtocol BuffersPythonRRubySass (Sass)Sass (Scss)SchemeSQLShellSwiftSVGTSXTypeScriptWebAssemblyYAMLXML`   .  ├── config/             # Database configuration  ├── handler/            # Handles HTTP request/response  ├── middleware/         # JWT authentication middleware  ├── model/              # GORM models  ├── repository/         # Database queries  ├── routes/             # Route registration  ├── service/            # Business logic layer  ├── utils/              # Response & helper functions  └── main.go             # App entry point   `

⚙️ Installation & Setup
-----------------------

### 1️⃣ Clone the repository

Plain textANTLR4BashCC#CSSCoffeeScriptCMakeDartDjangoDockerEJSErlangGitGoGraphQLGroovyHTMLJavaJavaScriptJSONJSXKotlinLaTeXLessLuaMakefileMarkdownMATLABMarkupObjective-CPerlPHPPowerShell.propertiesProtocol BuffersPythonRRubySass (Sass)Sass (Scss)SchemeSQLShellSwiftSVGTSXTypeScriptWebAssemblyYAMLXML`   git clone https://github.com/saulsanto22/auth-jwt-gin-go.git  cd auth-jwt-gin-go   `

### 2️⃣ Install dependencies

Plain textANTLR4BashCC#CSSCoffeeScriptCMakeDartDjangoDockerEJSErlangGitGoGraphQLGroovyHTMLJavaJavaScriptJSONJSXKotlinLaTeXLessLuaMakefileMarkdownMATLABMarkupObjective-CPerlPHPPowerShell.propertiesProtocol BuffersPythonRRubySass (Sass)Sass (Scss)SchemeSQLShellSwiftSVGTSXTypeScriptWebAssemblyYAMLXML`   go mod tidy   `

### 3️⃣ Setup MySQL Database

Buat database baru, misalnya go\_rest.

### 4️⃣ Edit file config/config.go

Ubah baris DSN sesuai konfigurasi lokalmu:

Plain textANTLR4BashCC#CSSCoffeeScriptCMakeDartDjangoDockerEJSErlangGitGoGraphQLGroovyHTMLJavaJavaScriptJSONJSXKotlinLaTeXLessLuaMakefileMarkdownMATLABMarkupObjective-CPerlPHPPowerShell.propertiesProtocol BuffersPythonRRubySass (Sass)Sass (Scss)SchemeSQLShellSwiftSVGTSXTypeScriptWebAssemblyYAMLXML`   dsn := "root:@tcp(127.0.0.1:3306)/go_rest?charset=utf8mb4&parseTime=True&loc=Local"   `

### 5️⃣ Run the app

Plain textANTLR4BashCC#CSSCoffeeScriptCMakeDartDjangoDockerEJSErlangGitGoGraphQLGroovyHTMLJavaJavaScriptJSONJSXKotlinLaTeXLessLuaMakefileMarkdownMATLABMarkupObjective-CPerlPHPPowerShell.propertiesProtocol BuffersPythonRRubySass (Sass)Sass (Scss)SchemeSQLShellSwiftSVGTSXTypeScriptWebAssemblyYAMLXML`   go run main.go   `

Server akan berjalan di:👉 http://localhost:8080

🔐 API Endpoints
----------------

### Public Routes

MethodEndpointDescriptionPOST/registerRegister new userPOST/loginLogin and get JWT

### Protected Routes

MethodEndpointDescriptionHeader RequirementGET/auth/profileGet user profileAuthorization: Bearer

🧪 Example Requests
-------------------

### Register User

Plain textANTLR4BashCC#CSSCoffeeScriptCMakeDartDjangoDockerEJSErlangGitGoGraphQLGroovyHTMLJavaJavaScriptJSONJSXKotlinLaTeXLessLuaMakefileMarkdownMATLABMarkupObjective-CPerlPHPPowerShell.propertiesProtocol BuffersPythonRRubySass (Sass)Sass (Scss)SchemeSQLShellSwiftSVGTSXTypeScriptWebAssemblyYAMLXML`   POST /register  Content-Type: application/json  {    "name": "Saul Santo",    "email": "saul@example.com",    "password": "123456"  }   `

### Login

Plain textANTLR4BashCC#CSSCoffeeScriptCMakeDartDjangoDockerEJSErlangGitGoGraphQLGroovyHTMLJavaJavaScriptJSONJSXKotlinLaTeXLessLuaMakefileMarkdownMATLABMarkupObjective-CPerlPHPPowerShell.propertiesProtocol BuffersPythonRRubySass (Sass)Sass (Scss)SchemeSQLShellSwiftSVGTSXTypeScriptWebAssemblyYAMLXML`   POST /login  Content-Type: application/json  {    "email": "saul@example.com",    "password": "123456"  }   `

Response:

Plain textANTLR4BashCC#CSSCoffeeScriptCMakeDartDjangoDockerEJSErlangGitGoGraphQLGroovyHTMLJavaJavaScriptJSONJSXKotlinLaTeXLessLuaMakefileMarkdownMATLABMarkupObjective-CPerlPHPPowerShell.propertiesProtocol BuffersPythonRRubySass (Sass)Sass (Scss)SchemeSQLShellSwiftSVGTSXTypeScriptWebAssemblyYAMLXML`   {    "status": "success",    "message": "Berhasil!",    "data": {      "token": "your_jwt_token_here"    }  }   `

🧩 Tech Stack
-------------

*   **Language:** Go (Golang)
    
*   **Framework:** Gin
    
*   **ORM:** GORM
    
*   **Database:** MySQL
    
*   **Auth:** JWT
    
*   **Tools:** bcrypt, Go Modules
    

💡 Future Improvements
----------------------

*   Add Swagger documentation
    
*   Add role-based access control (Admin/User)
    
*   Add Dockerfile & docker-compose
    
*   Add unit tests
    
*   Add refresh token feature
    

👨‍💻 Author
------------

**Saul Santo Anju**Backend Developer — passionate about building clean, secure, and maintainable APIs.

📎 [GitHub](https://github.com/saulsanto22) | [LinkedIn](https://linkedin.com/in/saulsanto22)