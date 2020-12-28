from django.shortcuts import render
from django.http import HttpResponse

# Create your views here.

def home(req):
    return HttpResponse('<h2>Home</h2>')

def posts(req):
    return HttpResponse('<h2>Posts</h2>')

def post(req):
    return HttpResponse('<h2>Post Title</h2>')

def profile(req):
    return HttpResponse('<h2>User Profile</h2>')
