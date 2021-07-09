## Requirements
1. GCC for sqlite3 requirements. (For windows users, just download [mingw](http://mingw-w64.org/doku.php)
or mingw on [msys2](https://www.msys2.org/))

## How to run (Windows)
1. Initiate module
    ```cmd
    go mod init <mod_name>
    ```

2. Download dependencies:
    ```cmd
    go mod tidy
    ```

3. Build project:
    ```cmd
    go build
    ```

4. Run project:
    ```cmd
    <mod_name>.exe
    ```

5. Open localhost:8081 or just use "Simple Go ORM (GORM)" Folder inside Golang Practice.postman_collection (requires postman) for testing purposes.

6. Run step 2 everytime you add new libraries/dependencies
7. Run step 3 and 4 everytime you make changes to the code