# Secure Books
This is a slightly harder WACTF challenge that requires you to break into an application and then exploit it via a client side authentication and SQL Injection vulnerability respectively.

This challenge will be themed around an online accounting platform, hence the 'books' part of the name, called secure books. Basically it is just a web app CRUD interface to let you fill in a table of basically financial records with the filter based on the owner of the record.

## Technology
The application will make use of the following technology:
* Code-Igniter
* Apache2
* SQLite
* [wkhtmltopdf](https://github.com/wkhtmltopdf/wkhtmltopdf)

## Database Layout
This application will be fairly simple in regards to how it's database is set out. To facilitat the required functionality we will use the three tables:
* Users
* Records
* Secrets

### Users Table
The users table will contain user information. This is going to be roughly along the lines of the following:
* ID (Primary Key)
* Username
* Password

Key things to note here are that the password field will be a hashed value (but using a bad algorithm).

User registration will be disabled requiring the challengers to exploit the client side authentication component in order to get into the application.

### Records Table
This table will refer to all the records entereed in by a user. I guess I could add functionality to support uploading a CSV for this to bulk add records but maybe later huh.
* ID (Primary Key)
* Description
* Amount ($$$)
* Who_From
* Who_To
* Type (Credit/Debit)
* T_Month (Month)
* Owner

We will escape input that is displayed in the web app so that there is no XSS but there is the SSRF. Only esacpe when displaying.

### Secrets Table
This contains the flag with the following schema:
* id
* flag
