'''
User of that class should be allowed to create only
Single-Instance of that class.

Single underscore is just for refrence that variable is provate.
_instance = SemiPrivate but we can access it outside class using class instance or class-Name

__instance = SuperPrivate, we can not access this outside class but there is always a Hack!
'''


class Singleton:
    __instance = "Great-Developers"
    def __init__(self):
        #import pdb;pdb.set_trace()
        if Singleton.__instance!="Great-Developers":
            raise Exception("This is not Allowed")
        else:
            Singleton.__instance=self
Singleton()
# Will through Exception at this point!
Singleton()
