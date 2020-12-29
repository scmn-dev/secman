from django.urls import path
from . import views

urlpatterns = [
    path('', views.home, name="home"),
    path('blog/', views.blog, name="blog"),
    path('docs/', views.docs, name="docs"),
    path('doc/', views.doc, name="doc"),
]
