# samudai-assignment

This folder consists of 3 APIs: User's signup and login, dashboard and manage access. I have tested them using postman. Below mentioned are the steps and request body to be defined if anyone wants to clone and run it locally.

## Login-Signup-API

* **Register**: POST Request - "localhost:8088/register" 
~ JSON Body: 
```
{
&emsp; "name": "kenny",
&emsp; "username": "ken",
&emsp; "email": "abc@gmail.com",
&emsp; "password": "abc@123"
}
```
Role field will be assigned automatically a default value, i.e. "user"

* **Login**: POST Request - "localhost:8088/login"
~ JSON Body: 
```
{
&emsp; "email": "abc@gmail.com",
&emsp; "password": "abc@123"
}
```

* **Get Users**: GET Request - "localhost:8088/getusers"
&emsp; getting a list of all the users in database

* **Delete Users**: DELETE Request - "localhost:8088/deleteuser/:id"
&emsp; delete the user with the given user id

* **Update Users**: PUT Request - "localhost:8088/updateuser/:id"
```
{
&emsp; "name": "kenny",
&emsp; "username": "ken",
&emsp; "email": "abc@gmail.com",
&emsp; "password": "abc@123"
}
```
* **Find User with email**: GET Request - "localhost:8088/find/:email"
&emsp; finds the user with the provided email id


## Dashboard API

Any user can create an empty dashboard by providing a single name field. Then they can add the widgets from a list of widgets as per their roles and permissions.

* **Create Dashboard**: POST Request - "localhost:8088/:id/createdashboard"
&emsp; creates a dashboard for the user with provided user_id and a string of widgets seperated by colon, 
for eg. we can pass "visualizations:settings" string in widgets field.
```
{
&emsp; "dashboard_name": "My Dashboard",
&emsp; "widgets": "",
}
```

* **Get Dashboard**: GET Request - "localhost:8088/:id/getdashboard"
&emsp; gets a list of dashboard created by a particular user by providing their user_id

* **Add Widgets**: GET request - "localhost:8088/addwidget/:id/:widget"
&emsp; adds a widgets in the user dashboard by providing dashboard id and widget name, i.e. "localhost:8088/addwidget/2/settings"


## Manage Access API

Admin can give or revoke access from users. By default, **user** is the default role given to any user, which can then be changed by admin. Moderator or any other user can't give access to users.

- There are 3 roles at present: user, moderator, admin

* **Change Role**: POST Request - "localhost:8088/changerole"
```
{
&emsp; "admin_id": "1",
&emsp; "user_id": "3",
&emsp; "role": "moderator"
}
```

* **Get Permissions**: GET Request - "localhost:8088/getPermissions/:id"
Provides a list of permissions provided to particular user by providing their user_id.
Permissions List:
- login
- signup
- logout
- reset
- settings
- visualizations
- read
- write
- verify

For more on permissions and roles that which role has what permissions, check out config/db.go 


