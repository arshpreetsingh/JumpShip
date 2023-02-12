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

from abc import ABCMeta, abstractmethod

class C(metaclass=ABCMeta):
    @abstractmethod
    def my_abstract_method(self):
        print('foo')

class D(C):
    pass

x = C()

y = D()