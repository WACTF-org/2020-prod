# Bird House
This is a simple web challenge for WACTF. 

The goal of the challenge is to get challengers to guess the name of the file that has the flag by iterating backwards through uploaded files.

The challenge aims to be a simple PHP web application that supports the following features:

* User account creation
* File upload (images only)
* Browsing of files

Each user will be allowed to upload images and may select to have their images as public or private. 
If an image is private then only the user that uploaded it and the site administrator may view it.
If an image is public then everyone can view it.

The flag the challenger has to retrieve will be a private image uploaded by the admin with a file number of 0.
All files uploaded will increment this number by 1 and there will be several files present on the application so it should be obvious that the numbers are incremental.

To prevent challengers from bruteforcing files on the webserver to find the flag, all images uploaded will be stored in the database instead.
The database will be a simple SQL solution such as SQLite or MySQL/MariaDB.

## Technology
The application will make use of the following technology:
* Code-Igniter
* Apache2
* SQLite

## Database Layout
This application will be fairly simple in regards to how it's database is set out. To facilitat the required functionality we will use the two tables:
* Users
* Files

### Users Table
The users table will contain user information. This is going to be roughly along the lines of the following:
* ID (Primary Key)
* Username
* Password

Key things to note here are that the password field will be a hashed value.

### Files Table
This table will refer to all the files uploaded by users of the application. The structure will be along the following lines:
* ID (Primary Key)
* Description
* Filename
* Filesize
* Filetype
* Owner
* Public

Key things to note here are that Owner will tie back to the username of the user that uploaded the file. The Public column will be a boolean that controls whether the file (image) will be shown to other users. The vulnerability in the application will lie in the fact that this field is only used for the main page where all the uploaded images can be seen. The page that loads an individual image will allow anyone to view an image if they use a valid ID, the primary key, and edit it if they are the owner.