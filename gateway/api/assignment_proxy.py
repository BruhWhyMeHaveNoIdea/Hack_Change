import os
import requests
from fastapi import APIRouter, HTTPException

router = APIRouter()

# Порт 8080 для assignment-service
ASSIGNMENT_SVC_URL = os.getenv("ASSIGNMENT_SVC_URL", "http://assignment-service:8080")

@router.post("/assignments/upload")
async def upload_assignment(request: dict):
    try:
        response = requests.post(
            f"{ASSIGNMENT_SVC_URL}/upload",  # http://assignment-service:8080/upload
            json=request,
            timeout=30
        )
        response.raise_for_status()
        return response.json()
    except requests.exceptions.RequestException as e:
        raise HTTPException(status_code=502, detail=f"Assignment service unavailable: {str(e)}")

@router.get("/assignments/{assignment_id}/feedback")
async def get_feedback(assignment_id: int):
    try:
        response = requests.get(
            f"{ASSIGNMENT_SVC_URL}/{assignment_id}/feedback",  # http://assignment-service:8080/{assignment_id}/feedback
            timeout=10
        )
        response.raise_for_status()
        return response.json()
    except requests.exceptions.RequestException as e:
        raise HTTPException(status_code=502, detail=f"Assignment service unavailable: {str(e)}")