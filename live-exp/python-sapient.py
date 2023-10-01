# Python program to demonstrate private members
# of the parent class


class C(object):
	def __init__(self):
		self.c = 21

		# d is private instance variable
		self.__d = 42
    	def my_data(self):
        	print(self.__d)

class D(object):
	def __init__(self):
		self.e = 84
		self.__d = 22
	def my-data(self):
		print(self.__d)


object1 = D()
#object2 = C()
# produces an error as d is private instance variable
print(object1.e)
