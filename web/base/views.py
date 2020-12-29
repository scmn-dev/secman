from django.shortcuts import render
from django.http import HttpResponse

# Create your views here.

def home(req):
    return render(req, 'base/home.html')

def blog(req):
    return render(req, 'base/blog.html')
    
def docs(req):
    return render(req, 'base/docs.html')

def doc(req):
    return render(req, 'base/doc.html')

