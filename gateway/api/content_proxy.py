from fastapi import APIRouter, Depends, HTTPException
from fastapi.requests import Request
import requests
import os

router = APIRouter(prefix = "/api")

CONTENT_SERVICE_URL = os.getenv("CONTENT_SERVICE_URL", "http://courses-service:9000")

@router.get("/courses")
async def get_courses(request: Request):
    response = requests.get(
        f"{CONTENT_SERVICE_URL}/courses",
        params=dict(request.query_params),
        timeout=30
    )
    return response.json()

@router.get("/courses/{course_id}")
async def get_course_detail(course_id: int, request: Request):
    response = requests.get(
        f"{CONTENT_SERVICE_URL}/courses/{course_id}",
        timeout=30
    )
    return response.json()

@router.post("/progress")
async def update_progress(request: Request):
    data = await request.json()
    response = requests.post(
        f"{CONTENT_SERVICE_URL}/progress",
        json=data,
        timeout=30
    )
    return response.json()