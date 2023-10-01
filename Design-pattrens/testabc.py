from abc import ABCMeta, abstractmethod

# class C():
#     @abstractmethod
#     def my_abstract_method(self):
#         print('foo')
#
# class D(C):
#     pass
#
# x = C()
# y = D()

# More about ABC can be learnt from here as well:
# https://pymotw.com/3/abc/
from abc import ABCMeta, abstractmethod

class C(metaclass=ABCMeta):
    @abstractmethod
    def my_abstract_method(self):
        print('foo')

class D(C):
    pass

x = C()

y = D()