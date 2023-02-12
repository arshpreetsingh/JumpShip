'''
Assuming that based on some geolocation
we are providing different content to the user.

factory_method() will provide required streaming access to user
based on geolocation.
'''
class IndianContent:
    def __init__(self,geoLocation):
        self.geolocatoin = geoLocation
    def deliver_content(self):
        content = self.geolocatoin+" Indian-Content Bollywood Movies"
        return content

class FrenchContent:
    def __init__(self,geoLocation):
        self.geolocatoin = geoLocation
    def deliver_content(self):
        content = self.geolocatoin+" French-Content French Movies"
        return content

class HollywoodContent:
    def __init__(self,geoLocation):
        self.geolocatoin = geoLocation

    def deliver_content(self):
        content = self.geolocatoin+" Hollywood-Content English Movies"
        return content

def content_factory(user_type):
    if user_type == "Indian":
        location = "India"
        content_instance = IndianContent(location)
    elif user_type == "French":
        location = "France"
        content_instance = FrenchContent(location)
    else:
        location = "USA"
        content_instance = HollywoodContent(location)
    return content_instance

content_instance = content_factory("Indian")
print(content_instance.deliver_content())

content_instance = content_factory("French")
print(content_instance.deliver_content())

content_instance = content_factory("")
print(content_instance.deliver_content())