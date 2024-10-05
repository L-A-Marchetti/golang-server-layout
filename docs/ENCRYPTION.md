# Implementing a Secure Password System with Bcrypt for SQLite3

## 1. Overview
The password system is crucial for ensuring that user credentials are stored securely in the database. Instead of storing passwords in plain text, we will hash them using the bcrypt algorithm, which is designed to be computationally intensive and resistant to brute-force attacks.

### Why Use Bcrypt?
- Security: Bcrypt automatically handles salting and hashing, making it difficult for attackers to reverse-engineer passwords.
- Cost Factor: Bcrypt allows you to adjust the cost factor, which determines how computationally expensive the hashing process is. This can be increased over time as hardware improves.

## 2. Setting Up Bcrypt
### 2.1 Install the Bcrypt Package
First, ensure that you have the bcrypt package installed in your Go environment:

```bash
go get golang.org/x/crypto/bcrypt
```

### 2.2 Hashing Passwords
When a user registers or updates their password, we need to hash it before storing it in the database.

Example: Hashing a Password

```go
package db

import (
    "golang.org/x/crypto/bcrypt"
    "log"
)

// HashPassword hashes the password using bcrypt
func HashPassword(password string) (string, error) {
    // Generate a hashed password with a default cost
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        log.Println("Error hashing password:", err)
        return "", err // Return error if hashing fails
    }
    return string(hashedPassword), nil // Return hashed password as string
}
```

Explanation of Hashing Process:
- GenerateFromPassword: This function takes the plaintext password and generates a hashed version. The bcrypt.DefaultCost defines the computational cost of hashing.
- Error Handling: If an error occurs during hashing, it is logged, and an error is returned.

### 2.3 Verifying Passwords
When a user attempts to log in, we need to verify that the entered password matches the stored hashed password.

Example: Verifying a Password

```go
// CheckPasswordHash compares a plain password with its hashed version
func CheckPasswordHash(password, hash string) bool {
    // Compare the provided password with the stored hash
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil // Returns true if passwords match
}
```

Explanation of Verification Process:
- CompareHashAndPassword: This function checks if the provided plaintext password matches the hashed password stored in the database.
- Return Value: It returns true if they match; otherwise, it returns false.

### 2.4 Storing Hashed Passwords in the Database
When creating or updating a user's record in your database, ensure that you hash their password before storing it.

Example: Creating a User with Hashed Password

```go
import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

// CreateUser handles user registration and stores hashed password
func CreateUser(db *sql.DB, username string, email string, password string) error {
    // Hash the password before storing it
    hashedPassword, err := HashPassword(password)
    if err != nil {
        return err // Return error if hashing fails
    }

    // SQL statement to insert a new user into the users table
    insertSQL := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`
    
    // Execute insert statement with hashed password
    _, err = db.Exec(insertSQL, username, email, hashedPassword)
    if err != nil {
        return err // Return error if execution fails
    }

    return nil // Successfully created user
}
```

Explanation of User Creation Process:
- Hashing Password: The user's plaintext password is hashed using HashPassword.
- Database Insertion: The hashed password is then stored in the database along with other user details.
- Error Handling: Any errors during hashing or database operations are returned for handling by calling functions.

### 2.5 Verifying User Login
When a user attempts to log in, you need to retrieve their hashed password from the database and compare it with the provided password.

Example: Verifying User Login

```go
// VerifyLogin checks if the provided credentials are valid
func VerifyLogin(db *sql.DB, username string, password string) (bool, error) {
    var hashedPassword string
    
    // Retrieve the hashed password for the given username
    err := db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&hashedPassword)
    if err != nil {
        if err == sql.ErrNoRows {
            return false, nil // User not found
        }
        return false, err // Database error
    }
    
    // Compare the provided password with the stored hash
    if CheckPasswordHash(password, hashedPassword) {
        return true, nil // Passwords match
    }
    
    return false, nil // Passwords don't match
}
```

Explanation of Login Verification Process:
- Retrieve Hashed Password: Query the database to get the stored hashed password for the given username.
- Compare Passwords: Use CheckPasswordHash to compare the provided password with the stored hash.
- Return Result: Return true if the passwords match, false otherwise.

## 3. Encrypting Data in SQLite3 Database
While SQLite3 doesn't have built-in encryption, you can encrypt sensitive data before storing it in the database. Here's an example of how to encrypt and decrypt data using AES encryption:

```go
import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "errors"
    "io"
)

var secretKey = []byte("your-32-byte-secret-key-here")

// EncryptData encrypts the given data using AES encryption
func EncryptData(data string) (string, error) {
    block, err := aes.NewCipher(secretKey)
    if err != nil {
        return "", err
    }

    plaintext := []byte(data)
    ciphertext := make([]byte, aes.BlockSize+len(plaintext))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

    return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// DecryptData decrypts the given encrypted data
func DecryptData(encryptedData string) (string, error) {
    ciphertext, err := base64.URLEncoding.DecodeString(encryptedData)
    if err != nil {
        return "", err
    }

    block, err := aes.NewCipher(secretKey)
    if err != nil {
        return "", err
    }

    if len(ciphertext) < aes.BlockSize {
        return "", errors.New("ciphertext too short")
    }
    iv := ciphertext[:aes.BlockSize]
    ciphertext = ciphertext[aes.BlockSize:]

    stream := cipher.NewCFBDecrypter(block, iv)
    stream.XORKeyStream(ciphertext, ciphertext)

    return string(ciphertext), nil
}
```

To use these functions when storing or retrieving sensitive data:

```go
// Storing encrypted data
sensitiveData := "sensitive information"
encryptedData, err := EncryptData(sensitiveData)
if err != nil {
    // Handle error
}
// Store encryptedData in the database

// Retrieving and decrypting data
decryptedData, err := DecryptData(encryptedData)
if err != nil {
    // Handle error
}
// Use decryptedData
```

## 4. Adding Password Protection to SQLite3 Database
SQLite3 itself doesn't provide built-in password protection for the entire database. However, you can use third-party extensions or implement application-level encryption. Here are two approaches:

### 4.1 Using SQLCipher Extension
SQLCipher is a popular extension that adds encryption to SQLite databases. To use it:

1. Install SQLCipher: Follow the installation instructions for your platform from the SQLCipher documentation.

2. Use the SQLCipher driver in your Go code:

```go
import (
    "database/sql"
    _ "github.com/sqlcipher/sqlcipher"
)

func OpenEncryptedDatabase(dbPath, password string) (*sql.DB, error) {
    db, err := sql.Open("sqlite3", dbPath)
    if err != nil {
        return nil, err
    }

    _, err = db.Exec("PRAGMA key = ?", password)
    if err != nil {
        db.Close()
        return nil, err
    }

    return db, nil
}
```

## 5. References for Further Reading
- Bcrypt Documentation: [Bcrypt Package](https://godoc.org/golang.org/x/crypto/bcrypt)
- SQLite Documentation: [SQLite](https://www.sqlite.org/docs.html)
- SQLCipher Documentation: [SQLCipher](https://www.zetetic.net/sqlcipher/documentation/)
- Go Crypto Package: [crypto Package](https://golang.org/pkg/crypto/)
