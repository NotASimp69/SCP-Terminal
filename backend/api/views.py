from django.shortcuts import render
from rest_framework.views import Response
from rest_framework import status

# Create your views here.
def fetch(request):
    return Response(status=status.HTTP_200_OK)