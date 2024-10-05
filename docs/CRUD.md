# Implementing Basic CRUD Functions for Each Table
This document outlines the process for creating basic Create, Read, Update, and Delete (CRUD) operations for each table in our forum project's database. Each operation will be implemented using Go's database/sql package, and we will emphasize the importance of transactions to maintain data integrity.
## 1. Overview of CRUD Operations
What are CRUD Operations?
CRUD stands for Create, Read, Update, and Delete. These operations are fundamental to interacting with databases. Each operation corresponds to a specific SQL command:

    Create: Inserting new records into a table.
    Read: Retrieving existing records from a table.
    Update: Modifying existing records in a table.
    Delete: Removing records from a table.

Why Use Transactions?
Transactions ensure that a series of database operations are completed successfully as a single unit. If any operation within the transaction fails, all changes can be rolled back to maintain data integrity. This is crucial in scenarios where multiple related operations must succeed or fail together.
## 2. Creating CRUD Functions
### 2.1 Create Functions
Example: Create User
The CreateUser function handles user registration by inserting a new user record into the users table.

```go
// CreateUser handles the form submission to create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        // Extract user information from the form
        username := r.FormValue("username")
        email := r.FormValue("email")
        password := r.FormValue("password") // Password should be hashed before storing

        // Setup database connection
        db := SetupDatabase()
        defer db.Close()

        // Start a new transaction
        tx, err := db.Begin()
        if err != nil {
            http.Error(w, "Error starting transaction", http.StatusInternalServerError)
            return
        }

        // SQL statement to insert a new user into the users table
        insertSQL := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`
        _, err = tx.Exec(insertSQL, username, email, password)
        if err != nil {
            tx.Rollback() // Rollback if execution fails
            http.Error(w, "Error executing statement", http.StatusInternalServerError)
            return
        }

        // Commit the transaction if everything is successful
        if err = tx.Commit(); err != nil {
            http.Error(w, "Error committing transaction", http.StatusInternalServerError)
            return
        }

        // Redirect to home page after successful creation
        http.Redirect(w, r, "/", http.StatusSeeOther)
    } else {
        // Redirect for non-POST requests
        http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}
```
Explanation of Create User Process:

    Form Submission: The function first checks if the request method is POST. If it is not, it redirects to the home page.
    Extracting Data: It retrieves user data (username, email, password) from the form submission.
    Database Connection: It establishes a connection to the database using SetupDatabase() and ensures it closes after use.
    Transaction Handling:
        A new transaction is initiated using db.Begin().
        If an error occurs while starting the transaction, an error response is sent back.
    Executing SQL Statement: The SQL command inserts the new user into the users table.
        If execution fails, it rolls back the transaction to undo any changes made during this transaction.
    Committing Changes: If everything goes well, it commits the transaction to save changes permanently.
    Redirecting User: Finally, it redirects the user to the home page after successful registration.

### 2.2 Read Functions
Example: Get Users
The GetUsers function retrieves all users from the database.

```go
// GetUsers retrieves all users from the database
func GetUsers() ([]User, error) {
    db := SetupDatabase()
    defer db.Close()

    rows, err := db.Query("SELECT id, username, email FROM users")
    if err != nil {
      return nil, err
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
      var user User
      if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
          return nil, err
      }
      users = append(users, user)
    }
    
    return users,nil 
}
```
Explanation of Get Users Process:

    Database Connection: The function establishes a connection to the database and ensures it closes after use.
    Executing Query: It executes a SQL query to retrieve all users' IDs, usernames, and emails from the users table.
    Error Handling: If an error occurs during query execution, it returns an error.
    Processing Results:
        It iterates through each row returned by the query.
        For each row, it scans values into a User struct and appends it to a slice of users.
    Returning Results: Finally, it returns the slice of users.

### 2.3 Update Functions
Example: Update User
The UpdateUser function updates an existing user's information in the database.

```go
// UpdateUser updates an existing user's information in the database.
func UpdateUser(userID int64 ,username string ,email string ) error {
   db:=SetupDatabase()
   defer db.Close()

   tx ,err:=db.Begin()
   if err!=nil{
      return fmt.Errorf("error starting transaction: %v",err)
   }

   updateSQL:=`UPDATE users SET username=?, email=? WHERE id=?`
   _,err=tx.Exec(updateSQL ,username,email,userID)

   if err!=nil{
      tx.Rollback()
      return fmt.Errorf("error executing statement: %v",err)
   }

   if err=tx.Commit();err!=nil{
      return fmt.Errorf("error committing transaction: %v",err)
   }

   return nil 
}
```
Explanation of Update User Process:

    Database Connection: The function connects to the database and ensures proper closure afterward.
    Transaction Handling:
        A new transaction is initiated to ensure that all changes are made atomically.
    Executing Update Statement:
        It prepares an SQL statement to update the user's username and email based on their ID.
    Error Handling:
        If any error occurs during execution or committing changes, appropriate actions (rollback or returning errors) are taken to maintain data integrity.

### 2.4 Delete Functions
Example: Delete User
The DeleteUser function removes a user from the database.

```go
// DeleteUser removes a user from the database.
func DeleteUser(userID int64) error {
    db := SetupDatabase()
    defer db.Close()

    tx, err := db.Begin()
    if err != nil {
       return fmt.Errorf("error starting transaction: %v", err)
    }

    deleteSQL := `DELETE FROM users WHERE id=?`
    _, err = tx.Exec(deleteSQL,userID)

    if err != nil {
       tx.Rollback() // Rollback on error
       return fmt.Errorf("error executing statement: %v", err)
    }

    if err = tx.Commit(); err != nil {
       return fmt.Errorf("error committing transaction: %v", err)
    }

    return nil 
}
```
Explanation of Delete User Process:

    Database Connection: Establishes a connection and ensures closure afterward.
    Transaction Handling:
        Begins a new transaction for deleting user data safely.
    Executing Delete Statement:
        Executes an SQL command that deletes a user based on their ID.
    Error Handling:
        If any error occurs during deletion or committing changes occurs rollback or returning errors as necessary.

Summary of Transaction Usage

    Transactions are crucial when performing multiple related operations on the database to ensure that all changes are applied successfully or none at all.
    Always start with Begin() when you want to group multiple operations together that need atomicity.
    Use Rollback() in case of any errors during execution; this will revert all changes made during that transaction.
    Finally call Commit() only when all operations have been executed successfully; this will save all changes permanently in your database.
