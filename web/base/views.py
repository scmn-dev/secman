from django.shortcuts import render
from django.http import HttpResponse

# Create your views here.

def home(req):
    return HttpResponse('<h2>Home</h2>')

def docs(req):
    return HttpResponse('<h2>Docs</h2>')

def doc(req):
    return HttpResponse('<h2>Doc Title</h2>')

def blog(req):
    return HttpResponse('<h2>Blog</h2>')
