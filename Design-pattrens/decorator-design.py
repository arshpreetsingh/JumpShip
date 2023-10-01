'''

self.text which is class's object can be inside the class.
this call can be used to create Decorator-pattern for any of the class as required!
'''

sudo kubeadm reset --cri-socket=unix:///var/run/cri-dockerd.sock

sudo kubeadm init --pod-network-cidr = 10.244
.0
.0 / 16 - -cri - socket = unix:///var/run/cri-dockerd.sock - -v = 5

sudo kubeadm reset --cri-socket=unix:///var/run/cri-dockerd.sock
unix:///var/run/containerd/containerd.sock

class Bold:
    def __init__(self,text):
        self.text = text
    def make_bold(self):
        return "<b>"+self.text+"</b>"
class Italic:
    def __init__(self,text):
        self.text = text
    def make_italic(self):
        return "<I>"+self.text+"</I>"
class UL:
    def __init__(self,text):
        self.text = text
    def make_underline(self):
        return "<UL>"+self.text+"</UL>"
def wrapper(text):
    final_text = UL(Italic(Bold(text).make_bold()).make_italic()).make_underline()
    return final_text

print(wrapper("this is me"))