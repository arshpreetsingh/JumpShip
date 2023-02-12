'''
Difference Between Factory and Abstract-Factory:

Factory-method is a Method.
Abstract-Factory is an Object.

user-Login class.
There could be multiple methods to make Login,
1. Login using Gmail.
2. Login using OpenID.
3. Login using PhoneOTP
'''

class userLogin:
    def __init__(self, userObject):
        self.userObject = userObject
    def make_login(self):
        # call Login Function
        logged_in_object = self.userObject()
        print("Login",logged_in_object.LOGIN())
        print(logged_in_object.username)
        print(logged_in_object.password)

class gmailLogin:

    def LOGIN(self):
        # Logic to Login for Gmailuser
        return "Login-Successful"
    def username(self):
        username="GmailUser"
        return username
    def password(self):
        password="GmailUserPass"
        return password
# Same methods could be implemented
class openIDLogin:
    pass
class phoneOTPLogin:
    pass

# userLogin is Abstract-Factory-Object providing Logins for multiple IDs
user_login = userLogin(gmailLogin)
user_login.make_login()