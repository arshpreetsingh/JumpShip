'''
Prototype Method is to provide Abstract-Layer or possible assumed Structure
for the implementation.

Assuming that we have Class defining Avatar() in a game.
Different people could create Different Set of Avatars for his/her Choice.
basic Structure of Avatar will be same but it would be Customizable from inherited-Class.

basically Prototype is to Update and cerate the Class for basic and abstrac-methods required in the
System.
'''

from abc import abstractmethod, ABC

class Avatar(ABC):

    @abstractmethod
    def face_image(self):
        pass

    @abstractmethod
    def body_size(self):
        pass

    def get_face_image(self):
        return self.face_image()

    def get_body_size(self):
        return self.body_size()

class KrishnAvatar(Avatar):

    def __init__(self):
        super().__init__()

    def face_image(self):
        self.face_image = "Krishna's Face Image"
        print(self.face_image)

    def body_size(self):
        self.body_size = "100"
        print(self.body_size)


class ShivAvatar(Avatar):

    def __init__(self):
        super().__init__()

    def face_image(self):
        self.face_image = "Shiva's Face Image"
        print(self.face_image)

    def body_size(self):
        self.body_size = "200"
        print(self.body_size)
# So from here face_image() and body_size() are isolated for specific user
# But can be called from Prototype class.

krish = KrishnAvatar()
krish.get_face_image()
krish.get_body_size()

shiv = ShivAvatar()
shiv.get_face_image()
shiv.get_body_size()