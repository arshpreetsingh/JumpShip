'''
Assuming that On-Boarding a user is very complex process so we have
broken that complex process into following three different classes.

So OnBoardUser() Class will other small processes in required manner to
complete the execution process.

'''
class EmailVerification:
    def start(self):
        print("Email verification started")
class ImageVerification:
    def start(self):
        print("Image Verification Started")
class LocationVerification:
    def start(self):
        print("Location verification Started")

class OnBoardUser:
    def __init__(self):
        self.verify_email = EmailVerification()
        self.verify_image = ImageVerification()
        self.verify_location = LocationVerification()
    def start_verification(self):
        self.verify_email.start()
        self.verify_image.start()
        self.verify_location.start()

# Will call all methods of user verfication into one class
user = OnBoardUser()
user.start_verification()
